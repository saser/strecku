package server

import (
	"context"
	"net"
	"testing"

	streckuv1 "github.com/Saser/strecku/saser/strecku/v1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/testing/protocmp"
)

const bufSize = 1024 * 1024

type fixture struct {
	srv *Server
	lis *bufconn.Listener
}

func setUp(t *testing.T) *fixture {
	t.Helper()

	f := &fixture{
		srv: New(),
		lis: bufconn.Listen(bufSize),
	}
	s := grpc.NewServer()
	streckuv1.RegisterStreckUServer(s, f.srv)
	go func() {
		if err := s.Serve(f.lis); err != nil {
			t.Errorf("s.Serve(f.lis) = %v", err)
			t.FailNow()
		}
	}()
	t.Cleanup(s.GracefulStop)
	return f
}

func (f *fixture) dial(context.Context, string) (net.Conn, error) {
	return f.lis.Dial()
}

func (f *fixture) insecureClient(ctx context.Context, t *testing.T) streckuv1.StreckUClient {
	t.Helper()
	cc, err := grpc.DialContext(ctx, "bufconn", grpc.WithContextDialer(f.dial), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	return streckuv1.NewStreckUClient(cc)
}

func (f *fixture) backdoorCreateUser(t *testing.T, user *streckuv1.User, password string) {
	t.Helper()
	f.srv.users = append(f.srv.users, user)
	f.srv.userKeys[user.Name] = len(f.srv.users) - 1
	t.Cleanup(func() {
		delete(f.srv.userKeys, user.Name)
	})
	f.srv.passwords[user.Name] = password
	t.Cleanup(func() {
		delete(f.srv.passwords, user.Name)
	})
}

func TestServer_AuthenticateUser(t *testing.T) {
	ctx := context.Background()

	f := setUp(t)
	emailAddress := "user@example.com"
	password := "password"
	user := &streckuv1.User{
		Name:         newUserName(),
		EmailAddress: emailAddress,
		DisplayName:  "User",
	}
	f.backdoorCreateUser(t, user, password)
	client := f.insecureClient(ctx, t)

	for _, test := range []struct {
		name     string
		req      *streckuv1.AuthenticateUserRequest
		wantUser *streckuv1.User
		wantCode codes.Code
	}{
		{name: "Correct", req: &streckuv1.AuthenticateUserRequest{EmailAddress: emailAddress, Password: password}, wantUser: user},
		{name: "IncorrectEmailAddress", req: &streckuv1.AuthenticateUserRequest{EmailAddress: "incorrect@example.com", Password: password}, wantCode: codes.Unauthenticated},
		{name: "IncorrectPassword", req: &streckuv1.AuthenticateUserRequest{EmailAddress: emailAddress, Password: "incorrect password"}, wantCode: codes.Unauthenticated},
	} {
		t.Run(test.name, func(t *testing.T) {
			authUser, err := client.AuthenticateUser(ctx, test.req)
			if err == nil {
				if diff := cmp.Diff(test.wantUser, authUser, protocmp.Transform()); diff != "" {
					t.Errorf("test.wantUser != authUser (-want +got):\n%v", diff)
				}
			} else {
				if got := status.Code(err); got != test.wantCode {
					t.Errorf("status.Code(err) = %v; want %v", got, test.wantCode)
				}
			}
		})
	}
}

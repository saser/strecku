//+build tools

package tools

import (
	_ "github.com/googleapis/api-linter/cmd/api-linter"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

package purchases

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

func TestGenerateName(t *testing.T) {
	got := GenerateName()
	prefix := CollectionID + "/"
	if !strings.HasPrefix(got, prefix) {
		t.Errorf("GenerateName() = %q; want prefix %q", got, prefix)
	}
	id := strings.TrimPrefix(got, prefix)
	if _, err := uuid.Parse(id); err != nil {
		t.Errorf("uuid.Parse(%q) err = %v; want nil", id, err)
	}
}

func TestValidateName(t *testing.T) {
	for _, test := range []struct {
		name string
		want error
	}{
		{name: "purchases/6f2d193c-1460-491d-8157-7dd9535526c6", want: nil},
		{name: "", want: ErrNameEmpty},
		{name: "invalidprefix/6f2d193c-1460-491d-8157-7dd9535526c6", want: ErrNameInvalidFormat},
		{name: "purchases/not a UUID", want: ErrNameInvalidFormat},
		{name: "6f2d193c-1460-491d-8157-7dd9535526c6", want: ErrNameInvalidFormat},
	} {
		if got := ValidateName(test.name); !cmp.Equal(got, test.want, cmpopts.EquateErrors()) {
			t.Errorf("ValidateName(%q) = %v; want %v", test.name, got, test.want)
		}
	}
}
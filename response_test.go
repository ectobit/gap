package gap_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.ectobit.com/gap"
)

func TestAddError(t *testing.T) {
	t.Parallel()

	res := &gap.Response[struct{}]{}
	res.AddError("test")

	if diff := cmp.Diff([]string{"test"}, res.Errors); diff != "" {
		t.Errorf("Wrap() mismatch (-want +got):\n%s", diff)
	}
}

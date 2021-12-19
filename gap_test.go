package gap_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.ectobit.com/gap"
	"go.ectobit.com/lax"
	"go.uber.org/zap/zaptest"
)

func TestWrapGet(t *testing.T) {
	t.Parallel()

	helloHandler := func(req *gap.Request[struct{}]) *gap.Response[hello] {
		return &gap.Response[hello]{
			Data: &hello{
				Message: "Hello world!",
			},
		}
	}

	server := httptest.NewServer(http.HandlerFunc(gap.Wrap(helloHandler, lax.NewZapAdapter(zaptest.NewLogger(t)))))

	res, err := http.DefaultClient.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	defer close(t, res.Body)

	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(`{"data":{"message":"Hello world!"}}`, string(got)); diff != "" {
		t.Errorf("Wrap() mismatch (-want +got):\n%s", diff)
	}

}

func close(t *testing.T, body io.Closer) {
	t.Helper()

	if err := body.Close(); err != nil {
		t.Fatal(err)
	}
}

type hello struct {
	Message string `json:"message"`
}

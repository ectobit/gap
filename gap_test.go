package gap_test

import (
	"bytes"
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

	if diff := cmp.Diff(helloJSON(true), string(got)); diff != "" {
		t.Errorf("Wrap() mismatch (-want +got):\n%s", diff)
	}

}

func TestWrapPost(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in         string
		wantStatus int
		want       string
	}{
		"invalid json body": {"", http.StatusBadRequest, `{"errors":["invalid json body"]}`},
		"ok":                {helloJSON(false), http.StatusOK, helloJSON(true)},
	}

	helloHandler := func(req *gap.Request[hello]) *gap.Response[hello] {
		return &gap.Response[hello]{
			Data: req.Data,
		}
	}

	server := httptest.NewServer(http.HandlerFunc(gap.Wrap(helloHandler, lax.NewZapAdapter(zaptest.NewLogger(t)))))

	for n, test := range tests { //nolint:paralleltest
		test := test

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			buf := bytes.NewBufferString(test.in)

			res, err := http.DefaultClient.Post(server.URL, "application/json", buf)
			if err != nil {
				t.Fatal(err)
			}

			defer close(t, res.Body)

			if test.wantStatus != res.StatusCode {
				t.Fatalf("Wrap() status code %d; want %d", res.StatusCode, test.wantStatus)
			}

			got, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(test.want, string(got)); diff != "" {
				t.Errorf("Wrap() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func close(t *testing.T, body io.Closer) {
	t.Helper()

	if err := body.Close(); err != nil {
		t.Fatal(err)
	}
}

func helloJSON(isResponse bool) string {
	if isResponse {
		return `{"data":{"message":"Hello world!"}}`
	}

	return `{"message":"Hello world!"}}`
}

type hello struct {
	Message string `json:"message"`
}

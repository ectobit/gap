package gap

import "net/http"

type Request[T any] struct {
	*http.Request
	Data *T
}

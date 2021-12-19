package gap

type Response[T any] struct {
	StatusCode int      `json:"-"`
	Data       *T       `json:"data,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}

func NewErrorResponse[O any](statusCode int, err string) *Response[O] {
	r := &Response[O]{StatusCode: statusCode}
	r.Errors = append(r.Errors, err)

	return r
}

func (r *Response[T]) AddError(err string) {
	r.Errors = append(r.Errors, err)
}

package gap

import (
	"encoding/json"
	"net/http"

	"go.ectobit.com/lax"
)

// Handler defines generic HTTP handler function.
type Handler[I, O any] func(*Request[I]) *Response[O]

// Wrap wraps generic handler into Go idiomatic HTTP handler function.
// Request json body is automatically JSON decoded and passed to generic handler.
// Response from generic handler is automatically JSON encoded and passed to idiomatic HTTP handler.
func Wrap[I, O any](handler Handler[I, O], log lax.Logger) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var reqData I

		if req.Method != http.MethodGet {
			if err := json.NewDecoder(req.Body).Decode(&reqData); err != nil {
				log.Warn("json decode", lax.Error(err))
				render(res, NewErrorResponse[O](http.StatusBadRequest, "invalid json body"), log)

				return
			}
		}

		render(res, handler(&Request[I]{req, &reqData}), log)

	})
}

func render[O any](res http.ResponseWriter, response *Response[O], log lax.Logger) {
	res.Header().Set("Content-Type", "application/json")

	resData, err := json.Marshal(response)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)

		return
	}

	if response.StatusCode == 0 {
		response.StatusCode = http.StatusOK
	}

	res.WriteHeader(response.StatusCode)
	if _, err := res.Write(resData); err != nil {
		log.Warn("response write", lax.Error(err))
	}
}

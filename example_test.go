package gap_test

import (
	"fmt"
	"io"
	"net/http"

	"go.ectobit.com/gap"
	"go.ectobit.com/lax"
)

func Example_get() {
	helloHandler := func(req *gap.Request[struct{}]) *gap.Response[hello] {
		return &gap.Response[hello]{
			Data: &hello{
				Message: "Hello world!",
			},
		}
	}

	log, _ := lax.NewDefaultZapAdapter("json", "debug")

	go func() {
		http.HandleFunc("/hello", gap.Wrap(helloHandler, log))
		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Error("listen and serve", lax.Error(err))
		}
	}()

	res, _ := http.DefaultClient.Get("http://localhost:3000/hello")

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error("body close", lax.Error(err))
		}
	}()

	got, _ := io.ReadAll(res.Body)
	fmt.Println(string(got))

	// Output: {"data":{"message":"Hello world!"}}
}

package hello

import (
	"fmt"
	"io"
	"net/http"
)

// HelloWorld is an HTTP Cloud Function.
func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	bytes, _ := io.ReadAll(request.Body)
	parameters := request.URL.Query()
	a := parameters.Get("a")
	b := parameters.Get("b")
	fmt.Fprintf(writer, "a=%s, b=%s, body=%v", a, b, string(bytes))
}

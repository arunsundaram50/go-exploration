package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type MyApp struct {
}

func (a MyApp) ServeHTTP(output http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/hello" {
		output.Write([]byte("hello, world"))
	} else if request.URL.Path == "/add" {
		query := request.URL.Query()
		a := query.Get("a")
		b := query.Get("b")
		i, err := strconv.Atoi(a)
		if err != nil {
			output.Write([]byte("Error a=" + a))
			return
		}
		j, err := strconv.Atoi(b)
		if err != nil {
			output.Write([]byte("Error b=" + b))
			return
		}
		sum := i + j
		s := fmt.Sprintf("a=%s, b=%s, %s+%s=%d", a, b, a, b, sum)
		output.Write([]byte(s))
	} else {
		output.Write([]byte("nothing to do!"))
	}
}

func main() {
	var myApp MyApp
	fmt.Print("Listening to localhost:8080...\n")
	http.ListenAndServe(":8080", myApp)
	fmt.Print("Program terminated\n")
}

package main

import (
	"net/http"
)

type MyApp struct {
}

func (a MyApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		w.Write([]byte("hello, world"))
	} else if r.URL.Path == "/add" {
		w.Write([]byte("1+1="))
	} else {
		w.Write([]byte("nothing to do!"))
	}
}

func main() {
	var myApp MyApp

	http.ListenAndServe(":8000", myApp)
}

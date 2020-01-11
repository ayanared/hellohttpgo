package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World"))
}

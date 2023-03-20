package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", header)
	log.Fatal(http.ListenAndServe("localhost:7700", nil))
}

func header(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "method: %s url: %s", req.Method, req.URL.Path)

	fmt.Fprintf(rw, "\n%s\n", "HTTP Заголвки:")

	for k, v := range req.Header {
		fmt.Fprintf(rw, "Имя заголовка: %s \nЗначение: %s\n", k, v)
		fmt.Fprintf(rw, "%s", "\n")
	}
}


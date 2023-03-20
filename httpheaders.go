package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    serve()
}

// Запускаем обработчик http запросов
func serve() {
	var ip = "localhost"
	var port = "8000" // default port

	if len(os.Args) == 2 {
          port = os.Args[1]
	}

	http.HandleFunc("/", headers)
	http.HandleFunc("/info", reqInfo)

	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
}

func reqInfo(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Method: %s\n", req.Method)
	fmt.Fprintf(rw, "Host: %s\n", req.Host)
	fmt.Fprintf(rw, "Proto: %s\n", req.Proto)
}

func headers(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%s\n", "HTTP Заголовки:")
	for k, v := range req.Header {
		fmt.Fprintf(rw, "%s %s\n", k, v)
	}
}

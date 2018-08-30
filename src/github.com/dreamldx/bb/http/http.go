package http

import (
	"fmt"
	"io"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func StartHTTPServer(port int) {
	fmt.Printf("Start HTTP server at %d\n", port)
	http.HandleFunc("/", helloWorldHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Sprintf("Failed to start server due to %v.\n", err)
	}
}

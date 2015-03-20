package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func main() {
	http.HandleFunc("/hello", hello)

	ln, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen on port: 5555")
	fmt.Printf("concurrent: %d\n", 4)
	http.Serve(ln, nil)
}

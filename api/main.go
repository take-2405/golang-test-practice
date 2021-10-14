package main

import (
	"net/http"
)


func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", helloHandler)
	server.ListenAndServe()
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
	return
}
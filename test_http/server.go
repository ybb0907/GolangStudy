package main

import (
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))

}

func main() {
	http.HandleFunc("/test", handle)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"
)

func add(w http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8008",nil)
}

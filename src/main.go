package main

import (
    "github.com/gorilla/mux"
    "net/http"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    http.ListenAndServe(":8080", router)
}
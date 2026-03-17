package routers

import (
    "github.com/gorilla/mux"
    "net/http"
)

func UserRouter(router *mux.Router) {
    userRouter := router.PathPrefix("/users").Subrouter()
    userRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("User list"))
    })
}
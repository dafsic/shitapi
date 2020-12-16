package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"shitapi/handlers"
)

//注册路由，设置分组，注册中间件
func mapRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/hello/{name}", handlers.Hello).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/api/info", handlers.Info).Methods(http.MethodGet, http.MethodOptions)

	return r
}

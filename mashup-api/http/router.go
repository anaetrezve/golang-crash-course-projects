package router

import "net/http"

type Router interface {
	Get(url string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}

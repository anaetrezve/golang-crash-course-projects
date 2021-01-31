package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")
}

func (*muxRouter) Serve(port string) {
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, muxDispatcher)
}

package router

import "net/http"

var Route Router

type Router interface {
	Handle() http.Handler
}

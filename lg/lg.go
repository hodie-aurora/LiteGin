package lg

import (
	"net/http"
)

// Engine implements the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of Engine
func New() *Engine {
	return &Engine{newRouter()}
}

// addRoute adds a new route to the router
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET adds a GET request handler to the router
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST adds a POST request handler to the router
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run starts the HTTP server on the specified address
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP handles incoming HTTP requests
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	engine.router.handle(ctx)
}

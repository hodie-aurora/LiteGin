package lg

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by lg
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine implements the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute adds a new route to the router
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Router is %s - %s", method, pattern)
	engine.router[key] = handler
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
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s %s\n", req.Method, req.URL)
	}
}

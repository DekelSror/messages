package main

import (
	"net/http"
)

type RouteHandler func(http.ResponseWriter, *http.Request, Server)

type Server struct {
	myCalls map[string][]string
	routes  map[string]RouteHandler
}

func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	url := r.URL.Path

	if server.routes[url] != nil {
		server.routes[url](w, r, server)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (server Server) Teach(path string, handler RouteHandler) {
	if server.routes[path] == nil {
		server.routes[path] = handler
	}
}

func MakeServer() *Server {
	server := new(Server)
	server.myCalls = make(map[string][]string)
	server.routes = make(map[string]RouteHandler)

	return server
}

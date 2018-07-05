package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	log.Printf("listening on TCP %s", addr)

	// root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("root handler - not found: %s", r.URL.Path)
		http.Error(w, "hellowasm_server - Not found", 404)
	})

	registerStatic("/console/", "run/console")

	server := &http.Server{Addr: addr}
	if errServe := server.ListenAndServe(); errServe != nil {
		log.Printf("serve: %v", errServe)
	}
	log.Printf("exiting")
}

type staticHandler struct {
	innerHandler http.Handler
}

func (handler staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("staticHandler.ServeHTTP url=%s from=%s", r.URL.Path, r.RemoteAddr)
	handler.innerHandler.ServeHTTP(w, r)
}

func registerStatic(path, dir string) {
	log.Printf("mapping www path %s to directory %s", path, dir)
	http.Handle(path, staticHandler{http.StripPrefix(path, http.FileServer(http.Dir(dir)))})
}

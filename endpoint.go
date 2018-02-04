package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type endpoint struct {
	listener   *net.TCPListener
	runPlugins func() []result
}

func (e *endpoint) start() {
	mux := http.NewServeMux()
	mux.Handle("/results", e)

	srv := http.Server{Addr: e.listener.Addr().String(), Handler: mux}
	log.Println("Started /rungopherrun server!")

	if err := srv.Serve(e.listener); err != http.ErrServerClosed {
		log.Printf("endpoint terminated due to unexpected error")
	}
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(e.runPlugins())
	if err != nil {
		log.Printf("Failed to marshal response")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(body)
}

func newEndpoint(port int, runPlugins func() []result) *endpoint {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0").To4(),
		Port: port,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &endpoint{l, runPlugins}
}

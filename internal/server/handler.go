package server

import (
	"net/http"
)

func Handler(pool *ServerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if peer := pool.GetNextServer(); peer != nil {
			peer.ReverseProxy.handler(w, r)
			return
		}
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
	}
}

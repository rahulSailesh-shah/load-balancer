package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/rahulSailesh-shah/load_balancer/internal/configs"
)

func Run(addr string) error {
	cfg, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	pool := NewPool()
	for _, be := range cfg.Backends {
		u, err := url.Parse(be.Destination_URL)
		if err != nil {
			return fmt.Errorf("parse backend URL %q: %w", be.Destination_URL, err)
		}

		proxy := NewProxy(u)
		proxy.reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("[%s] proxy error: %v\n", u.Host, err)

			// TODO: retry mechanism with the same server

			pool.MarkServer(u, false)

			// TODO: track the number of attempts and retry with other servers
		}

		pool.Add(NewServer(u, proxy))
		log.Printf("Configured backend: %s\n", u.Host)
	}

	StartHealthCheck(pool, 2*time.Minute)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler(pool))

	log.Printf("Starting LB on %s\n", addr)
	return http.ListenAndServe(addr, mux)
}

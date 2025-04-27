package main

import (
	"log"

	"github.com/rahulSailesh-shah/load_balancer/internal/server"
)

func main() {
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}

// package server

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/url"
// 	"time"

// 	"github.com/rahulSailesh-shah/load_balancer/internal/configs"
// )

// type Server struct {
// 	URL          *url.URL
// 	Alive        bool
// 	ReverseProxy *Proxy
// }

// func (s *Server) IsAlive(u *url.URL) bool {
// 	return s.Alive
// }

// func (s *Server) SetAlive(alive bool) {
// 	s.Alive = alive
// }

// type ServerPool struct {
// 	Servers []*Server
// 	Current int
// }

// func (s *ServerPool) AddBackend(server *Server) {
// 	s.Servers = append(s.Servers, server)
// }

// func (s *ServerPool) GetNextIndex() int {
// 	return (s.Current + 1) % len(s.Servers)
// }

// func (s *ServerPool) GetNextServer() *Server {
// 	// TODO: check no. of attempts for the same request

// 	next := s.GetNextIndex()
// 	l := len(s.Servers) + next
// 	for i := next; i < l; i++ {
// 		idx := i % len(s.Servers)
// 		if s.Servers[idx].IsAlive(s.Servers[idx].URL) {
// 			s.Current = idx
// 			return s.Servers[idx]
// 		}
// 	}

// 	return nil
// }

// func (s *ServerPool) MarkServer(serverUrl *url.URL, status bool) {
// 	for _, server := range s.Servers {
// 		if server.URL.String() == serverUrl.String() {
// 			server.SetAlive(status)
// 		}
// 	}
// }

// func (s *ServerPool) HealthCheck() {
// 	for _, s := range s.Servers {
// 		status := "up"
// 		alive := isBackendAlive(s.URL)
// 		s.SetAlive(alive)
// 		if !alive {
// 			status = "down"
// 		}
// 		log.Printf("%s [%s]\n", s.URL, status)
// 	}
// }

// func lb(w http.ResponseWriter, r *http.Request) {
// 	// TODO: validate the number of attempts for the same request

// 	peer := serverPool.GetNextServer()
// 	if peer != nil {
// 		peer.ReverseProxy.handler(w, r)
// 		return
// 	}
// 	http.Error(w, "Service not available", http.StatusServiceUnavailable)
// }

// func healthCheck() {
// 	t := time.NewTicker(time.Minute * 2)
// 	for {
// 		select {
// 		case <-t.C:
// 			log.Println("Starting health check...")
// 			serverPool.HealthCheck()
// 			log.Println("Health check completed")
// 		}
// 	}
// }

// func isBackendAlive(u *url.URL) bool {
// 	timeout := 2 * time.Second
// 	conn, err := net.DialTimeout("tcp", u.Host, timeout)
// 	if err != nil {
// 		log.Println("Site unreachable, error: ", err)
// 		return false
// 	}
// 	defer conn.Close()
// 	return true
// }

// var serverPool = &ServerPool{
// 	Servers: make([]*Server, 0),
// 	Current: 0,
// }

// // Run starts server and listens on defined port
// func Run() error {
// 	config, err := configs.NewConfiguration()

// 	fmt.Println(config)

// 	if err != nil {
// 		return fmt.Errorf("could not load configuration: %v", err)
// 	}

// 	for _, resource := range config.Backends {
// 		url, _ := url.Parse(resource.Destination_URL)
// 		proxy := NewProxy(url)

// 		// Error handler for the proxy
// 		proxy.reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
// 			log.Printf("[%s] %s\n", url.Host, err.Error())

// 			// TODO: retry mechanism with the same server

// 			serverPool.MarkServer(url, false)
// 			// TODO: track the number of attempts

// 		}

// 		//TODO: Add the server to the server pool
// 		serverPool.AddBackend(&Server{
// 			URL:          url,
// 			Alive:        true,
// 			ReverseProxy: proxy,
// 		})
// 		log.Printf("Configured server: %s\n", url.Host)
// 	}

// 	// create http server
// 	server := http.Server{
// 		Addr:    ":8080",
// 		Handler: http.HandlerFunc(lb),
// 	}

// 	// start health checking
// 	go healthCheck()

// 	fmt.Printf("Server started at %s\n", server.Addr)
// 	err = server.ListenAndServe()

// 	return err
// }

package server

import (
	"net/url"
)

// ServerPool holds all backends and the pointer for round-robin.
type ServerPool struct {
	servers []*Server
	current int
}

// NewPool initializes an empty pool.
func NewPool() *ServerPool {
	return &ServerPool{
		servers: []*Server{},
		current: -1,
	}
}

func (p *ServerPool) Add(s *Server) {
	p.servers = append(p.servers, s)
}

func (s *ServerPool) GetNextIndex() int {
	return (s.current + 1) % len(s.servers)
}

func (s *ServerPool) GetNextServer() *Server {
	// TODO: check no. of attempts for the same request

	next := s.GetNextIndex()
	l := len(s.servers) + next
	for i := next; i < l; i++ {
		idx := i % len(s.servers)
		server := s.servers[idx]
		if server.IsAlive() {
			s.current = idx
			return s.servers[idx]
		}
	}

	return nil
}

func (s *ServerPool) MarkServer(serverUrl *url.URL, status bool) {
	for _, server := range s.servers {
		if server.URL.String() == serverUrl.String() {
			server.SetAlive(status)
		}
	}
}

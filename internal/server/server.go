package server

import "net/url"

// Server represents one backend.
type Server struct {
	URL          *url.URL
	alive        bool
	ReverseProxy *Proxy
}

func NewServer(u *url.URL, proxy *Proxy) *Server {
	return &Server{
		URL:          u,
		alive:        true,
		ReverseProxy: proxy,
	}
}

func (s *Server) IsAlive() bool {
	return s.alive
}

func (s *Server) SetAlive(alive bool) {
	s.alive = alive
}

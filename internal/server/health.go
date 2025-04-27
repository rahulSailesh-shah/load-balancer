package server

import (
	"log"
	"net"
	"net/url"
	"time"
)

func StartHealthCheck(pool *ServerPool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			log.Println("Health check: starting")
			pool.CheckAll()
			log.Println("Health check: completed")
		}
	}()
}

func (p *ServerPool) CheckAll() {
	for _, srv := range p.servers {
		alive := isAlive(srv.URL)
		srv.SetAlive(alive)
		status := "down"
		if alive {
			status = "up"
		}
		log.Printf("%s [%s]\n", srv.URL.Host, status)
	}
}

func isAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

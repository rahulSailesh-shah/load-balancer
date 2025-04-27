package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type Proxy struct {
	URL          *url.URL
	reverseProxy *httputil.ReverseProxy
	handler      http.HandlerFunc
}

func NewProxy(u *url.URL) *Proxy {
	proxy := httputil.NewSingleHostReverseProxy(u)

	handler := ProxyRequestHandler(proxy, u, u.Host)

	return &Proxy{
		URL:          u,
		reverseProxy: proxy,
		handler:      handler,
	}
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Update the headers to allow for SSL redirection
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host

		//trim reverseProxyRoutePrefix
		path := r.URL.Path
		r.URL.Path = strings.TrimLeft(path, endpoint)

		// Note that ServeHttp is non blocking and uses a go routine under the hood
		fmt.Printf("[ PROXY ] Redirecting request to %s at %s\n", r.URL, time.Now().UTC())
		proxy.ServeHTTP(w, r)
	}
}

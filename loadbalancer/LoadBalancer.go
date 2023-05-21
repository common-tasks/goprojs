package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type ServerPool struct {
	servers []*url.URL
	current int
}

func (s *ServerPool) getNextServer() url.URL {
	url := s.servers[s.current]
	s.current = (s.current + 1) % (len(s.servers))
	return *url
}
func (s *ServerPool) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got request at load balancer")
	server := s.getNextServer()
	fmt.Println("served by ", server)
	proxy := httputil.NewSingleHostReverseProxy(&server)
	proxy.ServeHTTP(w, r)
}
func parseURL(s string) *url.URL {
	url, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return url
}
func LoadBalance() {
	pool := &ServerPool{
		servers: []*url.URL{
			parseURL("http://localhost:8081"),
			parseURL("http://localhost:8082"),
			parseURL("http://localhost:8083"),
			parseURL("http://localhost:8083"),
		},
	}
	http.HandleFunc("/", pool.handler)
	http.ListenAndServe(":8080", nil)
}
func RouteRequest() {
	fmt.Println("received request at ", time.Now())
}

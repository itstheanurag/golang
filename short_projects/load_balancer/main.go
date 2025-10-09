package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type simpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}


func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)

	handleErr(err)

	return  &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}


func handleErr(err error) {
	if err != nil {
		fmt.Printf("something went wrong, an error occureed, details: %v", err)
		os.Exit(1)
	}
}

type Server interface {
	Address() string
	IsAlive() bool
	Server(rw http.ResponseWriter, r *http.Request)
}

type LoadBalancer struct {
	port string
	roundRobinCount int
	servers []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
   return &LoadBalancer {
	port: port,
	roundRobinCount: 0,
	servers: servers,
   }
}


func (lb *LoadBalancer) getNextAvailableServer() Server {
    serverCount := len(lb.servers)
    for i := 0; i < serverCount; i++ {
        idx := (lb.roundRobinCount + i) % serverCount
        if lb.servers[idx].IsAlive() {
            lb.roundRobinCount = (idx + 1) % serverCount
            return lb.servers[idx]
        }
    }
    return nil
}

func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, r *http.Request) {
    server := lb.getNextAvailableServer()
    if server == nil {
        http.Error(rw, "No servers available", http.StatusServiceUnavailable)
        return
    }
    server.Server(rw, r)
}

func (s *simpleServer) Server(rw http.ResponseWriter, req *http.Request) {
    s.proxy.ServeHTTP(rw, req)
}


func  (s *simpleServer) Address() string { return s.addr} 

func (s *simpleServer) IsAlive() bool {
	return true
}



func main() {
	fmt.Println("The minimal load balancer written in the golang, a short fun learning project")
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.google.com"),
		newSimpleServer("https://www.amazon.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func (rw http.ResponseWriter, req *http.Request) {
		lb.serverProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	http.ListenAndServe(":" + *&lb.port, nil)
}
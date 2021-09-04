package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"

	"golang.org/x/net/proxy"
)

type ProxyHandler struct {
	Target *url.URL
	Proxy  *httputil.ReverseProxy
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s -> %s\n", r.URL, path.Join(p.Target.String(), r.URL.String()))
	// set host header to target host
	r.Host = p.Target.Host
	var PTransport = &http.Transport{
		Dial: proxy.FromEnvironment().Dial,
	}
	p.Proxy.Transport = PTransport
	p.Proxy.ServeHTTP(w, r)
}

func main() {
	targetURL := flag.String("targetURL", "", "target url to proxy to")
	flag.Parse()
	if *targetURL == "" {
		log.Fatal("targetURL is required")
	}
	remote, err := url.Parse(*targetURL)
	if err != nil {
		panic(err)
	}

	log.Printf("Proxying to: %s\n", *targetURL)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.Handle("/", &ProxyHandler{remote, proxy})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

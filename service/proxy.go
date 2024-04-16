package service

import (
	"go-proxy/lib"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type GoProxy struct {
	Port int
}

// initProxy
func NewProxy() *GoProxy {
	config := lib.Config()
	var port int
	if config.Port == 0 {
		port = 8080 // default port
	} else {
		port = config.Port
	}
	return &GoProxy{
		Port: port,
	}
}

func (g *GoProxy) ProxyAddressRequest(w http.ResponseWriter, r *http.Request) {
	// set target url
	proxyAddress := g.GetAddressByUrl(r.URL.Path)
	targetURL, err := url.Parse(proxyAddress)
	if err != nil {
		lib.ProxyLogger().Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// replace request url
	url_path := g.SubPrefix(r.URL.Path, proxyAddress)
	r.URL.Path = url_path
	lib.ProxyLogger().Info("proxy address: " + proxyAddress + ", url: " + url_path)

	// proxy request
	proxy.ServeHTTP(w, r)

}

func (g *GoProxy) GetAddressByUrl(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 2 {
		return "http://" + parts[2]
	}
	return ""
}

func (g *GoProxy) SubPrefix(url string, proxyAddress string) string {
	prefix := "/proxy/" + proxyAddress[len("http://"):]
	if len(url) >= len(prefix) {
		url = url[len(prefix):]
	}
	return url
}

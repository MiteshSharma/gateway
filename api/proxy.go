package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/MiteshSharma/gateway/util"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Registry interface {
	GetService(string) (string, error)
}

type ServiceProxyHandler struct {
	registry Registry
}

func NewServiceProxyHandler(_registry Registry) *ServiceProxyHandler {
	h := &ServiceProxyHandler{
		registry: _registry,
	}
	return h
}

func InitProxy(router *mux.Router) {
	registry := NewLocalRegistry()
	serviceProxy := NewServiceProxyHandler(registry)
	router.PathPrefix("/").HandlerFunc(serviceProxy.proxyHandler())
}

func (h *ServiceProxyHandler) proxyHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteUrl, err := h.registry.GetService(r.Host)
		if err != nil {
			utils.Logger.Error("Error fetching service url ", zap.Error(err))
			return
		}

		remote, err := url.Parse(remoteUrl)
		if err != nil {
			utils.Logger.Error("Error parsing service url ", zap.Error(err))
			return
		}
		path := "/*catchall"
		reverseProxy := httputil.NewSingleHostReverseProxy(remote)
		reverseProxy.Director = func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", remote.Host)
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			wildcardIndex := strings.IndexAny(path, "*")
			proxyPath := singleJoiningSlash(remote.Path, req.URL.Path[wildcardIndex:])
			if strings.HasSuffix(proxyPath, "/") && len(proxyPath) > 1 {
				proxyPath = proxyPath[:len(proxyPath)-1]
			}
			req.URL.Path = proxyPath
		}
		reverseProxy.ServeHTTP(w, r)
	}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

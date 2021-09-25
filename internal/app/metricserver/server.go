package metricserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	srv := &server{
		router: mux.NewRouter(),
	}

	srv.configureRouter()

	return srv
}

func (srv *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(rw, r)
}

func (srv *server) configureRouter() {
	srv.router.Handle("/metrics", promhttp.Handler()).Methods("GET")
}

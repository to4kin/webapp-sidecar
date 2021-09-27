package metricserver

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	countFiles = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "files_count",
		Help: "The number of files in upload folder",
	})

	totalSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "files_size_total",
		Help: "The total size in bytes of files in the upload folder",
	})
)

type server struct {
	router *mux.Router
}

func newServer(config *Config) *server {
	srv := &server{
		router: mux.NewRouter(),
	}

	srv.registerMetrics(config)
	srv.configureRouter(config)

	return srv
}

func (srv *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(rw, r)
}

func (srv *server) configureRouter(config *Config) {
	srv.router.Handle(config.MetricsPath, promhttp.Handler()).Methods("GET")
}

func (srv *server) registerMetrics(config *Config) {
	go func() {
		for {
			files, _ := ioutil.ReadDir(config.UploadFolder)
			countFiles.Set(float64(len(files)))

			var total int64 = 0
			for _, file := range files {
				total += file.Size()
			}
			totalSize.Set(float64(total))

			time.Sleep(time.Duration(config.CheckFolderInterval) * time.Second)
		}
	}()
}

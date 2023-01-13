package metrics

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

type Metrics interface {
	StartPrometheusHandler()
	WaitSync()
}

type MetricData struct {
	Uri, Port string
	waitGroup *sync.WaitGroup
}

func NewPrometheusInstance(uri, port string) Metrics {
	waitGroup := &sync.WaitGroup{}
	return &MetricData{Uri: uri, Port: port, waitGroup: waitGroup}
}

func (data *MetricData) WaitSync() {
	data.waitGroup.Wait()
}

func (data *MetricData) StartPrometheusHandler() {

	data.waitGroup.Add(1)

	go func() {

		http.Handle(data.Uri, promhttp.Handler())
		err := http.ListenAndServe(data.Port, nil)

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server: %s\n", err)
		}

		data.waitGroup.Done()
	}()

	fmt.Printf("Started prometheus handler at localhost%s%s\n", data.Port, data.Uri)

}

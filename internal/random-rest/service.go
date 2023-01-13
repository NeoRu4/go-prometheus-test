package random_rest

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go-prometheus/internal/metrics"
	"io"
	"log"
	"net/http"
	"time"
)

type Queries interface {
	QueryOne(postNum int)
}

type QueriesData struct {
	posts         []Post
	metricLatency *prometheus.HistogramVec
	metricCount   prometheus.Counter
}

func NewQueryService() Queries {

	service := &QueriesData{}
	service.Init()

	return service
}

func (data *QueriesData) Init() {
	data.metricLatency = metrics.LatencyOfQuerySummary("func_latency")
	data.metricCount = metrics.CountOfSuccess("success_func")
}

func (data *QueriesData) QueryOne(postNum int) {
	startedAt := time.Now()
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postNum))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	var res Post
	json.Unmarshal(b, &res)

	data.posts = append(data.posts, res)
	elapsed := time.Since(startedAt).Seconds()

	data.metricLatency.WithLabelValues(fmt.Sprintf("jsonplaceholder/post/%d", postNum)).Observe(elapsed)
	data.metricCount.Inc()
}

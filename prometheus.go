package utils

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	// 啟動 Web 服務
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("啟動失敗")
		}
		log.Info("監控端口,端口為：" + strconv.Itoa(port))
	}()
}

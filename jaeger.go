package utils

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// 建立jaeger鏈路追蹤實例
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		// 服務名稱
		ServiceName: serviceName,
		// 採樣的類型
		// Type				Param		说明
		// "const"			0或1		採樣器對所有tracer做出相同決定；全部採樣，或全部不採樣
		// "probabilistic"	0.0~1.0		採樣器隨機採樣，Param 為採樣機率
		// "ratelimiting"	N			採樣器以固定的速率採樣，Param=2.0，則限制每秒2條
		// "remote"			無			採樣器詢問 Jaeger 代理，獲取當前服務中使用的適當採樣策略
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		// 報告方式
		// 以下為 Reporter構造
		// type ReporterConfig struct {
		// 	QueueSize                  int `yaml:"queueSize"`
		// 	BufferFlushInterval        time.Duration
		// 	LogSpans                   bool   `yaml:"logSpans"`
		// 	LocalAgentHostPort         string `yaml:"localAgentHostPort"`
		// 	DisableAttemptReconnecting bool   `yaml:"disableAttemptReconnecting"`
		// 	AttemptReconnectInterval   time.Duration
		// 	CollectorEndpoint          string            `yaml:"collectorEndpoint"`
		// 	User                       string            `yaml:"user"`
		// 	Password                   string            `yaml:"password"`
		// 	HTTPHeaders                map[string]string `yaml:"http_headers"`
		// }

		// 常用配置
		// QUEUESIZE：設置 Quere 大小，儲存採樣 span 資訊，Queue滿了就一次發送到 jaeger 後端；defaultQueueSize 默認 100；
		// BufferFlushInterval：強制清除、推送 Queue 時間，對於流量不高的程式，Queue可能長時間不满，遇到超時狀況可以自動推送一次。默認為 1 秒。
		// LogSpans：是否把 Log 也推送，span 中可以攜帶一些日誌資訊。
		// LocalAgentHostPort：要推送到的 Jaeger agent，默認端口 6831，是 Jaeger 接收壓縮格式的 thrift 協議的資料端口。
		// CollectorEndpoint：要推送到的 Jaeger Collector，用 Collector 就不用 agent 了。
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	// 實例化 Tracer
	tracer, closer, err := cfg.NewTracer()
	return tracer, closer, err
}

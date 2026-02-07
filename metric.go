package prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Metric struct {
	Name      string
	Collector prometheus.Collector
}

func getMetricByName(ctx core.Ctx, name string) prometheus.Collector {
	config, ok := ctx.Ref(PROMPT).(*Config)
	if !ok || config == nil {
		return nil
	}

	for _, metric := range config.Metrics {
		if metric.Name == name {
			return metric.Collector
		}
	}

	return nil
}

func InjectCounter(ctx core.Ctx, name string) prometheus.Counter {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(prometheus.Counter)
	if !ok {
		return nil
	}

	return metric
}

func InjectCounterVec(ctx core.Ctx, name string) *prometheus.CounterVec {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(*prometheus.CounterVec)
	if !ok {
		return nil
	}

	return metric
}

func InjectGauge(ctx core.Ctx, name string) prometheus.Gauge {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(prometheus.Gauge)
	if !ok {
		return nil
	}

	return metric
}

func InjectGaugeVec(ctx core.Ctx, name string) *prometheus.GaugeVec {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(*prometheus.GaugeVec)
	if !ok {
		return nil
	}

	return metric
}

func InjectHistogram(ctx core.Ctx, name string) prometheus.Histogram {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(prometheus.Histogram)
	if !ok {
		return nil
	}

	return metric
}

func InjectHistogramVec(ctx core.Ctx, name string) *prometheus.HistogramVec {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(*prometheus.HistogramVec)
	if !ok {
		return nil
	}

	return metric
}

func InjectSummary(ctx core.Ctx, name string) prometheus.Summary {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(prometheus.Summary)
	if !ok {
		return nil
	}

	return metric
}

func InjectSummaryVec(ctx core.Ctx, name string) *prometheus.SummaryVec {
	collector := getMetricByName(ctx, name)
	if collector == nil {
		return nil
	}

	metric, ok := collector.(*prometheus.SummaryVec)
	if !ok {
		return nil
	}

	return metric
}

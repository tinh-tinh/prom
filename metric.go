package prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Metric struct {
	Name      string
	Collector prometheus.Collector
}

func InjectCounter(ctx core.Ctx, name string) prometheus.Counter {
	metric, ok := ctx.Ref(GetMetricName(name)).(prometheus.Counter)
	if !ok {
		return nil
	}

	return metric
}

func InjectCounterVec(ctx core.Ctx, name string) *prometheus.CounterVec {
	metric, ok := ctx.Ref(GetMetricName(name)).(*prometheus.CounterVec)
	if !ok {
		return nil
	}

	return metric
}

func InjectGauge(ctx core.Ctx, name string) prometheus.Gauge {
	metric, ok := ctx.Ref(GetMetricName(name)).(prometheus.Gauge)
	if !ok {
		return nil
	}

	return metric
}

func InjectGaugeVec(ctx core.Ctx, name string) *prometheus.GaugeVec {
	metric, ok := ctx.Ref(GetMetricName(name)).(*prometheus.GaugeVec)
	if !ok {
		return nil
	}

	return metric
}

func InjectHistogram(ctx core.Ctx, name string) prometheus.Histogram {
	metric, ok := ctx.Ref(GetMetricName(name)).(prometheus.Histogram)
	if !ok {
		return nil
	}

	return metric
}

func InjectHistogramVec(ctx core.Ctx, name string) *prometheus.HistogramVec {
	metric, ok := ctx.Ref(GetMetricName(name)).(*prometheus.HistogramVec)
	if !ok {
		return nil
	}

	return metric
}

func InjectSummary(ctx core.Ctx, name string) prometheus.Summary {
	metric, ok := ctx.Ref(GetMetricName(name)).(prometheus.Summary)
	if !ok {
		return nil
	}

	return metric
}

func InjectSummaryVec(ctx core.Ctx, name string) *prometheus.SummaryVec {
	metric, ok := ctx.Ref(GetMetricName(name)).(*prometheus.SummaryVec)
	if !ok {
		return nil
	}

	return metric
}

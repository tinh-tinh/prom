package prom

import (
	"errors"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const PROMPT core.Provide = "PROMPT"

type Config struct {
	Metrics []Metric
	Opt     promhttp.HandlerOpts
}

func promHandler(module core.Module) core.Controller {
	ctrl := module.NewController("/metrics")
	pro, ok := module.Ref(PROMPT).(*Config)
	if !ok || pro == nil {
		panic(errors.New("prometheus not config"))
	}

	if len(pro.Metrics) > 0 {
		reg := prometheus.NewRegistry()
		for _, metric := range pro.Metrics {
			err := reg.Register(metric.Collector)
			if err != nil {
				log.Print(err)
			}
		}
		handler := promhttp.HandlerFor(
			reg,
			pro.Opt,
		)

		ctrl.Handler("", handler)
	} else {
		ctrl.Handler("", promhttp.Handler())
	}

	return ctrl
}

func ForRoot(config *Config) core.Modules {
	return func(module core.Module) core.Module {
		promModule := module.New(core.NewModuleOptions{})

		promModule.NewProvider(core.ProviderOptions{
			Name:  PROMPT,
			Value: config,
		})
		promModule.Export(PROMPT)

		promModule.Controllers(promHandler)

		return promModule
	}
}

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var fooMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "foo_metric", Help: "foo help"}, []string {"hostname", "function"},)
var barMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "bar_metric", Help: "bar help"}, []string {"hostname", "function"},)

func updateMetrics() {
	go func() {
		for {
			fooMetric.WithLabelValues("ccsw1.sqx.colo", "cc").Add(10)
			barMetric.WithLabelValues("ocsw1.dft.colo", "oc").Add(5)
			time.Sleep(2 * time.Second)
		}
	}()
}


func init() {
	prometheus.MustRegister(fooMetric)
	prometheus.MustRegister(barMetric)

	updateMetrics()
}





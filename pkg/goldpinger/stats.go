// Copyright 2018 Bloomberg Finance L.P.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goldpinger

import (
	"context"
	"time"

	"github.com/bloomberg/goldpinger/v3/pkg/models"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

var (
	goldpingerStatsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goldpinger_stats_total",
			Help: "Statistics of calls made in goldpinger instances",
		},
		[]string{
			"goldpinger_instance",
			"group",
			"action",
		},
	)

	goldpingerNodesHealthGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goldpinger_nodes_health_total",
			Help: "Number of nodes seen as healthy/unhealthy from this instance's POV",
		},
		[]string{
			"goldpinger_instance",
			"status",
		},
	)

	goldpingerClusterHealthGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goldpinger_cluster_health_total",
			Help: "1 if all check pass, 0 otherwise",
		},
		[]string{
			"goldpinger_instance",
		},
	)

	goldpingerResponseTimePeersHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "goldpinger_peers_response_time_s",
			Help:    "Histogram of response times from other hosts, when making peer calls",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30},
		},
		[]string{
			"goldpinger_instance",
			"call_type",
			"host_ip",
			"pod_ip",
			"hostName",
		},
	)

	goldpingerResponseTimeKubernetesHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "goldpinger_kube_master_response_time_s",
			Help:    "Histogram of response times from kubernetes API server, when listing other instances",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30},
		},
		[]string{
			"goldpinger_instance",
		},
	)

	goldpingerErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goldpinger_errors_total",
			Help: "Statistics of errors per instance",
		},
		[]string{
			"goldpinger_instance",
			"type",
		},
	)
	goldpingerDnsErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goldpinger_dns_errors_total",
			Help: "Statistics of DNS errors per instance",
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerTelnetErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goldpinger_telnet_errors_total",
			Help: "Statistics of telnet errors per instance",
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerELSErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goldpinger_els_errors_total",
			Help: "Statistics of els errors per instance",
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerTelnetHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "goldpinger_telnet_response_time_s",
			Help:    "Histogram of response times from telnet services",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30},
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerELSHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "goldpinger_els_response_time_s",
			Help:    "Histogram of response times from els services",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 30},
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerTelnetConnectivityGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goldpinger_telnet_connectivity",
			Help: "1 if connectivity pass, 0 otherwise",
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerELSConnectivityGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goldpinger_els_connectivity",
			Help: "1 if connectivity pass, 0 otherwise",
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	goldpingerPeerConnectivityGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goldpinger_peer_connectivity",
			Help: "Peer connectivity status",
		},
		[]string{
			"goldpinger_instance",
			"host",
		},
	)

	bootTime = time.Now()
)

func init() {
	prometheus.MustRegister(goldpingerStatsCounter)
	prometheus.MustRegister(goldpingerNodesHealthGauge)
	prometheus.MustRegister(goldpingerClusterHealthGauge)
	prometheus.MustRegister(goldpingerResponseTimePeersHistogram)
	prometheus.MustRegister(goldpingerResponseTimeKubernetesHistogram)
	prometheus.MustRegister(goldpingerErrorsCounter)
	prometheus.MustRegister(goldpingerDnsErrorsCounter)
	prometheus.MustRegister(goldpingerELSErrorsCounter)
	prometheus.MustRegister(goldpingerTelnetErrorsCounter)
	prometheus.MustRegister(goldpingerELSHistogram)
	prometheus.MustRegister(goldpingerTelnetHistogram)
	prometheus.MustRegister(goldpingerTelnetConnectivityGauge)
	prometheus.MustRegister(goldpingerELSConnectivityGauge)
	prometheus.MustRegister(goldpingerPeerConnectivityGauge)
	zap.L().Info("Metrics setup - see /metrics")
}

func GetStats(ctx context.Context) *models.PingResults {
	// GetStats no longer populates the received and made calls - use metrics for that instead
	return &models.PingResults{
		BootTime: strfmt.DateTime(bootTime),
	}
}

// counts various calls received and made
func CountCall(group string, call string) {
	goldpingerStatsCounter.WithLabelValues(
		GoldpingerConfig.Hostname,
		group,
		call,
	).Inc()
}

// counts healthy and unhealthy nodes
func CountHealthyUnhealthyNodes(healthy, unhealthy float64) {
	goldpingerNodesHealthGauge.WithLabelValues(
		GoldpingerConfig.Hostname,
		"healthy",
	).Set(healthy)
	goldpingerNodesHealthGauge.WithLabelValues(
		GoldpingerConfig.Hostname,
		"unhealthy",
	).Set(unhealthy)
}

// SetClusterHealth sets the cluster health gauge to 1 (healthy) or 0 (unhealthy)
func SetClusterHealth(healthy bool) {
	value := 1.0
	if !healthy {
		value = 0
	}
	goldpingerClusterHealthGauge.WithLabelValues(
		GoldpingerConfig.Hostname,
	).Set(value)
}

// counts instances of various errors
func CountError(errorType string) {
	goldpingerErrorsCounter.WithLabelValues(
		GoldpingerConfig.Hostname,
		errorType,
	).Inc()
}

// counts instances of dns errors
func CountDnsError(host string) {
	goldpingerDnsErrorsCounter.WithLabelValues(
		GoldpingerConfig.Hostname,
		host,
	).Inc()
}

// counts instances of telnet errors
func CountTelnetError(host string) {
	goldpingerTelnetErrorsCounter.WithLabelValues(
		GoldpingerConfig.Hostname,
		host,
	).Inc()
}

// counts instances of ELS errors
func CountElsError(host string) {
	goldpingerELSErrorsCounter.WithLabelValues(
		GoldpingerConfig.Hostname,
		host,
	).Inc()
}

// returns a timer for easy observing of the durations of calls to kubernetes API
func GetLabeledKubernetesCallsTimer() *prometheus.Timer {
	return prometheus.NewTimer(
		goldpingerResponseTimeKubernetesHistogram.WithLabelValues(
			GoldpingerConfig.Hostname,
		),
	)
}

// returns a timer for easy observing of the duration of calls to peers
func GetLabeledPeersCallsTimer(callType, hostIP, podIP string) *prometheus.Timer {
	return prometheus.NewTimer(
		goldpingerResponseTimePeersHistogram.WithLabelValues(
			GoldpingerConfig.Hostname,
			callType,
			hostIP,
			podIP,
		),
	)
}

// returns a timer for easy observing of the duration of calls to telnet peers
func GetLabeledTelnetCallsTimer(host string) *prometheus.Timer {
	return prometheus.NewTimer(
		goldpingerTelnetHistogram.WithLabelValues(
			GoldpingerConfig.Hostname,
			host,
		),
	)
}

// returns a timer for easy observing of the duration of calls to ELS peers
func GetLabeledELSCallsTimer(host string) *prometheus.Timer {
	return prometheus.NewTimer(
		goldpingerELSHistogram.WithLabelValues(
			GoldpingerConfig.Hostname,
			host,
		),
	)
}

// SetTelnetConnectivityStatus sets the telnet endpoint connectivity gauge to 1 (healthy) or 0 (unhealthy)
func SetTelnetConnectivityStatus(healthy bool, host string) {
	value := 1.0
	if !healthy {
		value = 0
	}
	goldpingerTelnetConnectivityGauge.WithLabelValues(
		GoldpingerConfig.Hostname,
		host,
	).Set(value)
}

// SetELSConnectivityStatus sets the ELS endpoint connectivity gauge to 1 (healthy) or 0 (unhealthy)
func SetELSConnectivityStatus(healthy bool, host string) {
	value := 1.0
	if !healthy {
		value = 0
	}
	goldpingerELSConnectivityGauge.WithLabelValues(
		GoldpingerConfig.Hostname,
		host,
	).Set(value)
}

// SetPeerConnectivityStatus sets the Peer endpoint connectivity gauge to 1 (healthy) or 0 (unhealthy)
func SetPeerConnectivityStatus(healthy bool, host string) {
	value := 1.0
	if !healthy {
		value = 0
	}
	goldpingerPeerConnectivityGauge.WithLabelValues(
		GoldpingerConfig.Hostname,
		host,
	).Set(value)
}

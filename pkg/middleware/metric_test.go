package middleware

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestSetGaugeValue(t *testing.T) {
	metric := &Metric{
		Type:        Gauge,
		Name:        "test_gauge",
		Description: "A test gauge metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{Name: "test_gauge", Help: "A test gauge metric"},
			[]string{"label1"},
		),
	}

	err := metric.SetGaugeValue([]string{"value1"}, 10.5)
	assert.NoError(t, err, "Setting gauge value should not produce an error")
}

func TestSetGaugeValueErrors(t *testing.T) {
	metric := &Metric{
		Type:        None,
		Name:        "test_none",
		Description: "A test none metric",
		Labels:      []string{"label1"},
	}

	err := metric.SetGaugeValue([]string{"value1"}, 10.5)
	assert.Error(t, err, "Setting gauge value for None type should produce an error")

	metric.Type = Counter
	err = metric.SetGaugeValue([]string{"value1"}, 10.5)
	assert.Error(t, err, "Setting gauge value for Counter type should produce an error")
}

func TestInc(t *testing.T) {
	counterMetric := &Metric{
		Type:        Counter,
		Name:        "test_counter",
		Description: "A test counter metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewCounterVec(
			prometheus.CounterOpts{Name: "test_counter", Help: "A test counter metric"},
			[]string{"label1"},
		),
	}

	err := counterMetric.Inc([]string{"value1"})
	assert.NoError(t, err, "Incrementing counter should not produce an error")

	gaugeMetric := &Metric{
		Type:        Gauge,
		Name:        "test_gauge",
		Description: "A test gauge metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{Name: "test_gauge", Help: "A test gauge metric"},
			[]string{"label1"},
		),
	}

	err = gaugeMetric.Inc([]string{"value1"})
	assert.NoError(t, err, "Incrementing gauge should not produce an error")
}

func TestIncErrors(t *testing.T) {
	metric := &Metric{
		Type:        None,
		Name:        "test_none",
		Description: "A test none metric",
		Labels:      []string{"label1"},
	}

	err := metric.Inc([]string{"value1"})
	assert.Error(t, err, "Incrementing None type should produce an error")

	metric.Type = Histogram
	err = metric.Inc([]string{"value1"})
	assert.Error(t, err, "Incrementing Histogram type should produce an error")
}

func TestAdd(t *testing.T) {
	counterMetric := &Metric{
		Type:        Counter,
		Name:        "test_counter",
		Description: "A test counter metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewCounterVec(
			prometheus.CounterOpts{Name: "test_counter", Help: "A test counter metric"},
			[]string{"label1"},
		),
	}

	err := counterMetric.Add([]string{"value1"}, 5.5)
	assert.NoError(t, err, "Adding to counter should not produce an error")

	gaugeMetric := &Metric{
		Type:        Gauge,
		Name:        "test_gauge",
		Description: "A test gauge metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{Name: "test_gauge", Help: "A test gauge metric"},
			[]string{"label1"},
		),
	}

	err = gaugeMetric.Add([]string{"value1"}, 5.5)
	assert.NoError(t, err, "Adding to gauge should not produce an error")
}

func TestAddErrors(t *testing.T) {
	metric := &Metric{
		Type:        None,
		Name:        "test_none",
		Description: "A test none metric",
		Labels:      []string{"label1"},
	}

	err := metric.Add([]string{"value1"}, 5.5)
	assert.Error(t, err, "Adding to None type should produce an error")

	metric.Type = Histogram
	err = metric.Add([]string{"value1"}, 5.5)
	assert.Error(t, err, "Adding to Histogram type should produce an error")
}

func TestObserve(t *testing.T) {
	histogramMetric := &Metric{
		Type:        Histogram,
		Name:        "test_histogram",
		Description: "A test histogram metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{Name: "test_histogram", Help: "A test histogram metric", Buckets: []float64{0.1, 0.5, 1.0}},
			[]string{"label1"},
		),
	}

	err := histogramMetric.Observe([]string{"value1"}, 0.7)
	assert.NoError(t, err, "Observing histogram should not produce an error")

	summaryMetric := &Metric{
		Type:        Summary,
		Name:        "test_summary",
		Description: "A test summary metric",
		Labels:      []string{"label1"},
		vec: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{Name: "test_summary", Help: "A test summary metric", Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01}},
			[]string{"label1"},
		),
	}

	err = summaryMetric.Observe([]string{"value1"}, 0.7)
	assert.NoError(t, err, "Observing summary should not produce an error")
}

func TestObserveErrors(t *testing.T) {
	metric := &Metric{
		Type:        None,
		Name:        "test_none",
		Description: "A test none metric",
		Labels:      []string{"label1"},
	}

	err := metric.Observe([]string{"value1"}, 0.7)
	assert.Error(t, err, "Observing None type should produce an error")

	metric.Type = Counter
	err = metric.Observe([]string{"value1"}, 0.7)
	assert.Error(t, err, "Observing Counter type should produce an error")
}

package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMonitor(t *testing.T) {
	// Ensure the GetMonitor function returns a singleton instance
	monitor1 := GetMonitor()
	monitor2 := GetMonitor()

	assert.Equal(t, monitor1, monitor2, "GetMonitor should return the same instance")
}

func TestSetMetricPath(t *testing.T) {
	monitor := GetMonitor()
	path := "/test/metrics"
	monitor.SetMetricPath(path)

	assert.Equal(t, path, monitor.metricPath, "Metric path should be set correctly")
}

func TestSetSlowTime(t *testing.T) {
	monitor := GetMonitor()
	slowTime := int32(10)
	monitor.SetSlowTime(slowTime)

	assert.Equal(t, slowTime, monitor.slowTime, "Slow time should be set correctly")
}

func TestSetDuration(t *testing.T) {
	monitor := GetMonitor()
	duration := []float64{0.5, 1.0, 2.5}
	monitor.SetDuration(duration)

	assert.Equal(t, duration, monitor.reqDuration, "Request duration should be set correctly")
}

func TestAddMetric(t *testing.T) {
	monitor := GetMonitor()

	// Test adding a Counter metric
	counterMetric := &Metric{
		Name:        "test_counter",
		Description: "A test counter metric",
		Type:        Counter,
		Labels:      []string{"label1"},
	}
	err := monitor.AddMetric(counterMetric)
	assert.NoError(t, err, "Adding a counter metric should not produce an error")
	assert.Contains(t, monitor.metrics, counterMetric.Name, "Counter metric should be added to monitor metrics")

	// Test adding a Gauge metric
	gaugeMetric := &Metric{
		Name:        "test_gauge",
		Description: "A test gauge metric",
		Type:        Gauge,
		Labels:      []string{"label1"},
	}
	err = monitor.AddMetric(gaugeMetric)
	assert.NoError(t, err, "Adding a gauge metric should not produce an error")
	assert.Contains(t, monitor.metrics, gaugeMetric.Name, "Gauge metric should be added to monitor metrics")

	// Test adding a Histogram metric without buckets (should produce an error)
	histogramMetric := &Metric{
		Name:        "test_histogram",
		Description: "A test histogram metric",
		Type:        Histogram,
		Labels:      []string{"label1"},
	}
	err = monitor.AddMetric(histogramMetric)
	assert.Error(t, err, "Adding a histogram metric without buckets should produce an error")

	// Test adding a Summary metric without objectives (should produce an error)
	summaryMetric := &Metric{
		Name:        "test_summary",
		Description: "A test summary metric",
		Type:        Summary,
		Labels:      []string{"label1"},
	}
	err = monitor.AddMetric(summaryMetric)
	assert.Error(t, err, "Adding a summary metric without objectives should produce an error")
}

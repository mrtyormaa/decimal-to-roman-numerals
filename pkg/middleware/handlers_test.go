package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func setupRouter(m *Monitor) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	m.Use(r)
	return r
}

func setupRouterWithoutExpose(m *Monitor) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	m.UseWithoutExposingEndpoint(r)
	return r
}

func setupRouterWithExpose(m *Monitor) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	m.Expose(r)
	return r
}

func TestUse(t *testing.T) {
	monitor := GetMonitor()
	r := setupRouter(monitor)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/debug/metrics", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUseWithoutExposingEndpoint(t *testing.T) {
	monitor := GetMonitor()
	r := setupRouterWithoutExpose(monitor)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/debug/metrics", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestExpose(t *testing.T) {
	monitor := GetMonitor()
	r := setupRouterWithExpose(monitor)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/debug/metrics", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMonitorInterceptor(t *testing.T) {
	monitor := GetMonitor()
	r := setupRouter(monitor)

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())
}

func TestGinMetricHandle(t *testing.T) {
	monitor := GetMonitor()
	monitor.initGinMetrics()

	r := setupRouter(monitor)
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.String(http.StatusOK, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())

	metric := monitor.GetMetric(metricSlowRequest)
	counterVec, ok := metric.vec.(*prometheus.CounterVec)
	assert.True(t, ok, "Metric should be a CounterVec")

	counter, err := counterVec.GetMetricWithLabelValues("/test", "GET", "200")
	assert.NoError(t, err)
	assert.NotNil(t, counter)
}

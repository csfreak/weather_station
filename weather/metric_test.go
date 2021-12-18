package weather

import (
	"testing"
)

//These are dummy tests that only check for runtime errors, not functionality

func TestMetricInit(t *testing.T) {
	MetricInit()
}

func TestMetricsUpdate(t *testing.T) {
	UpdateMetrics(&test_weather_1)
}

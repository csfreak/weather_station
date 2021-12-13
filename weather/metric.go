package weather

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	temperatureMetricVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weatherstation_temperature_farenheit",
			Help: "Current Temperature",
		},
		[]string{
			"location",
			"station",
		},
	)
	humidityMetricVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weatherstation_humidity",
			Help: "Current Relative Humidity",
		},
		[]string{
			"location",
			"station",
		},
	)
	windSpeedMetricVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weatherstation_windspeed_mph",
			Help: "Current Wind Speed",
		},
		[]string{
			"type",
			"station",
		},
	)
	windDirectionMetricVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weatherstation_wind_direction_degrees",
			Help: "Current Wind Direction in Degrees",
		},
		[]string{
			"station",
		},
	)
	rainMetricVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weatherstation_rain_rate_inhr",
			Help: "Current Rain Rate in Inches per Hour",
		},
		[]string{
			"station",
		},
	)
	barometricPressureVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weatherstation_barometric_pressure_inhg",
			Help: "Current Barometric Presure in Inches of Mercury",
		},
		[]string{
			"type",
			"station",
		},
	)
)

func MetricInit() {
	prometheus.MustRegister(temperatureMetricVec)
	prometheus.MustRegister(humidityMetricVec)
	prometheus.MustRegister(windSpeedMetricVec)
	prometheus.MustRegister(windDirectionMetricVec)
	prometheus.MustRegister(rainMetricVec)
	prometheus.MustRegister(barometricPressureVec)
}

func UpdateMetrics(w *Weather) {
	weatherStation := fmt.Sprintf("%s-%s", w.Model, w.StationType)

	temperatureMetricVec.With(
		prometheus.Labels{
			"location": "indoor",
			"station":  weatherStation,
		}).Set(w.IndoorTempF)
	temperatureMetricVec.With(
		prometheus.Labels{
			"location": "outdoor",
			"station":  weatherStation,
		}).Set(w.OutdoorTempF)
	humidityMetricVec.With(
		prometheus.Labels{
			"location": "indoor",
			"station":  weatherStation,
		}).Set(w.IndoorHumidity)
	humidityMetricVec.With(
		prometheus.Labels{
			"location": "outdoor",
			"station":  weatherStation,
		}).Set(w.OutdoorHumidity)
	windSpeedMetricVec.With(
		prometheus.Labels{
			"type":    "sustained",
			"station": weatherStation,
		}).Set(w.WindspeedMPH)
	windSpeedMetricVec.With(
		prometheus.Labels{
			"type":    "gust",
			"station": weatherStation,
		}).Set(w.WindGustMPH)
	windDirectionMetricVec.With(
		prometheus.Labels{
			"station": weatherStation,
		}).Set(float64(w.WindDirectionDegree))
	rainMetricVec.With(
		prometheus.Labels{
			"station": weatherStation,
		}).Set(w.RainRateInch)
	barometricPressureVec.With(
		prometheus.Labels{
			"type":    "absolute",
			"station": weatherStation,
		}).Set(w.BarometricPressureABSIn)
	barometricPressureVec.With(
		prometheus.Labels{
			"type":    "relative",
			"station": weatherStation,
		}).Set(w.BarometricPressureRelIn)
}

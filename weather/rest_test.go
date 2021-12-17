package weather

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	rest_path      = "/v1/stations"
	test_weather_1 = Weather{
		StationType:             "test",
		Model:                   "test",
		Timestamp:               time.Now(),
		IndoorTempF:             1,
		IndoorHumidity:          1,
		OutdoorTempF:            1,
		OutdoorHumidity:         1,
		WindspeedMPH:            1,
		WindGustMPH:             1,
		WindDirectionDegree:     1,
		RainRateInch:            1,
		BarometricPressureABSIn: 1,
		BarometricPressureRelIn: 1,
	}
	test_weather_2 = Weather{
		StationType:             "test",
		Model:                   "test",
		Timestamp:               time.Now(),
		IndoorTempF:             2,
		IndoorHumidity:          2,
		OutdoorTempF:            2,
		OutdoorHumidity:         2,
		WindspeedMPH:            2,
		WindGustMPH:             2,
		WindDirectionDegree:     2,
		RainRateInch:            2,
		BarometricPressureABSIn: 2,
		BarometricPressureRelIn: 2,
	}
	test_host_1 = "host1"
	test_host_2 = "host2"
)

func TestRestGetHandler_implements_http_handler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, rest_path, nil)
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	RestGetHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
}

func TestRestGetHandler_invalid_method(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, rest_path, nil)
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	RestGetHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %d, got %d", http.StatusMethodNotAllowed, result.StatusCode)
	}
}

func TestUpdateStation_single(t *testing.T) {
	Stations = nil
	UpdateStation(&test_weather_1, test_host_1)

	if len(Stations) != 1 {
		t.Errorf("expected 1 station, got %d", len(Stations))
	}

	if Stations[0].Host != test_host_1 && Stations[0].LastData != test_weather_1 {
		t.Error("found wrong station, expected ", test_weather_1, " found ", Stations[0].LastData)
	}
}

func TestUpdateStation_double(t *testing.T) {
	Stations = nil
	UpdateStation(&test_weather_1, test_host_1)
	UpdateStation(&test_weather_2, test_host_2)

	if len(Stations) != 2 {
		t.Errorf("expected 2 station, got %d", len(Stations))
	}

	if Stations[0].Host != test_host_1 && Stations[0].LastData != test_weather_1 {
		t.Error("found wrong station, expected ", test_weather_1, " found ", Stations[0].LastData)
	}

	if Stations[1].Host != test_host_2 && Stations[1].LastData != test_weather_2 {
		t.Error("found wrong station, expected ", test_weather_2, " found ", Stations[1].LastData)
	}
}

func TestUpdateStation_duplicate(t *testing.T) {
	Stations = nil
	UpdateStation(&test_weather_1, test_host_1)
	UpdateStation(&test_weather_2, test_host_1)

	if len(Stations) != 1 {
		t.Errorf("expected 1 station, got %d", len(Stations))
	}

	if Stations[0].Host != test_host_1 && Stations[0].LastData != test_weather_2 {
		t.Error("found wrong station, expected ", test_weather_2, " found ", Stations[0].LastData)
	}

}

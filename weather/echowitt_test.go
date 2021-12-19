package weather

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	good_post_data = "PASSKEY=TESTKEY&stationtype=TestStation&dateutc=2006-01-02+15:04:05&tempinf=65.7&humidityin=54&baromrelin=29.986&baromabsin=29.986&tempf=44.6&humidity=57&winddir=2&windspeedmph=1.1&windgustmph=1.1&rainratein=0.000&eventrainin=0.331&dailyrainin=0.331&weeklyrainin=0.343&monthlyrainin=0.343&yearlyrainin=0.343&totalrainin=0.343&solarradiation=0.100&uv=0&model=MODEL"
	bad_post_data  = "PASSKEY=TESTKEY&stationtype=TestStation&dateutc=2006-01-02+15:04:05&tempinf=BADPARSE"
	non_form_data  = "This isn't form data; This is a sentance."
	echowitt_path  = "/v1/echowhitt"
)

func TestEchowittHandler_implements_http_handler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, echowitt_path, nil)
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	EcowittHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
}

func TestEchowittHandler_valid_data_parse(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, echowitt_path, strings.NewReader(good_post_data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	EcowittHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, result.StatusCode)
	}
}
func TestEchowittHandler_invalid_data_parse(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, echowitt_path, strings.NewReader(bad_post_data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	EcowittHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, result.StatusCode)
	}
}

func TestEchowittHandler_invalid_form_parse(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, echowitt_path, strings.NewReader(non_form_data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	EcowittHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, result.StatusCode)
	}
}

func TestEchowittHandler_invalid_method(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, echowitt_path, nil)
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	EcowittHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()
	if result.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %d, got %d", http.StatusMethodNotAllowed, result.StatusCode)
	}
}

func TestEchowittHandler_xff(t *testing.T) {
	Stations = nil
	req := httptest.NewRequest(http.MethodPost, echowitt_path, strings.NewReader(good_post_data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Forwarded-For", test_host_2)
	resp := httptest.NewRecorder()

	// Call EchowittHandler
	EcowittHandler(resp, req)
	result := resp.Result()
	defer result.Body.Close()

	if Stations[0].Host != test_host_2 {
		t.Error("found wrong host, expected ", test_host_2, " found ", Stations[0].Host)
	}
}

package estc

import (
	"encoding/json"
	"net/http"
)

// TestNewConfig test new config
func TestNewConfig(url string) *Config {
	return &Config{
		APIKey:       "testapikey",
		APISecret:    "testapisecret",
		BaseEndpoint: url,
		Debug:        true,
	}
}

func estimateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ETC{
		StatusCode:      SuccessStatusCode,
		ConfidenceLevel: 10,
		Time:            1000,
	})
	return
}

// DefaultHandlerMap default url and handler map
var DefaultHandlerMap = map[string]http.Handler{
	"/api/v1/estimate-time": http.HandlerFunc(estimateHandler),
}

// TestNewMux creates mux for test/dev server
func TestNewMux(hm map[string]http.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	for url, handler := range hm {
		mux.Handle(url, handler)
	}
	return mux
}

// TestNewServer creates dev server
func TestNewServer(hm map[string]http.Handler, port string) *http.Server {
	mux := TestNewMux(hm)
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}

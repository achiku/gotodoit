package estc

import (
	"encoding/json"
	"io/ioutil"
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
	json.NewEncoder(w).Encode(ETCResponse{
		StatusCode:      SuccessStatusCode,
		ConfidenceLevel: 10,
		Time:            1000,
	})
	return
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req UpdateUserRequest
	if err := json.Unmarshal(rawbody, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UpdateUserResponse{
		ID:         req.ID,
		Email:      req.Email,
		StatusCode: SuccessStatusCode,
		UserName:   req.UserName,
	})
	return
}

// DefaultHandlerMap default url and handler map
var DefaultHandlerMap = map[string]http.Handler{
	"/v1/estimate-time": http.HandlerFunc(estimateHandler),
	"/v1/users":         http.HandlerFunc(updateUserHandler),
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

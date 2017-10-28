package estc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_EstimateTimeToCompleteSimple(t *testing.T) {
	req := &Task{
		Name:       "CS 101 week one homework",
		Difficulty: 2,
	}
	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ETCResponse{
			StatusCode:      SuccessStatusCode,
			ConfidenceLevel: 10 * req.Difficulty,
			Time:            1000,
		})
		return
	}
	ts := httptest.NewServer(http.HandlerFunc(f))
	defer ts.Close()

	client := NewClient(TestNewConfig(ts.URL), &http.Client{}, nil)
	ctx := context.Background()
	res, err := client.EstimateTimeToComplete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != SuccessStatusCode {
		t.Errorf("want %d got %d", SuccessStatusCode, res.StatusCode)
	}
}

func TestClient_EstimateTimeToCompleteDefault(t *testing.T) {
	ts := httptest.NewServer(TestNewMux(DefaultHandlerMap))
	defer ts.Close()

	client := NewClient(TestNewConfig(ts.URL), &http.Client{}, nil)
	ctx := context.Background()
	req := &Task{
		Name:       "CS 101 week one homework",
		Difficulty: 2,
	}
	res, err := client.EstimateTimeToComplete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != SuccessStatusCode {
		t.Errorf("want %d got %d", SuccessStatusCode, res.StatusCode)
	}
}

func TestClient_EstimateTimeToComplete(t *testing.T) {
	errHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ETCResponse{
			StatusCode: ErrorStatusCode,
		})
		return
	})
	httpErrHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	})
	normalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ETCResponse{
			StatusCode:      SuccessStatusCode,
			ConfidenceLevel: 10,
			Time:            1000,
		})
		return
	})
	lowConfidenceHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ETCResponse{
			StatusCode:      SuccessStatusCode,
			ConfidenceLevel: 1,
			Time:            100000,
		})
		return
	})

	data := []struct {
		Handler    http.Handler
		StatusCode int
		Confidence int
		Fail       bool
	}{
		{Handler: errHandler, StatusCode: ErrorStatusCode, Confidence: 0, Fail: true},
		{Handler: httpErrHandler, StatusCode: ErrorStatusCode, Confidence: 0, Fail: true},
		{Handler: normalHandler, StatusCode: SuccessStatusCode, Confidence: 10, Fail: false},
		{Handler: lowConfidenceHandler, StatusCode: SuccessStatusCode, Confidence: 1, Fail: false},
	}

	for _, d := range data {
		mp := map[string]http.Handler{
			"/api/v1/estimate-time": d.Handler,
		}
		ts := httptest.NewServer(TestNewMux(mp))
		// not good for performance to defer in loop,
		// but it's just few loops, and not production code...
		defer ts.Close()

		ctx := context.Background()
		client := NewClient(TestNewConfig(ts.URL), &http.Client{}, nil)
		req := &Task{
			Name:       "CS 101 week one homework",
			Difficulty: 2,
		}
		res, err := client.EstimateTimeToComplete(ctx, req)
		if err != nil {
			if d.Fail {
				continue
			}
			t.Fatal(err)
		}
		if res.StatusCode != d.StatusCode {
			t.Errorf("want StatusCode=%d got %d", d.StatusCode, res.StatusCode)
		}
		if res.ConfidenceLevel != d.Confidence {
			t.Errorf(
				"want Confidence=%d got %d", d.Confidence, res.ConfidenceLevel)
		}
	}
}

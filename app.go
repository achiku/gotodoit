package gotodoit

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/achiku/gotodoit/estc"
	"github.com/achiku/gotodoit/model"
	"github.com/rs/xlog"
)

// App application
type App struct {
	DB         model.DBer
	Config     *Config
	EstcClient *estc.Client
}

// AppHandler internal
type AppHandler struct {
	h func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

// ServeHTTP serve
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := xlog.FromRequest(r)
	encoder := json.NewEncoder(w)
	reqInfo := xlog.F{"http_request": r}

	statusCode, res, err := ah.h(w, r)
	if err != nil {
		logger.Error(err, reqInfo)
		w.WriteHeader(statusCode)
		encoder.Encode(res)
		return
	}
	w.WriteHeader(statusCode)
	encoder.Encode(res)
	return
}

// NewApp creates app
func NewApp(cfgPath string) (*App, error) {
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	db, err := NewDB(cfg)
	if err != nil {
		return nil, err
	}
	estClient := estc.NewClient(&estc.Config{
		APIKey:       cfg.EstcConfig.APIKey,
		APISecret:    cfg.EstcConfig.APISecret,
		BaseEndpoint: cfg.EstcConfig.BaseEndpoint,
		Debug:        cfg.EstcConfig.Debug,
	}, &http.Client{}, log.New(os.Stdout, "[estc] ", log.LstdFlags))
	app := &App{
		DB:         db,
		Config:     cfg,
		EstcClient: estClient,
	}
	return app, nil
}

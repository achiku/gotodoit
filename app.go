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

// BaseApp base app global values
type BaseApp struct {
	Config     *Config
	EstcClient *estc.Client
}

// APIApp application
type APIApp struct {
	BaseApp
	DB     model.DBer
	Logger *log.Logger
}

// AppHandler internal
type AppHandler struct {
	h func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

// ServeHTTP serve
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := xlog.FromRequest(r)
	encoder := json.NewEncoder(w)

	statusCode, res, err := ah.h(w, r)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(statusCode)
		encoder.Encode(res)
		return
	}
	w.WriteHeader(statusCode)
	encoder.Encode(res)
	return
}

// NewApp creates app
func NewApp(cfgPath string) (*APIApp, error) {
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
	app := &APIApp{
		BaseApp: BaseApp{
			Config:     cfg,
			EstcClient: estClient,
		},
		DB:     db,
		Logger: log.New(os.Stdout, "[api-server] ", log.LstdFlags),
	}
	return app, nil
}

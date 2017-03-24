package gotodoit

import "github.com/achiku/gotodoit/model"

// NewDB creates db
func NewDB(cfg *Config) (model.DBer, error) {
	c := &model.DBConfig{
		DBName:  cfg.DBName,
		Host:    cfg.DBHost,
		Port:    cfg.DBPort,
		User:    cfg.DBUser,
		SSLMode: cfg.DBSSLMode,
	}
	db, err := model.NewDB(c)
	if err != nil {
		return nil, err
	}
	return db, nil
}

package gotodoit

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/achiku/qg"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/rs/xlog"
)

// Priority
const (
	QueNameHighPriority = "high"
	QueNameLowPriority  = ""
)

// NewQueClient create que client
func NewQueClient(db *sql.DB) (*qg.Client, error) {
	qc := qg.NewClient(db)
	return qc, nil
}

// NewWorkerDB create DB
func NewWorkerDB(config *Config) (*sql.DB, error) {
	dbCfg := &stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     config.DBHost,
			User:     config.DBUser,
			Password: config.DBPass,
			Database: config.DBName,
			Port:     config.DBPort,
		},
		AfterConnect: qg.PrepareStatements,
	}
	stdlib.RegisterDriverConfig(dbCfg)
	db, err := sql.Open("pgx", dbCfg.ConnectionString(""))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	// db.SetConnMaxLifetime(time.Second * 10)
	return db, nil
}

// StartWorker starts workers
func StartWorker(confPath string) error {
	appCfg, err := NewConfig(confPath)
	if err != nil {
		return err
	}
	db, err := NewWorkerDB(appCfg)
	if err != nil {
		return err
	}
	qc, err := NewQueClient(db)
	if err != nil {
		return err
	}

	jobs := JobApp{
		BaseApp: BaseApp{
			Config: appCfg,
		},
	}
	wm := qg.WorkMap{
		"updateUserJob": jobs.UpdateUserInfo,
	}
	wPoolLow := qg.NewWorkerPool(qc, wm, appCfg.NumWorkers)
	wPoolLow.Queue = QueNameLowPriority
	wPoolHigh := qg.NewWorkerPool(qc, wm, appCfg.NumWorkers)
	wPoolHigh.Queue = QueNameHighPriority

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	go func() {
		for {
			s := <-signalChan
			if s == syscall.SIGTERM {
				log.Printf("received SIGTERM. shutting down worker. (PID=%d)", os.Getpid())
				wPoolLow.Shutdown()
				wPoolHigh.Shutdown()
				log.Printf("shutting down worker. (PID=%d)", os.Getpid())
			}
		}
	}()
	xlog.Infof("starting workers. num workers: %d, (PID=%d)", appCfg.NumWorkers, os.Getpid())
	var wg sync.WaitGroup
	wg.Add(2)
	go wPoolLow.Start()
	go wPoolHigh.Start()
	wg.Wait()
	return nil
}

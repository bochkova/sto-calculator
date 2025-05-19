package app

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sto-calculator/internal/config"
	"sto-calculator/internal/handlers"
	"sto-calculator/internal/provider/postgres"
	"sto-calculator/internal/service"
	log "sto-calculator/pkg/logging"

	"github.com/go-chi/chi/v5"
)

type App struct {
	ctx     context.Context
	config  *config.Config
	closers []io.Closer
	server  *http.Server
}

func NewApp(configPath string) *App {
	var closers []io.Closer

	ctx, cancel := context.WithCancel(context.Background())
	go caughtSignals(cancel)

	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config from %s: %v", configPath, err)
	}
	log.WithFields(log.Fields{"config": cfg}).Info("Loaded config")

	if err = log.Configure(&log.Config{
		Level: cfg.Logger.Level,
	}); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}

	db, err := initDB(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}
	closers = append([]io.Closer{db}, closers...)

	newService := service.NewService(postgres.NewDB(db), service.NewFunctions())

	if err = newService.Init(ctx); err != nil {
		log.Fatalf("Failed to init service: %v", err)
	}

	router := chi.NewRouter()
	handlers.Register(router, handlers.NewHandlers(newService))

	return &App{
		ctx:     ctx,
		config:  cfg,
		closers: closers,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
			Handler: router,
		},
	}
}

func (a *App) Run() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Infof("Server started. Listening on %s", a.server.Addr)

	select {
	case <-a.ctx.Done():
		if err := a.server.Shutdown(context.Background()); err != nil {
			log.Errorf("Failed to shutdown server: %v", err)
		}
	}

	for _, closer := range a.closers {
		if err := closer.Close(); err != nil {
			log.Errorf("Failed to close: %v", err)
		}
	}
}

func caughtSignals(cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM,
	)

	log.Infof("Caught signal: %v. Shutting down...", <-quit)

	cancel()
}

package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/steevehook/mdc19/notes-rest-api/config"
	"github.com/steevehook/mdc19/notes-rest-api/controllers"
	"github.com/steevehook/mdc19/notes-rest-api/logging"
	"github.com/steevehook/mdc19/notes-rest-api/repositories"
	"github.com/steevehook/mdc19/notes-rest-api/services"
)

type App struct {
	Config config.Manager
	Router *httprouter.Router
	Server http.Server
}

func New() (*App, error) {
	err := logging.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	configManager, err := config.New()
	if err != nil {
		logging.Logger.Error("could not initialize config")
		return nil, err
	}

	db, err := repositories.DBConnection(repositories.DBOptions{
		URL:                     configManager.DBUrl(),
		MaxOpenConnections:      configManager.DBMaxOpenConnections(),
		MaxIdleConnections:      configManager.DBMaxIdleConnections(),
		ConnectionMaxLifetime:   configManager.DBConnMaxLifetime(),
		EnablePreparedStmtCache: configManager.DBEnablePreparedStmtCache(),
	})
	if err != nil {
		logging.Logger.Error("could not connect to the database")
		return nil, err
	}

	notesRepo := repositories.NewNotes(db)
	notesService := services.NewNotes(notesRepo)
	router := controllers.New(controllers.RouterConfig{
		NotesService: notesService,
	})
	app := &App{
		Config: configManager,
		Router: router,
	}

	return app, nil
}

func (app *App) Start() error {
	logging.Logger.Info(
		"application started",
		zap.String("listen", app.Config.Listen()),
	)
	server := http.Server{
		Addr:    app.Config.Listen(),
		Handler: app.Router,
	}
	app.Server = server
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		logging.Logger.Info("http server is closed")
		return nil
	}
	return err
}

func (app *App) Shutdown() {
	// Need to call Sync() before exiting. See Sync docs
	defer logging.Logger.Sync()
	ch := make(chan struct{})
	go func() {
		logging.Logger.Info("shutting down http server")
		if err := app.Server.Shutdown(context.Background()); err != nil {
			logging.Logger.Error("error on server shutdown", zap.Error(err))
		}
		close(ch)
	}()
	select {
	case <-ch:
		logging.Logger.Info("application was shut down")
	case <-time.After(2 * time.Second):
		logging.Logger.Error("could not shut down in", zap.Duration("timeout", 2*time.Second))
	}
}

type Shutdowner interface {
	Shutdown()
}

func ListenForSignals(signals []os.Signal, servers ...Shutdowner) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	sig := <-c
	logging.Logger.Info("received shutdown signal", zap.String("signal", sig.String()))

	for _, server := range servers {
		server.Shutdown()
	}
	os.Exit(0)
}

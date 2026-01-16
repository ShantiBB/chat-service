package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"chat-service/internal/config"
	"chat-service/internal/http/handler"
	"chat-service/internal/repository/postgres"
	"chat-service/internal/service"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
}

func (app *App) MustLoad() {
	slog.SetDefault(app.Logger)

	repo := postgres.New(app.Config)
	repo.Migrate()

	svc := service.New(repo)
	h := handler.New(svc, app.Config)

	mux := http.NewServeMux()
	h.Router(mux)

	server := app.server(mux)

	slog.Info("server started", "host", app.Config.Server.Host, "port", app.Config.Server.Port)

	app.runWithGracefulShutdown(server)
}

func (app *App) server(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			app.Config.Server.Host,
			app.Config.Server.Port,
		),
		Handler:      mux,
		ReadTimeout:  app.Config.Server.ReadTimeout,
		WriteTimeout: app.Config.Server.WriteTimeout,
		IdleTimeout:  app.Config.Server.IdleTimeout,
	}
}

func (app *App) runWithGracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen error", "error", err)
		}
	}()

	<-stop
	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), app.Config.Server.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}

	slog.Info("server stopped")
}

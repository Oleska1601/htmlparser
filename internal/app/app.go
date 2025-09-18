package app

import (
	"context"
	"fmt"
	"htmlparser/config"
	"htmlparser/internal/controller"
	"htmlparser/internal/database/repo"
	"htmlparser/internal/usecase"
	"htmlparser/pkg/logger"
	"htmlparser/pkg/postgres"
	"htmlparser/pkg/redis"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title           HTML Parser
// @version         1.0
// @description     Server for parsing html
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8081
// @BasePath  /
func Run(cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger := logger.New(cfg.Logger.Level)
	fmt.Println(cfg.Postgres.PgURL)
	postgres, err := postgres.New(ctx, cfg.Postgres.PgURL, postgres.SetMaxPoolSize(cfg.Postgres.MaxPoolSize))
	if err != nil {
		logger.Error("Run postgres.New", slog.Any("error", err))
		return
	}
	defer postgres.Close()
	pgRepo := repo.New(postgres)

	if err := pgRepo.ApplyMigrations(); err != nil {
		logger.Error("Run pgRepo.ApplyMigrations", slog.Any("error", err))
		return
	}
	redisClient, err := redis.New(ctx, cfg.Redis.RedisURL, cfg.Redis.RedisPassword, cfg.Redis.TTL)
	if err != nil {
		logger.Error("Run redis.New", slog.Any("error", err))
	}
	u := usecase.New(redisClient, pgRepo, logger)
	srv := controller.New(cfg.Server.Host, cfg.Server.Port, u, logger)

	go func() {
		if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("srv.Server.ListenAndServe", slog.Any("error", err))
		}
	}()

	logger.Info("server started", slog.Int("port", cfg.Server.Port))

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer shutdownCancel()
	if err := srv.Server.Shutdown(shutdownCtx); err != nil {
		logger.Error("srv.Server.Shutdown", slog.Any("error", err))
	}

	logger.Info("server exited properly")
}

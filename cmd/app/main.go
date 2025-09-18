package main

import (
	"htmlparser/config"
	"htmlparser/internal/app"
	"log/slog"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("main config.New", slog.Any("error", err))
		return
	}
	app.Run(cfg)
}

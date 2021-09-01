package main

import (
	c "github.com/ffo32167/weather/configs"
	l "github.com/ffo32167/weather/internal/logger"
	g "github.com/ffo32167/weather/internal/rpc"
	ch "github.com/ffo32167/weather/internal/storage"

	"github.com/sirupsen/logrus"
)

const (
	cache   = "mem"
	appName = "weatherParser"
)

func main() {
	// прочитать конфиг
	cfg, err := c.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	// настроить логер
	l.NewLog(appName, cfg.AppPath, cfg.SourceLinesInLog, cfg.LogLevel)

	// создать кэш
	logrus.Info("Reading Cache")
	cache := ch.ChooseCache(cache)
	cache.Load(cfg.AppPath)

	// запустить сервер
	logrus.Info("GRPC service starting up...")
	g.ServerStart(cfg, cache)
}

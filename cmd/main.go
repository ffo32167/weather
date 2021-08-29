package main

import (
	c "weather/configs"
	l "weather/internal/logger"
	g "weather/internal/rpc"
	ch "weather/internal/storage"

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

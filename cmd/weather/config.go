package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	tp "github.com/ffo32167/weather/internal/types"
)

func newConfig(configType, configPath string) (*tp.Config, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, fmt.Errorf("can't get path: %w", err)
	}

	switch {
	case configType == "JSON" && configPath != "":
		return newConfigJSON(configPath)
	case configType == "JSON":
		return newConfigJSON(path)
	}

	return nil, errors.New("can't choose config source")
}

// NewConfigJSON загрузить конфигурацию из JSON, настроить логи
func newConfigJSON(path string) (cfg *tp.Config, err error) {
	cfg = &tp.Config{}
	cfg.AppPath = path
	if cfg.AppPath == "" {
		return nil, fmt.Errorf("can't get application path: %w", err)
	}
	file, err := os.Open(filepath.Join(cfg.AppPath, `config.json`))
	if err != nil {
		return nil, fmt.Errorf("can't open config.json: %w", err)
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can't read config.json: %w", err)
	}
	err = json.Unmarshal(buff, &cfg)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal config.json: %w", err)
	}
	return cfg, nil
}

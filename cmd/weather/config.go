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
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	switch {
	case configType == "JSON" && configPath != "":
		return newConfigJSON(configPath)
	case configType == "JSON":
		return newConfigJSON(path)
	}
	return nil, fmt.Errorf("can't choose config source")
}

// NewConfigJSON загрузить конфигурацию из JSON, настроить логи
func newConfigJSON(path string) (cfg *tp.Config, err error) {
	cfg = &tp.Config{}
	cfg.AppPath = path
	if cfg.AppPath == "" {
		return nil, errors.New("can't get application path")
	}
	file, err := os.Open(filepath.Join(cfg.AppPath, `config.json`))
	if err != nil {
		return nil, fmt.Errorf("can't open config.json: %w", err)
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("can't read config.json")
	}
	err = json.Unmarshal(buff, &cfg)
	if err != nil {
		return nil, errors.New("can't unmarshal config.json")
	}
	return cfg, nil
}

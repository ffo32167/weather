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

// NewConfig загрузить конфигурацию, настроить логи
func NewConfig() (cfg *tp.Config, err error) {
	cfg = &tp.Config{}
	cfg.AppPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
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

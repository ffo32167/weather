package cache

import (
	"github.com/ffo32167/weather/internal"
)

// Cacher работает с кэшем: готовит к работе, создает путь к кэшу, читает кэш и сохраняет данные месяца в кэш
type Cacher interface {
	Load(path string)
	Path(pathParts ...string) string
	MonthRead(path string) ([]internal.DayWeather, error)
	MonthWrite(string, []internal.DayWeather)
}

// ChooseCache выбрать кэш
func ChooseCache(cache string) (ch Cacher) {
	switch cache {
	default:
		ch = NewWeatherMemCache()
	}
	return
}

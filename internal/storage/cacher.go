package cache

import tp "github.com/ffo32167/weather/internal/types"

// Cacher работает с кэшем: готовит к работе, создает путь к кэшу, читает кэш и сохраняет данные месяца в кэш
type Cacher interface {
	Load(path string)
	Path(pathParts ...string) string
	MonthRead(path string) ([]tp.DayWeather, error)
	MonthWrite(string, []tp.DayWeather)
}

// ChooseCache выбрать кэш
func ChooseCache(cache string) (ch Cacher) {
	switch cache {
	default:
		ch = NewWeatherMemCache()
	}
	return
}

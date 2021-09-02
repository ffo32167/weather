package pageparse

import (
	"github.com/ffo32167/weather/internal"
	"io"

	tp "github.com/ffo32167/weather/internal/types"
)

// SiteParser Интерфейс для получения данных с сайтов(worldWeather/yandexWeather)
type SiteParser interface {
	CreateDataPath(country, city, month, year string) (address string)
	SiteParse(source io.Reader, city string, month string, config tp.Config) []internal.DayWeather
}

// ChooseSiteParser Выбирает источник данных
func ChooseSiteParser(site string, config *tp.Config) SiteParser {
	switch site {
	case "worldweather":
		return worldWeather{address: config.WorldWeatherAddress}
	default:
		return yandex{address: config.YandexWeatherAddress}
	}

}

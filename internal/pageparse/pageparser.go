package pageparse

import (
	"io"

	c "weather/configs"
	w "weather/internal/types"
)

// SiteParser Интерфейс для получения данных с сайтов(worldWeather/yandexWeather)
type SiteParser interface {
	CreateDataPath(country, city, month, year string) (address string)
	SiteParse(source io.Reader, city string, month string, config c.Config) []w.DayWeather
}

// ChooseSiteParser Выбирает источник данных
func ChooseSiteParser(site string, config *c.Config) SiteParser {
	switch site {
	case "worldweather":
		return worldWeather{address: config.WorldWeatherAddress}
	default:
		return yandex{address: config.YandexWeatherAddress}
	}

}

package pageparse

import (
	"io"

	c "github.com/ffo32167/weather/cmd/weather/configs"
	w "github.com/ffo32167/weather/internal/types"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

// worldWeather содержит информацию для парсинга сайта worldWeather
type worldWeather struct {
	address string
}

// CreateDataPath cоздаёт путь к нужной странице вида:
// https://world-weather.ru/pogoda/russia/moscow/january-2018/
func (w worldWeather) CreateDataPath(country, city, month, year string) (url string) {
	return (w.address + country + "/" + city + "/" + month + "-" + year)
}

func (worldWeather) SiteParse(source io.Reader, city string, month string, config c.Config) (data []w.DayWeather) {
	var day w.DayWeather
	doc, err := goquery.NewDocumentFromReader(source)
	if err != nil {
		logrus.Error("can't parse page as HTML")
	}
	doc.Find(config.WorldWeatherSection).Each(func(i int, s *goquery.Selection) {
		day.City = city
		day.DayNumber = s.Find(config.WorldWeatherDayNumber).Text()
		day.Month = month
		day.TempDay = s.Find(config.WorldWeatherTempDay).Text()
		day.TempNight = s.Find(config.WorldWeatherTempNight).Text()
		day.Condition, _ = s.Find(config.WorldWeatherCondition).Attr(config.WorldWeatherConditionAttr)
		data = append(data, day)
	})
	return data
}

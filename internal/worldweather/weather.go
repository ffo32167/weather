package worldweather

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/ffo32167/weather/internal"
	"github.com/ffo32167/weather/internal/types"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type Source struct {
	baseAddress string
	config      types.Config
}

func (s Source) Get(city string, month string) ([]internal.DayWeather, error) {
	year := "2021"
	url := s.baseAddress + "russia" + "/" + city + "/" + month + "-" + year

	var data []internal.DayWeather
	var day internal.DayWeather
	source := getDataFromSite(url)
	doc, err := goquery.NewDocumentFromReader(source)
	if err != nil {
		logrus.Error("can't parse page as HTML")
	}
	doc.Find(s.config.WorldWeatherSection).Each(func(i int, selection *goquery.Selection) {
		day.City = city
		day.DayNumber = selection.Find(s.config.WorldWeatherDayNumber).Text()
		day.Month = month
		day.TempDay = selection.Find(s.config.WorldWeatherTempDay).Text()
		day.TempNight = selection.Find(s.config.WorldWeatherTempNight).Text()
		day.Condition, _ = selection.Find(s.config.WorldWeatherCondition).Attr(s.config.WorldWeatherConditionAttr)
		data = append(data, day)
	})
	return data, nil
}

func getDataFromSite(sitePath string) io.Reader {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(sitePath)
	if err != nil {
		logrus.WithFields(logrus.Fields{"sitePath": sitePath}).Error("can't complete request to site")
	}
	if resp.StatusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"sitePath": sitePath, "status code": resp.StatusCode}).Error("weather page not found")
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		logrus.Error("cant copy response:", sitePath)
	}
	return buf
}

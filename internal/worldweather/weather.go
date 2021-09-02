package worldweather

import (
	"github.com/ffo32167/weather/internal"
)

type Source struct {
}

func (s Source) Get(city string, month string) ([]internal.DayWeather, error) {
	panic("implement me")
}

package weatherdata

import "bytes"

// DayWeather это данные погоды за день
type DayWeather struct {
	City      string
	Month     string
	DayNumber string
	TempDay   string
	TempNight string
	Condition string
}

// WeatherParams это параметры запроса
type WeatherParams struct {
	MonthsNumbers []int32
	Cities        []string
	Site          string
	Months        []string
	Year          string
	ReplyFormat   string
}

// Encoder кодирует ответ в различные форматы
type Encoder interface {
	Encode([][]DayWeather, []string) (bytes.Buffer, string)
}

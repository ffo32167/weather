package internal

// DayWeather это данные погоды за день
type DayWeather struct {
	City      string
	Month     string
	DayNumber string
	TempDay   string
	TempNight string
	Condition string
}

type Source interface {
	Get(city string, month string) ([]DayWeather, error)
}

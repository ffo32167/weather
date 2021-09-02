package types

// WeatherParams это параметры запроса
type WeatherParams struct {
	MonthsNumbers []int32
	Cities        []string
	Site          string
	Months        []string
	Year          string
	ReplyFormat   string
}

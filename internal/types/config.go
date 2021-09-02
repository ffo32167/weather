package types

// Config содержит пути к сайтам, алиасы полей для парсинга и настройки
type Config struct {
	AppPath  string
	LogLevel string `json:"LogLevel"`
	GrpcPort string `json:"GrpcPort"`
	// SourceLinesInLog определяет нужно ли указывать номера строк там где было вызвано логирование
	SourceLinesInLog *bool `json:"SourceLinesInLog"`
	// поля для парсинга данных из Yandex
	YandexWeatherAddress       string            `json:"YandexWeatherAddress"`
	YandexWeatherSection       string            `json:"YandexWeatherSection"`
	YandexWeatherDayNumber     string            `json:"YandexWeatherDayNumber"`
	YandexWeatherTemp          string            `json:"YandexWeatherTemp"`
	YandexWeatherCondition     string            `json:"YandexWeatherCondition"`
	YandexWeatherConditionAttr string            `json:"YandexWeatherConditionAttr"`
	YandexWeatherMap           map[string]string `json:"YandexWeatherMap"`
	// поля для парсинга данных из WorldWeather
	WorldWeatherAddress       string `json:"WorldWeatherAddress"`
	WorldWeatherSection       string `json:"WorldWeatherSection"`
	WorldWeatherDayNumber     string `json:"WorldWeatherDayNumber"`
	WorldWeatherTempDay       string `json:"WorldWeatherTempDay"`
	WorldWeatherTempNight     string `json:"WorldWeatherTempNight"`
	WorldWeatherCondition     string `json:"WorldWeatherCondition"`
	WorldWeatherConditionAttr string `json:"WorldWeatherConditionAttr"`
}

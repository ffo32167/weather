package types

import "bytes"

// Encoder кодирует ответ в различные форматы
type Encoder interface {
	Encode([][]DayWeather, []string) (bytes.Buffer, string)
}

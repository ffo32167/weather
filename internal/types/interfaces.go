package types

import (
	"bytes"
	"github.com/ffo32167/weather/internal"
)

// Encoder кодирует ответ в различные форматы
type Encoder interface {
	Encode([][]internal.DayWeather, []string) (bytes.Buffer, string)
}

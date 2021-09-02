package encode

import w "github.com/ffo32167/weather/internal/types"

// ChooseEncoder выбирает формат ответа
func ChooseEncoder(format string) (e w.Encoder) {
	switch format {
	case "JSON":
		e = newJSON()
	default:
		e = newCSV()
	}
	return
}

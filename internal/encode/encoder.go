package encode

import w "weather/internal/types"

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

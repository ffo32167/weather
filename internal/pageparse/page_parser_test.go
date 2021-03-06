package pageparse

import (
	"reflect"
	"testing"

	tp "github.com/ffo32167/weather/internal/types"
)

func TestChooseSiteParser(t *testing.T) {
	type args struct {
		site   string
		config *tp.Config
	}
	tests := []struct {
		name       string
		args       args
		wantSource SiteParser
	}{
		{"basic worldWeather", args{"worldweather", &cfg}, worldWeather{cfg.WorldWeatherAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSource := ChooseSiteParser(tt.args.site, tt.args.config); !reflect.DeepEqual(gotSource, tt.wantSource) {
				t.Errorf("ChooseSiteParser() = %v, want %v", gotSource, tt.wantSource)
			}
		})
	}
}

package yr

import (
	//"fmt"

	"strconv"
	"testing"

	"github.com/carolinesh/funtemps/conv"
)

// Test fahrenheit to celsius
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "6", want: "42.8"},
	}

	for _, tc := range tests {
		f, err := strconv.ParseFloat(tc.input, 64)
		if err != nil {
			t.Errorf("error ved konvertering av input %v", tc.input)
		}

		got := conv.FahrenheitToCelsius(f)
		s := strconv.FormatFloat(f, 'f', 2, 64)
		if s != tc.want {
			t.Errorf("forventet: %v, fikk: %v", tc.want, got)
		}
	}
}

// Test celsius to fahrenheit

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "42.8", want: "6"},
	}

	for _, tc := range tests {
		f, err := strconv.ParseFloat(tc.input, 64)
		if err != nil {
			t.Errorf("error ved konvertering av input %v", tc.input)
		}

		got := conv.CelsiusToFahrenheit(f)
		s := strconv.FormatFloat(f, 'f', 2, 64)
		if s != tc.want {
			t.Errorf("forventet: %v, fikk: %v", tc.want, got)
		}
	}
}

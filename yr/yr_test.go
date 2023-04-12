package yr

import (
	//"fmt"
	"reflect"
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
		{input: "10", want: "-12.2"},
	}

	for _, tc := range tests {
		f, err := strconv.ParseFloat(tc.input, 64)
		if err != nil {
			t.Errorf("error ved konvertering av input %v", tc.input)
		}

		str := strconv.FormatFloat(f, 'f', 6, 64)
		if str != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, str)
		}
	}

	// utføre testen

	for _, tc := range tests {
		got := conv.FahrenheitToCelsius('f')
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
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
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
	}

	for _, tc := range tests {
		f, err := strconv.ParseFloat(tc.input, 64)
		if err != nil {
			t.Errorf("error ved konvertering av input %v", tc.input)
		}

		str := strconv.FormatFloat(f, 'f', 6, 64)
		if str != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, str)
		}
	}

	// utføre testen

	for _, tc := range tests {
		got := conv.CelsiusToFahrenheit('f')
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}

/* Konvertere fra string til float64
	str := "6"
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("error ved konvertering")
	}


	// Konvertere fra float64 til string
   {
    f := 42.8
    str := strconv.FormatFloat(f,'f',6,64 ) }

	//Utføre testen
	tests := []test{
		{input: "6", want:"42.8" },
	}

	for _, tc := range tests {
		got := conv.FahrenheitToCelsius(f)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
*/

// Sett inn testfunksjonen fra den forige oppgaven her
// Bruke string conv pakke for å kunne konvertere string
// til float og float til string

package yr

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestReadLines(t *testing.T) {
	file, err := os.Open("../kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	lines, err := ReadLines(file)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	expected := 16756
	if len(lines) != expected {
		t.Fatalf("Unexpected number of lines. Expected %v, but got %v", expected, len(lines))
	}
}
func TestConvertTemperatures(t *testing.T) {
	lines, err := ConvertTemperatures()
	if err != nil {
		t.Fatalf("ConvertTemperatures failed: %v", err)
	}

	// Set up input values
	input := "Kjevik;SN39040;18.03.2022 01:50;6"
	want := "Kjevik;SN39040;18.03.2022 01:50;42.80F"

	// Find the input value in the output slice
	var got string
	for _, line := range lines {
		if strings.HasPrefix(line, input) {
			got = line
			break
		}
	}

	if got != want {
		fmt.Println(want)
		fmt.Println(got)
	}
}

package yr

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/carolinesh/funtemps/conv"
)

func ConvertTemperatures() ([]string, error) {
	file, err := OpenFile("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)

	ConvertedTemperatures := make([]string, 0)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			continue // ignorerer overskriftslinjen
		}

		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			return nil, fmt.Errorf("uventet antall felt i linje %d: %d", i, len(fields))
		}

		if fields[3] == "" {
			continue // ignorerer linjer med tomme temperaturfelt
		}

		TemperatureCelsius, err := strconv.ParseFloat(fields[3], 64)

		if err != nil {
			return nil, fmt.Errorf("kunne ikke parse temperatur i linje %d: %s", i, err)
		}
		TemperatureFahrenheit := CelsiusToFahrenheit(TemperatureCelsius)

		ConvertedTemperature := fmt.Sprintf("%s;%s;%.2fF", fields[0], strings.Join(fields[1:3], ";"), TemperatureFahrenheit)
		ConvertedTemperatures = append(ConvertedTemperatures, ConvertedTemperature)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ConvertedTemperatures, nil
}

func GetAndWriteTemperatures(filename string) error {
	lines, err := ConvertTemperatures()
	if err != nil {
		return err
	}
	return WriteLines(lines, filename)
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return conv.CelsiusToFahrenheit(celsius)
}

func WriteLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer CloseFile(file)
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Skriver overskriftslinjen
	fmt.Fprintln(writer, "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur (F)")

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	fmt.Fprint(writer, "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Caroline")

	return nil
}

func ReadLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Navn") {
			continue // ignorerer overskriftslinjen
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func OpenFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	return file, err
}

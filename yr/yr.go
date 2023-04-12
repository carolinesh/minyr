package yr

import (
	"bufio"
	"fmt"
	"github.com/carolinesh/funtemps/conv"
	"log"
	"os"
	"strconv"
	"strings"
)

// Funkjsonen for å lese filen, konvertere og printe ut ny fil.

func GetData() {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	// Bruker scanner og finner vedrien til gradene som er 4. elementet

	scanner := bufio.NewScanner(file)
	var tempValues []float64

	// looper igjennom for å hente ut verdiene

	for i := 0; scanner.Scan(); i++ {
		if i == 0 { //Hopper over første linjen i dataen
			continue
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ";")
		tempCelsius, err := strconv.ParseFloat(values[3], 64)
		if err != nil {
			panic(err)
		}
		tempValues = append(tempValues, tempCelsius)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// avlutt og create ny fil med de nye verdiene.

	outFile, err := os.Create("nyFil.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	for _, tempCelsius := range tempValues {
		tempFahrenheit := conv.CelsiusToFahrenheit(tempCelsius)
		fmt.Fprintf(writer, "%.2f\n", tempFahrenheit)
	}
	writer.Flush() //Tømmer bufferen for mellomlagret data
}

func Yr() {
	GetData()
}

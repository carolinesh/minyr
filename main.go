package main

import (

	//"io/ioutil"
	"bufio"
	"os"

	//"os"
	//"iotuil"
	"log"
	//"github.com/carolinesh/funtemps/conv"
)

// Funkjsonen for å lese filen

func main() {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal("Unable to read file: %v", err)
	}
	defer file.Close()

	// Bruker scanner og finner vedrien til gradene som er 4. elementet

	scanner := bufio.NewScanner(file)
	var value string
	for i := 0; scanner.Scan(); i++ {
		if i == 3 {
			value = scanner.Text()
			break
		}
	}
}

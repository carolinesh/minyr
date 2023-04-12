package main

import (
	"bufio"
	"fmt"
	"os"
	"yr"
)

// Det som brukeren skal kunne skrive inn i terminalviduet

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		if input == "exit" {
			fmt.Println("Programmet avslutter")
			os.Exit(0)

		} else if input == "convert" {
			fmt.Println("Koverterer målingene og printer ut ny fil")
			// Må legge til Y/N her?
			// Må kalle på metoden som gjør at filen konverterer og skriver ut ny til med gitt navn
			 func GetData()

		} else if input == "average" {
			fmt.Println("Regner ut gjennomsnittet")
			//Må kalle til en funkjson
		} else {
			fmt.Println("Velg mellom convert, average eller exit")
		}
	}
}

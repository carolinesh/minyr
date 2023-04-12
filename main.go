package main

import (
	"bufio"
	"fmt"
	"minyr/yr"
	"os"
	"strings"
)

func main() {
	fmt.Println("Velg mellom convert, average eller exit")

	var input string
	var answer string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = strings.ToLower(scanner.Text())
		if input == "exit" {
			fmt.Println("Programmet avslutter")
			os.Exit(0)

		} else if input == "convert" {
			fmt.Println("Velg Y eller N:")
			scanner.Scan()
			answer = strings.ToLower(scanner.Text())
			if answer == "y" {
				fmt.Println("Koverterer målingene og printer ut ny fil")
				yr.GetData()
			} else if answer == "n" {
				fmt.Println("n")
			} else {
				fmt.Println("Velg Y eller N:")
			}

		} else if input == "average" {
			fmt.Println("Regner ut gjennomsnittet")
		} else {
			fmt.Println("Velg mellom convert, average eller exit")
		}
	}
}

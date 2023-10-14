package main

import (
	"fmt"
)

func main() {

	mode := ""

	fmt.Println("mode: ")
	fmt.Scanln(&mode)

	switch mode {
	case "encode", "e":
		universal("+")
	case "decode", "d":
		universal("-")
	}

}

func universal(operator string) {
	text := ""
	shift := 0
	shifted := 0

	fmt.Println("text: ")

	fmt.Scanln(&text)

	fmt.Println("shift: ")

	fmt.Scanln(&shift)

	for _, rune := range text {
		if operator == "+" {
			shifted = (int(rune) + shift)
		}
		if operator == "-" {
			shifted = (int(rune) - shift)
		}
		fmt.Printf(string(shifted))
	}
}

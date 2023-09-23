package main

import (
	"fmt"

	aux "github.com/giancarlo-mena/code-wars-go/challenges"
)

func main() {
	printBuildTower(2)
}

func printReverseWords(str string) {
	fmt.Print(aux.ReverseWords(str))
}

func printDecodeMorse(str string) {
	fmt.Print(aux.DecodeMorse(str))
}

func printBuildTower(floors int) {
	fmt.Print(aux.TowerBuilder(floors))
}

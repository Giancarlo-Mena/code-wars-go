package main

import (
	"fmt"

	lvl5 "github.com/giancarlo-mena/code-wars-go/challenges/5"
	lvl6 "github.com/giancarlo-mena/code-wars-go/challenges/6"
	lvl7 "github.com/giancarlo-mena/code-wars-go/challenges/7"
)

func main() {
	printBitCounting(1234)
}

func printReverseWords(str string) {
	fmt.Print(lvl7.ReverseWords(str))
}

func printDecodeMorse(str string) {
	fmt.Print(lvl6.DecodeMorse(str))
}

func printBuildTower(floors int) {
	fmt.Print(lvl6.TowerBuilder(floors))
}

func printDeadfishSwim(str string) {
	fmt.Print(lvl6.Parse(str))
}

func printHumanReadableTime(seconds int) {
	fmt.Print(lvl5.HumanReadableTime(seconds))
}

func printTortoiseRace(v1, v2, g int) {
	fmt.Print(lvl6.Race(v1, v2, g))
}

func printCamelCaseMethod(str string) {
	fmt.Print(lvl6.CamelCase(str))
}

func printBitCounting(number uint) {
	fmt.Print((lvl6.CountBits(number)))
}
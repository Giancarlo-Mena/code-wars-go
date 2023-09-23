package main

import (
	"fmt"

	lvl6 "github.com/giancarlo-mena/code-wars-go/challenges/6"
	lvl7 "github.com/giancarlo-mena/code-wars-go/challenges/7"
)

func main() {
	printDeadfishSwim("ioioio")
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

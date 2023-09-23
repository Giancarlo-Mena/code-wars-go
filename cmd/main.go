package main

import (
	"fmt"

	aux "github.com/giancarlo-mena/code-wars-go/challenges"
)

func main() {
	printDecodeMorse(".... . -.--   .--- ..- -.. .")
}

func printReverseWords(str string) {
	fmt.Print(aux.ReverseWords(str))
}

func printDecodeMorse(str string) {
	fmt.Print(aux.DecodeMorse(str))
}

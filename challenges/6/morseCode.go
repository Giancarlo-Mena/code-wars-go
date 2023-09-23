package challenges

import "strings"

var morseCodeMap = map[string]byte{
	".-":   'A',
	"-...": 'B',
	"-.-.": 'C',
	"-..":  'D',
	".":    'E',
	"..-.": 'F',
	"--.":  'G',
	"....": 'H',
	"..":   'I',
	".---": 'J',
	"-.-":  'K',
	".-..": 'L',
	"--":   'M',
	"-.":   'N',
	"---":  'O',
	".--.": 'P',
	"--.-": 'Q',
	".-.":  'R',
	"...":  'S',
	"-":    'T',
	"..-":  'U',
	"...-": 'V',
	".--":  'W',
	"-..-": 'X',
	"-.--": 'Y',
	"--..": 'Z',

	"-----": '0',
	".----": '1',
	"..---": '2',
	"...--": '3',
	"....-": '4',
	".....": '5',
	"-....": '6',
	"--...": '7',
	"---..": '8',
	"----.": '9',
}

func DecodeMorse(morseCode string) string {
	decoded := ""

	for _, word := range strings.Split(strings.TrimSpace(morseCode), "  ") {
		for _, letter := range strings.Split(word, " ") {
			decoded += string(morseCodeMap[letter])
		}
		decoded += " "
	}

	return decoded
}

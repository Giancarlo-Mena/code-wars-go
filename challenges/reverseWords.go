package challenges

import "strings"

func ReverseWords(str string) string {
	strs := strings.Split(str, " ")
	result := make([]string, len(strs))

	for _, s := range strs {
		var aux string = ""

		for j := len(s) - 1; j >= 0; j-- {
			aux += string(s[j])
		}
		
		result = append(result, aux)
	}

	return strings.Join(result, " ")
}

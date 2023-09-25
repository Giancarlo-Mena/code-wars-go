package challenges

import "strings"

func CamelCase(s string) string {
	arr := strings.Split(s, " ")
	
	for i, word := range arr {
		if len(word) > 0 {
			arr[i] = strings.Replace(word, string(word[0]), string(word[0] - 32), 1)
		}
	}

    return strings.Join(arr, "")
}
package kata

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}
	
	result := ""
	count := 0

	date := map[string]int64{
		"year":   seconds / 31536000,
		"day":    seconds % 31536000 / 86400,
		"hour":   seconds % 86400 / 3600,
		"minute": seconds % 3600 / 60,
		"second": seconds % 60,
	}

	keys :=  []string{"year", "day", "hour", "minute", "second"}

	for _, key := range keys {
		value := date[key]
		if value == 0 {
			continue
		}

		if count != 0 {
			result += ", "
		}

		result += fmt.Sprint(value) + " " + key
		count++

		if value > 1 {
			result += "s"
		}
	}

	if count > 1 {
		lastComma := strings.LastIndex(result, ",")
		result = result[:lastComma] + " and" + result[lastComma+1:]
	}
	
	return result
}
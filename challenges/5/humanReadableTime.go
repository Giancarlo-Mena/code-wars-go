package challenges

import "fmt"

func HumanReadableTime(seconds int) string {
	if seconds >= 359999 {
		return "99:59:59"
	}

	result := ""
	sec := "00"
	min := "00"
	hour := "00"

	hour = fmt.Sprintf("%02d", seconds / 3600)
	seconds = seconds % 3600
	min = fmt.Sprintf("%02d",  seconds / 60)
	seconds = seconds % 60
	sec = fmt.Sprintf("%02d", seconds)

	result += hour + ":" + min + ":" + sec
	return result
}

func HumanBetterVersion(seconds int) string {
	hour := seconds / 3600
	min := seconds % 3600 / 60
	sec := seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}
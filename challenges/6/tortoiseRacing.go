package challenges

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	var d2 float32 = 0.0
	as := float32(v2-v1) / 3600
	count := 0

	for g > int(d2) {
		d2 += as
		count++
	}

	return [3]int{count / 3600, count % 60 / 60, count % 60}
}
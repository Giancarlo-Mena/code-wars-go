package challenges

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	as := float32(g) / float32(v2-v1)
	min := (as - float32(int(as))) * 60
	sec := (min - float32(int(min))) * 60

	return [3]int{int(as), int(min), int(sec)}
}
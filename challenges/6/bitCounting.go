package challenges

func CountBits(number uint) int {
	var result uint = 0;
	for number != 0 {
		result += number % 2
		number /= 2	
	}
	return int(result)
}


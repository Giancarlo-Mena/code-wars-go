package challenges

var result []int = []int{}
var number int = 0

func Parse(data string) []int {
	number = 0
	result = []int{}

	for _, char := range data {
		checkChar(byte(char))
	}

	return result
}

func checkChar(char byte) {
	switch char {
	case 'i':
		number++
	case 'd':
		number--
	case 's':
		number *= number
	case 'o':
		result = append(result, number)	
	default:
		break
	}
}
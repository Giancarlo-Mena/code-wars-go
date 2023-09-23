package challenges

import "strings"

func TowerBuilder(floors int) []string {
	space := ""
	blocks := "*"
	tower := make([]string, floors)

	for i := 0; i < floors-1; i++ {
		space += " "
	}

	for i := 0; i < len(tower); i++ {
		tower[i] = space + blocks + space
		space = strings.Replace(space, " ", "", 1)
		blocks += "**"

	}

	return tower
}
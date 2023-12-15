package leetcode

func judgeCircle(moves string) bool {
	var x, y int
	for i := 0; i < len(moves); i++ {
		switch string(moves[i]) {
		case "U":
			y++
		case "D":
			y--
		case "R":
			x++
		case "L":
			x--
		}
	}
	return x == 0 && y == 0
}

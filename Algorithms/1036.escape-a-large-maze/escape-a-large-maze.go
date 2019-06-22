package problem1036

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

// M is length and width of the large maze
const M = 1e6

// MAX 代表了 blocked 能够围起来的最大数量
// 题目里面的说了 len(blocked)<=200
// 这样的话，能够围起来的点，不会超过 20000 个
const MAX = 20000

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	isBlocked := make(map[int]bool, len(blocked))
	for _, b := range blocked {
		x, y := b[0], b[1]
		isBlocked[x<<32+y] = true
	}

	isOpen := func(source, target []int) bool {
		tx, ty := target[0], target[1]
		//
		points := make([][]int, 0, MAX)
		hasSeen := make(map[int]bool, MAX)
		//
		sx, sy := source[0], source[1]
		points, hasSeen[sx<<32+sy] = append(points, source), true
		//
		count := 0
		for len(points) > 0 {
			size := len(points)
			for i := 0; i < size; i++ {
				px, py := points[i][0], points[i][1]
				for k := 0; k < 4; k++ {
					x, y := px+dx[k], py+dy[k]
					if x == tx && y == ty {
						return true
					}
					p := x<<32 + y
					if 0 <= x && x < M &&
						0 <= y && y < M &&
						!isBlocked[p] &&
						!hasSeen[p] {
						points, hasSeen[p] = append(points, []int{x, y}), true
					}
				}
			}
			points = points[size:]
			count += size
			if count >= MAX {
				return true
			}
		}
		return false
	}

	return isOpen(source, target) && isOpen(target, source)
}

package lib06

func SolvePrt1(board Board, guard Player) int {

	width := len((board)[0])
	height := len(board)

	(board)[guard.Y][guard.X] = Visited

	visitedDirections := make(VisitedDirections, height)
	for i := range visitedDirections {
		visitedDirections[i] = make([][]Direction, width)
	}

	for {
		nextX, nextY := guard.GetNextPos()
		if !inside(width, height, nextX, nextY) {
			break
		}

		if (board)[nextY][nextX] != Blocked {
			for _, dir := range visitedDirections[nextY][nextX] {
				if dir == guard.Direction {
					// We are in loop
					return -1
				}
			}
			(board)[nextY][nextX] = Visited
			visitedDirections[nextY][nextX] = append(visitedDirections[nextY][nextX], guard.Direction)
			guard.X = nextX
			guard.Y = nextY
		} else {
			guard.Direction = (guard.Direction + 1) % 4
		}
	}

	return board.GetScore()
}

func inside(width, height, x, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

type VisitedDirections [][][]Direction

func SolvePrt2(board Board, guard Player) int {

	score := 0

	for i := range board {
		for j := range (board)[i] {
			if i == guard.Y && j == guard.X {
				continue
			}
			if (board)[i][j] == Blocked {
				continue
			}

			alteredBoard := board.Clone()
			alteredBoard[i][j] = Blocked

			res := SolvePrt1(alteredBoard, guard)
			if res == -1 {
				score++
			}
		}
	}

	return score
}

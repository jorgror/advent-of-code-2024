package lib06

import (
	"bufio"
	"io"
)

type Block int

const (
	Unvisited Block = iota
	Visited
	Blocked
)

type Board [][]Block

func (b Board) GetScore() int {
	score := 0
	for _, row := range b {
		for _, block := range row {
			if block == Visited {
				score++
			}
		}
	}
	return score
}

func (b Board) Print() {
	for i := range b {
		for j := range b[i] {
			switch b[i][j] {
			case Unvisited:
				print(".")
			case Visited:
				print("X")
			case Blocked:
				print("#")
			}
		}
		println()
	}
}

func (b Board) Clone() Board {
	clone := make(Board, len(b))
	for i := range b {
		clone[i] = make([]Block, len(b[i]))
		for j := range b[i] {
			clone[i][j] = b[i][j]
		}
	}
	return clone
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Player struct {
	X, Y      int
	Direction Direction
}

func (p Player) GetNextPos() (int, int) {
	switch p.Direction {
	case Up:
		p.Y--
	case Right:
		p.X++
	case Down:
		p.Y++
	case Left:
		p.X--
	}
	return p.X, p.Y
}

func Parse(input io.Reader) (Board, Player) {
	scanner := bufio.NewScanner(input)

	board := Board{}
	guard := Player{}
	// Parse the board
	for scanner.Scan() {
		line := scanner.Text()
		width := len(line)
		row := make([]Block, width)
		for i, c := range line {
			switch c {
			case '.':
				row[i] = Unvisited
			case '#':
				row[i] = Blocked
			case '^':
				row[i] = Visited
				guard = Player{X: i, Y: len(board), Direction: Up}
			case '>':
				row[i] = Visited
				guard = Player{X: i, Y: len(board), Direction: Right}
			case 'v':
				row[i] = Visited
				guard = Player{X: i, Y: len(board), Direction: Down}
			case '<':
				row[i] = Visited
				guard = Player{X: i, Y: len(board), Direction: Left}
			default:
				panic("invalid character")
			}
		}
		board = append(board, row)
	}

	return board, guard
}

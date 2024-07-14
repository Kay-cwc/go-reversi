package chessboard

import (
	"fmt"
	"strconv"
	"strings"
)

type Chessboard struct {
	dimension uint
	board     [][]string // 0 represent unfilled
}

const ChessDefault string = "-"
const ChessPlayer1 string = "X"
const ChessPlayer2 string = "O"

func InitChessboard(dimension uint) Chessboard {
	board := make([][]string, dimension+1)

	for row := range dimension + 1 {
		// special handling for the first row as the column indices (not for playing)
		if row == 0 {
			colIndices := make([]string, dimension+1)
			colIndices[0] = " "
			for i := uint64(0); i < uint64(dimension); i++ {
				colIndices[i+1] = strconv.FormatUint(i+1, 10)
			}
			board[row] = colIndices
			continue
		}
		// fill the rest of the chessboard
		// the first col of each row will be used as the row indices
		board[row] = make([]string, dimension+1)
		for col := range dimension + 1 {
			if col == 0 {
				board[row][col] = strconv.FormatUint(uint64(row), 10)
				continue
			}
			board[row][col] = ChessDefault
		}
	}
	// by default, {33, 34, 43, 44} will be filled
	board[4][4], board[5][5] = ChessPlayer1, ChessPlayer1
	board[4][5], board[5][4] = ChessPlayer2, ChessPlayer2
	return Chessboard{
		dimension,
		board,
	}
}

// print the chessboard in a readable format
// 0,1,2 will be parsed as string "-", "X", "O" for display
func Print(chessboard *Chessboard) {
	// dimension+1 as the board's xy are + 1 to contain the indices
	for row := range chessboard.dimension + 1 {
		rowVal := chessboard.board[row]
		printRowString(rowVal)
	}
}

var adjacent_def = [3]int{-1, 0, 1}

func moveCell(cell [2]uint, changes [2]int) [2]uint {
	x := uint(changes[0] + int(cell[0]))
	y := uint(changes[1] + int(cell[1]))

	return [2]uint{x, y}
}

func isInBound(cell [2]uint, dimension uint) bool {
	return cell[0] > 0 && cell[0] <= dimension && cell[1] > 0 && cell[1] <= dimension
}

func IsAdjacentToOpponent(chessboard *Chessboard, chess string, move [2]uint) bool {
	// first move +- xy by 1, then +- yx by 1
	var adjacentCells = make([][2]uint, 0, 8)

	for _, changeX := range adjacent_def {
		for _, changeY := range adjacent_def {
			adjacentCell := moveCell(move, [2]int{changeX, changeY})
			if !isInBound(adjacentCell, chessboard.dimension) {
				continue
			}
			// check if it is filled by opponent
			currentFill := chessboard.board[adjacentCell[0]][adjacentCell[1]]
			if currentFill == chess || currentFill == ChessDefault {
				continue
			}

			adjacentCells = append(adjacentCells, adjacentCell)
		}
	}

	return len(adjacentCells) > 0
}

// handle player move on chessboard. this function does not check if the rules comply the game rules
func Move(chessboard *Chessboard, chess string, move [2]uint) {
	chessboard.board[move[1]][move[0]] = chess
}

func printRowString(row []string) {
	fmt.Println(strings.Join(row[:], "|"))
}

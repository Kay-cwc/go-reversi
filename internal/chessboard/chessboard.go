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

func IsAvailable(chessboard *Chessboard, move [2]uint) bool {
	return chessboard.board[move[1]][move[0]] == ChessDefault
}

func findSurroundingCells(chessboard *Chessboard, chess string, cell [2]uint, change [2]int) [][2]uint {
	surroundingCells := make([][2]uint, 0, 8)
	for {
		nextCell := moveCell(cell, change)
		if !isInBound(nextCell, chessboard.dimension) {
			break
		}

		currentFill := chessboard.board[nextCell[1]][nextCell[0]]
		if currentFill == chess {
			// find an ending cell that is placed by the current player. valid adjacent cell
			surroundingCells = append(surroundingCells, nextCell)
			break
		} else if currentFill == ChessDefault {
			// no ending cell is placed by the current player. invalid adjacent cell
			break
		}
		// is filled by opponent, continue the path and search
		cell = nextCell
	}
	return surroundingCells
}

// check all the surrounding chess of the potential move
// only return those cells that is in bound and it filled by opponents
func IsAdjacentToOpponent(chessboard *Chessboard, chess string, move [2]uint) [][2]uint {
	fmt.Println(chess, move)
	// first move +- xy by 1, then +- yx by 1
	var adjacentCells = make([][2]uint, 0, 8)

	for _, changeX := range adjacent_def {
		for _, changeY := range adjacent_def {
			change := [2]int{changeX, changeY}
			adjacentCell := moveCell(move, change)
			if !isInBound(adjacentCell, chessboard.dimension) {
				continue
			}
			// check if it is filled by opponent
			currentFill := chessboard.board[adjacentCell[1]][adjacentCell[0]]
			if currentFill == chess || currentFill == ChessDefault {
				continue
			}

			fmt.Println(adjacentCell, currentFill, chess)
			// extends from the potential move to all surrounding cells and beyond
			// find the straight path that has a cell occupied by self
			surroundingCells := findSurroundingCells(chessboard, chess, move, change)
			if len(surroundingCells) > 0 {
				adjacentCells = append(adjacentCells, adjacentCell)
			}
		}
	}

	// in return, we don't need to tell the ending cell. the chessboard will handle later.
	// we just need to make sure at least one of the adjacent cell is valid
	return adjacentCells
}

type UserMove struct {
	Move          [2]uint
	AdjacentCells [][2]uint
}

// handle player move on chessboard. this function does not check if the rules comply the game rules
func Move(chessboard *Chessboard, chess string, userMove UserMove) {
	chessboard.board[userMove.Move[1]][userMove.Move[0]] = chess
	fmt.Println(userMove.AdjacentCells)
	// flip all surroundingCells
	for _, adjacentCell := range userMove.AdjacentCells {
		chessboard.board[adjacentCell[1]][adjacentCell[0]] = chess
		change := [2]int{int(adjacentCell[0]) - int(userMove.Move[0]), int(adjacentCell[1]) - int(userMove.Move[1])}
		surroundingCells := findSurroundingCells(chessboard, chess, userMove.Move, change)
		fmt.Println(surroundingCells)
		for _, surroundingCell := range surroundingCells {
			chessboard.board[surroundingCell[1]][surroundingCell[0]] = chess
		}
	}
}

func printRowString(row []string) {
	fmt.Println(strings.Join(row[:], "|"))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Chessboard struct {
	dimension uint
	board     [][]string // 0 represent unfilled
}

const chessDefault string = "-"
const chessPlayer1 string = "X"
const chessPlayer2 string = "O"

func initChessboard(dimension uint) Chessboard {
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
			board[row][col] = chessDefault
		}
	}
	// by default, {33, 34, 43, 44} will be filled
	board[4][4], board[5][5] = chessPlayer1, chessPlayer1
	board[4][5], board[5][4] = chessPlayer2, chessPlayer2
	return Chessboard{
		dimension,
		board,
	}
}

func printRowString(row []string) {
	fmt.Println(strings.Join(row[:], "|"))
}

// print the chessboard in a readable format
// 0,1,2 will be parsed as string "-", "X", "O" for display
func printChessboard(chessboard *Chessboard) {
	// dimension+1 as the board's xy are + 1 to contain the indices
	for row := range chessboard.dimension + 1 {
		rowVal := chessboard.board[row]
		printRowString(rowVal)
	}
}

type validatePromptResp[O uint | []uint] func(string) (O, string, bool)

func prompt[O uint | []uint](q string, validateFuncs []validatePromptResp[O]) O {
	var ans string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, q+":\n")
		ans, _ = r.ReadString('\n')
		if ans != "" {
			ans = strings.TrimSpace(ans)
			// should perform validation here
			for _, validateFunc := range validateFuncs {
				output, errorMsg, validationError := validateFunc(ans)
				if validationError {
					fmt.Println(errorMsg)
					continue
				}
				return output
			}
		}
	}
}

func isUintString(v string) (uint, string, bool) {
	intVal, err := strconv.ParseUint(v, 10, 64)
	return uint(intVal), "please input a positive integer", err != nil
}

type Game struct {
	player1    bool
	chessboard Chessboard
}

func initGame() Game {
	dimension := prompt(
		"tell me the dimension of the chessboard (default=8)",
		[]validatePromptResp[uint]{isUintString},
	)
	chessboard := initChessboard(dimension)
	fmt.Println("New Game Started!")
	printChessboard(&chessboard)

	return Game{
		chessboard: chessboard,
		player1:    true,
	}
}

func validateUserMovePrompt(v string) ([]uint, string, bool) {
	targets := strings.Split(v, ",")
	output := make([]uint, 2)
	if len(targets) != 2 {
		return output, "Invalid input", true
	}
	var hasErr bool
	var errMsg string
	output[0], errMsg, hasErr = isUintString(targets[0])
	if hasErr {
		return output, errMsg, hasErr
	}
	output[1], errMsg, hasErr = isUintString(targets[1])
	if hasErr {
		return output, errMsg, hasErr
	}
	// the value should be between 1-8
	// should be dynamic later for the chessboard dimension
	// maybe return a closure here
	if output[0] == 0 || output[0] > 8 {
		return output, "x must be between 1-8", true
	}
	if output[1] == 0 || output[1] > 8 {
		return output, "y must be between 1-8", true
	}

	return output, "", false
}

func userMove(game *Game) {
	playerName := "Player 1"
	chess := chessPlayer1
	if !game.player1 {
		playerName = "Player 2"
		chess = chessPlayer2
	}
	q := fmt.Sprintf("Player %s's move (input your move in x,y format):", playerName)
	move := prompt(q, []validatePromptResp[[]uint]{validateUserMovePrompt})
	// update state
	game.player1 = !game.player1
	game.chessboard.board[move[1]][move[0]] = chess
}

func main() {
	// entry point of the game
	// the game will be running as a loop
	game := initGame()
	for {
		userMove(&game)
		printChessboard(&game.chessboard)
	}
}

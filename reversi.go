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
	board     [][]uint // 0 represent unfilled
}

func initChessboard(dimension uint) Chessboard {
	board := make([][]uint, dimension)
	for row := range dimension {
		board[row] = make([]uint, dimension)
		for col := range dimension {
			board[row][col] = 0
		}
	}
	// by default, {33, 34, 43, 44} will be filled
	board[3][3], board[4][4] = 1, 1
	board[3][4], board[4][3] = 2, 2
	return Chessboard{
		dimension,
		board,
	}
}

var displayValMap = map[uint]string{
	0: "-",
	1: "X",
	2: "O",
}

func printRowString(row []string) {
	fmt.Println(strings.Join(row[:], "|"))
}

// print the chessboard in a readable format
// 0,1,2 will be parsed as string "-", "X", "O" for display
func printChessboard(chessboard *Chessboard) {
	dimension := chessboard.dimension
	// col indices
	// +1 because we need to reserve space for the row index
	colIndices := make([]string, dimension+1)
	colIndices[0] = " "
	for i := uint64(0); i < uint64(dimension); i++ {
		fmt.Println(i)
		colIndices[i+1] = strconv.FormatUint(i+1, 10)
	}
	fmt.Println(colIndices)
	printRowString(colIndices)

	for row := range dimension {
		rowVal := chessboard.board[row]
		rowDisplay := make([]string, dimension+1)
		rowDisplay[0] = strconv.FormatUint(uint64(row+1), 10)
		for col := range dimension {
			rowDisplay[col+1] = displayValMap[rowVal[col]]
		}
		printRowString(rowDisplay)
	}
}

type validatePromptResp[O uint | []uint] func(string) (O, string, bool)

func prompt[O uint | []uint](q string, validate validatePromptResp[O]) O {
	var ans string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, q+":\n")
		ans, _ = r.ReadString('\n')
		if ans != "" {
			ans = strings.TrimSpace(ans)
			// should perform validation here
			output, errorMsg, validationError := validate(ans)
			if validationError {
				fmt.Println(errorMsg)
				continue
			}
			// break when user type something
			return output
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
	dimension := prompt[uint]("tell me the dimension of the chessboard (default=8)", isUintString)
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
		return output, "Invalid input", false
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
	if !game.player1 {
		playerName = "Player 2"
	}
	q := fmt.Sprintf("Player %s's move (input your move in x,y format):", playerName)
	move := prompt(q, validateUserMovePrompt)
	// update state
	player := uint(1)
	if !game.player1 {
		player = 2
	}
	game.player1 = !game.player1
	game.chessboard.board[move[1]-1][move[0]-1] = player
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

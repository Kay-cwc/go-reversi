package main

import (
	"fmt"
	"strings"

	// "kayton/chiu/internal"
	chessboard "kayton/chiu/reversi/internal/chessboard"
	prompt "kayton/chiu/reversi/internal/prompt"
	validation "kayton/chiu/reversi/internal/validation"
)

type Chessboard struct {
	dimension uint
	board     [][]string // 0 represent unfilled
}

const chessDefault string = "-"
const chessPlayer1 string = "X"
const chessPlayer2 string = "O"

// func initChessboard(dimension uint) Chessboard {
// 	board := make([][]string, dimension+1)

// 	for row := range dimension + 1 {
// 		// special handling for the first row as the column indices (not for playing)
// 		if row == 0 {
// 			colIndices := make([]string, dimension+1)
// 			colIndices[0] = " "
// 			for i := uint64(0); i < uint64(dimension); i++ {
// 				colIndices[i+1] = strconv.FormatUint(i+1, 10)
// 			}
// 			board[row] = colIndices
// 			continue
// 		}
// 		// fill the rest of the chessboard
// 		// the first col of each row will be used as the row indices
// 		board[row] = make([]string, dimension+1)
// 		for col := range dimension + 1 {
// 			if col == 0 {
// 				board[row][col] = strconv.FormatUint(uint64(row), 10)
// 				continue
// 			}
// 			board[row][col] = chessDefault
// 		}
// 	}
// 	// by default, {33, 34, 43, 44} will be filled
// 	board[4][4], board[5][5] = chessPlayer1, chessPlayer1
// 	board[4][5], board[5][4] = chessPlayer2, chessPlayer2
// 	return Chessboard{
// 		dimension,
// 		board,
// 	}
// }

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

type Game struct {
	player1    bool
	chessboard chessboard.Chessboard
}

func initGame() Game {
	dimension := prompt.Ask(
		"tell me the dimension of the chessboard (default=8)",
		[]prompt.ValidatePromptResp[uint]{validation.IsUintString},
	)
	board := chessboard.InitChessboard(dimension)
	fmt.Println("New Game Started!")
	chessboard.Print(&board)

	return Game{
		chessboard: board,
		player1:    true,
	}
}

func userMove(game *Game) {
	playerName := "Player 1"
	chess := chessPlayer1
	if !game.player1 {
		playerName = "Player 2"
		chess = chessPlayer2
	}
	q := fmt.Sprintf("Player %s's move (input your move in x,y format):", playerName)
	move := prompt.Ask(q, []prompt.ValidatePromptResp[[]uint]{validation.ValidateUserMovePrompt})
	// update state
	game.player1 = !game.player1
	chessboard.Move(&game.chessboard, chess, move)
}

func main() {
	// entry point of the game
	// the game will be running as a loop
	game := initGame()
	for {
		userMove(&game)
		chessboard.Print(&game.chessboard)
	}
}

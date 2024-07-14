package game

import (
	"fmt"
	chessboard "kayton/chiu/reversi/internal/chessboard"
	parser "kayton/chiu/reversi/internal/parser"
	prompt "kayton/chiu/reversi/internal/prompt"
)

type Game struct {
	player1    bool
	Chessboard chessboard.Chessboard
}

func InitGame() Game {
	dimension := prompt.Ask(
		"tell me the dimension of the chessboard (default=8)",
		parser.IsUintString,
	)
	board := chessboard.InitChessboard(dimension)
	fmt.Println("New Game Started!")
	chessboard.Print(&board)

	return Game{
		Chessboard: board,
		player1:    true,
	}
}

// since the move depends on the current game state
// this function will return a closure instead
func isValidMove(game *Game) func(v string) ([2]uint, string, bool) {
	return func(v string) ([2]uint, string, bool) {
		userMove, errMsg, hasErr := parser.ValidateUserMoveInput(v)

		return userMove, errMsg, hasErr
	}
}

func UserMove(game *Game) {
	playerName := "Player 1"
	chess := chessboard.ChessPlayer1
	if !game.player1 {
		playerName = "Player 2"
		chess = chessboard.ChessPlayer2
	}
	q := fmt.Sprintf("Player %s's move (input your move in x,y format):", playerName)
	move := prompt.Ask(q, isValidMove(game))
	// update state
	game.player1 = !game.player1
	chessboard.Move(&game.Chessboard, chess, move)
}

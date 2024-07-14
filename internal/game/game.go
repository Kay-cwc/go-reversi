package game

import (
	"fmt"
	chessboard "kayton/chiu/reversi/internal/chessboard"
	parser "kayton/chiu/reversi/internal/parser"
	prompt "kayton/chiu/reversi/internal/prompt"
)

type Game struct {
	player     string
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
		player:     chessboard.ChessPlayer1,
	}
}

// since the move depends on the current game state
// this function will return a closure instead
func isValidMove(game *Game) func(v string) ([2]uint, string, bool) {
	return func(v string) ([2]uint, string, bool) {
		userMove, errMsg, hasErr := parser.ValidateUserMoveInput(v)
		isValidMove := chessboard.IsAdjacentToOpponent(&game.Chessboard, game.player, userMove)
		if !isValidMove {
			errMsg = "Invalid Move"
			hasErr = true
		}
		return userMove, errMsg, hasErr
	}
}

func isPlayer1(player string) bool {
	return player != chessboard.ChessPlayer1
}

func UserMove(game *Game) {
	playerName := "Player 1"
	if isPlayer1(game.player) {
		playerName = "Player 2"
	}
	q := fmt.Sprintf("Player %s's move (input your move in x,y format):", playerName)
	move := prompt.Ask(q, isValidMove(game))
	// update state
	chessboard.Move(&game.Chessboard, game.player, move)

	if isPlayer1(game.player) {
		game.player = chessboard.ChessPlayer2
	} else {
		game.player = chessboard.ChessPlayer1
	}
}

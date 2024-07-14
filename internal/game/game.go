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
func isValidMove(game *Game) func(v string) (chessboard.UserMove, string, bool) {
	return func(v string) (chessboard.UserMove, string, bool) {
		move, errMsg, hasErr := parser.ValidateUserMoveInput(v)
		var adjacentCells = make([][2]uint, 0, 8)
		userMove := chessboard.UserMove{
			Move:          move,
			AdjacentCells: adjacentCells,
		}
		if !chessboard.IsAvailable(&game.Chessboard, move) {
			errMsg = "Invalid Move - the cell is filled already"
			hasErr = true

			return userMove, errMsg, hasErr
		}

		adjacentCells = chessboard.IsAdjacentToOpponent(&game.Chessboard, game.player, move)
		if len(adjacentCells) == 0 {
			errMsg = "Invalid Move"
			hasErr = true
		}

		userMove.AdjacentCells = adjacentCells
		return userMove, errMsg, hasErr
	}
}

func isPlayer1(player string) bool {
	return player == chessboard.ChessPlayer1
}

func Move(game *Game) {
	q := fmt.Sprintf("Player %s move (input your move in x,y format):", game.player)
	userMove := prompt.Ask(q, isValidMove(game))
	// update state
	chessboard.Move(&game.Chessboard, game.player, userMove)

	if isPlayer1(game.player) {
		game.player = chessboard.ChessPlayer2
	} else {
		game.player = chessboard.ChessPlayer1
	}
}

package game

import (
	"fmt"
	chessboard "kayton/chiu/reversi/internal/chessboard"
	prompt "kayton/chiu/reversi/internal/prompt"
	validation "kayton/chiu/reversi/internal/validation"
)

type Game struct {
	player1    bool
	Chessboard chessboard.Chessboard
}

func InitGame() Game {
	dimension := prompt.Ask(
		"tell me the dimension of the chessboard (default=8)",
		[]prompt.ValidatePromptResp[uint]{validation.IsUintString},
	)
	board := chessboard.InitChessboard(dimension)
	fmt.Println("New Game Started!")
	chessboard.Print(&board)

	return Game{
		Chessboard: board,
		player1:    true,
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
	move := prompt.Ask(q, []prompt.ValidatePromptResp[[]uint]{validation.ValidateUserMovePrompt})
	// update state
	game.player1 = !game.player1
	chessboard.Move(&game.Chessboard, chess, move)
}

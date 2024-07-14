package main

import (
	"fmt"

	// "kayton/chiu/internal"
	chessboard "kayton/chiu/reversi/internal/chessboard"
	gameController "kayton/chiu/reversi/internal/game"
)

// entry point of the game
// the game will be running as a loop
func main() {
	fmt.Println(`[Game Rule]
1. must be adjacent to another opponent chess
2. must form a straight line with another chess of yours (straight or tilted)
3. the opponent chess in between your chess and your new move will be flipped
	`)

	game := gameController.InitGame()
	for {
		gameController.Move(&game)
		chessboard.Print(&game.Chessboard)
	}
}

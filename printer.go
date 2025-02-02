package main

import "fmt"

type Printer struct{}

func createPrinter() *Printer {
	return &Printer{}
}

func (l *Printer) printBoard(board *Board) {
	for i := 0; i < board.size; i++ {
		fmt.Println("")
	}

	for i := 0; i < board.size; i++ {
		fmt.Println(board.fields[i])
	}

	fmt.Println("Use Arrow Keys to Move; Press Esc to end the game.")
}

func (printer *Printer) printGameEndedMessage(game *Game, r string) {
	for i := 0; i < 10; i++ {
		fmt.Println("")
	}

	fmt.Println("========================")
	fmt.Println("====== GAME OVER =======")
	fmt.Println("========================")
	fmt.Println(r)
	fmt.Println("")
	fmt.Println("Score:", game.score)

	for i := 0; i < 10; i++ {
		fmt.Println("")
	}
}

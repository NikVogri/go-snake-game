package main

type Game struct {
	score int32
	ended bool
}

func createGame() *Game {
	return &Game{
		score: 0,
		ended: false,
	}
}

func (game *Game) incrementScore() {
	game.score++
}

func (game *Game) end() {
	game.ended = true
}

func (game *Game) shouldSpeedBeIncreased() bool {
	return game.score%5 == 0
}

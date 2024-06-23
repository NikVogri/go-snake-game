package main

type Game struct {
	score           int32
	speedMultiplier float64
	ended           bool
}

var game = Game{
	score:           0,
	speedMultiplier: 1.0,
	ended:           false,
}

func (game *Game) incrementScore() {
	game.score++
}

func (game *Game) increaseSpeed() {
	game.speedMultiplier -= 0.05
}

func (game *Game) end() {
	game.ended = true
}

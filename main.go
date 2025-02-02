package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

var BOARD_SIZE = 32
var BASE_GAME_TICK_DURATION = 150 * time.Millisecond

var game = createGame()
var board = createBoard(BOARD_SIZE)

var snake = createSnake(board)
var fruit = createFruit(board)

var ticker = createTicker()
var printer = createPrinter()

func main() {
	go listenForKeyboardEvents()

	ticker.setSpeed(BASE_GAME_TICK_DURATION)

	ticker.start(func(t *Ticker) {
		if game.ended {
			printer.printGameEndedMessage(game, "You've ended the game")
			t.stop()
			return
		}

		nextPosition := snake.getNextHeadPosition()

		if board.isOutsideBorder(nextPosition) {
			printer.printGameEndedMessage(game, "You went over the border")
			t.stop()
			return
		}

		if fruit.hasBeenEaten(nextPosition) {
			snake.growByOne()
			fruit.respawn(board)
			game.incrementScore()

			if game.shouldSpeedBeIncreased() {
				ticker.increaseSpeedByFactor(0.05)
			}
		}

		snake.moveToNextPosition(nextPosition)

		if snake.hasEatenItself() {
			printer.printGameEndedMessage(game, "You've eaten yourself")
			t.stop()
			return
		}

		board.update(snake, fruit)
		printer.printBoard(board)
	})

}

func listenForKeyboardEvents() {
	defer keyboard.Close()

	if err := keyboard.Open(); err != nil {
		log.Fatal("Failed to open keyboard:", err)
		return
	}

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err)
			break
		}

		if key == keyboard.KeyEsc {
			game.end()
			break
		}

		directionMap := map[keyboard.Key]Direction{
			keyboard.KeyArrowUp:    DIRECTION_NORTH,
			keyboard.KeyArrowDown:  DIRECTION_SOUTH,
			keyboard.KeyArrowLeft:  DIRECTION_WEST,
			keyboard.KeyArrowRight: DIRECTION_EAST,
		}

		if directionMap[key] == "" {
			continue
		}

		snake.changeDirection(directionMap[key])
	}
}

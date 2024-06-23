package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/eiannone/keyboard"
)

type Point struct {
	x int
	y int
}

var BOARD_SIZE = 32

var board = createBoard(BOARD_SIZE)
var snake = createSnake()
var fruit = createFruit()

func main() {
	board.draw()
	go listenForKeyboardEvents()

	fmt.Println(snake.headPosition)

	for {
		if game.ended {
			fmt.Println("========================")
			fmt.Println("====== GAME OVER =======")
			fmt.Println("========================")
			fmt.Println("You've ended the game   ")
			fmt.Println("Score: ", game.score)
			break
		}

		nextPosition := snake.getNextHeadPosition()

		if board.isOutsideBorder(nextPosition) {
			fmt.Println("========================")
			fmt.Println("====== GAME OVER =======")
			fmt.Println("========================")
			fmt.Println("You went over the border")
			fmt.Println("Score: ", game.score)
			break
		}

		// check if fruit was eaten
		if nextPosition.x == fruit.position.x && nextPosition.y == fruit.position.y {
			snake.growByOne()
			fruit.respawn()
			game.incrementScore()

			if game.score%5 == 0 {
				game.increaseSpeed()
			}
		}

		snake.moveToNextPosition(nextPosition)

		// check if body part was eaten
		if snake.hasEatenItself() {
			fmt.Println("========================")
			fmt.Println("====== GAME OVER =======")
			fmt.Println("========================")
			fmt.Println("You've eaten yourself   ")
			fmt.Println("Score: ", game.score)
			break
		}

		board.redraw()
		board.print()

		time.Sleep(time.Millisecond * time.Duration(math.Round(125*game.speedMultiplier)))
	}
}

func listenForKeyboardEvents() {
	if err := keyboard.Open(); err != nil {
		log.Fatal("Failed to open keyboard:", err)
		return
	}
	defer keyboard.Close()

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

		var newDirection string

		if key == keyboard.KeyArrowUp {
			newDirection = DIRECTION_NORTH
		}

		if key == keyboard.KeyArrowDown {
			newDirection = DIRECTION_SOUTH
		}

		if key == keyboard.KeyArrowLeft {
			newDirection = DIRECTION_WEST
		}

		if key == keyboard.KeyArrowRight {
			newDirection = DIRECTION_EAST
		}

		snake.changeDirection(newDirection)
	}
}

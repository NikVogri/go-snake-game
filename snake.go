package main

import "math"

const (
	DIRECTION_NORTH = "DIRECTION_NORTH"
	DIRECTION_EAST  = "DIRECTION_EAST"
	DIRECTION_SOUTH = "DIRECTION_SOUTH"
	DIRECTION_WEST  = "DIRECTION_WEST"
)

type SnakeBodyNode struct {
	position Point
}

type Snake struct {
	headPosition  Point
	direction     string
	body          []SnakeBodyNode
	pastPositions []Point
}

func createSnake() *Snake {
	return &Snake{
		headPosition: Point{
			// setting starting position
			x: int(math.Round(float64(board.size / 2))),
			y: int(math.Round(float64(board.size / 2))),
		},
		direction:     DIRECTION_EAST,
		body:          []SnakeBodyNode{},
		pastPositions: []Point{},
	}
}

func (snake *Snake) changeDirection(direction string) {
	snake.direction = direction
}

func (snake *Snake) getNextHeadPosition() Point {
	var newPosition Point

	if snake.direction == DIRECTION_NORTH {
		newPosition = Point{x: snake.headPosition.x, y: snake.headPosition.y - 1}
	}

	if snake.direction == DIRECTION_SOUTH {
		newPosition = Point{x: snake.headPosition.x, y: snake.headPosition.y + 1}
	}

	if snake.direction == DIRECTION_WEST {
		newPosition = Point{x: snake.headPosition.x - 1, y: snake.headPosition.y}
	}

	if snake.direction == DIRECTION_EAST {
		newPosition = Point{x: snake.headPosition.x + 1, y: snake.headPosition.y}
	}

	return newPosition
}

func (snake *Snake) growByOne() {

	node := SnakeBodyNode{
		position: snake.headPosition,
	}

	snake.body = append(snake.body, node)
}

func (snake *Snake) moveToNextPosition(point Point) {
	snake.headPosition = point
	snake.pastPositions = append(snake.pastPositions, point)
	snake.moveBodyNodesToNextPosition()

}

func (snake *Snake) moveBodyNodesToNextPosition() {
	for i := range snake.body {
		index := len(snake.pastPositions) - 2 - i

		if index < 0 {
			index = 0
		}

		snake.body[i].position = snake.pastPositions[index]

	}
}

func (snake *Snake) hasEatenItself() bool {
	hasEatenItself := false

	for _, node := range snake.body {
		if node.position.x == snake.headPosition.x && node.position.y == snake.headPosition.y {
			hasEatenItself = true
			break
		}
	}

	return hasEatenItself

}

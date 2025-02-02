package main

type Direction string

const (
	DIRECTION_NORTH Direction = "DIRECTION_NORTH"
	DIRECTION_EAST  Direction = "DIRECTION_EAST"
	DIRECTION_SOUTH Direction = "DIRECTION_SOUTH"
	DIRECTION_WEST  Direction = "DIRECTION_WEST"
)

var DIRECTION_MAP = map[Direction]*Vector2{
	DIRECTION_NORTH: {y: -1, x: 0},
	DIRECTION_SOUTH: {y: 1, x: 0},
	DIRECTION_WEST:  {y: 0, x: -1},
	DIRECTION_EAST:  {y: 0, x: 1},
}

type SnakeBodyNode struct {
	position Vector2
}

type Snake struct {
	headPosition  Vector2
	direction     Direction
	body          []SnakeBodyNode
	pastPositions []Vector2
}

func createSnake(board *Board) *Snake {
	return &Snake{
		headPosition:  board.getCenterPosition(),
		direction:     DIRECTION_EAST,
		body:          []SnakeBodyNode{},
		pastPositions: []Vector2{},
	}
}

func (snake *Snake) changeDirection(direction Direction) {
	snake.direction = direction
}

func (snake *Snake) getNextHeadPosition() Vector2 {
	if DIRECTION_MAP[snake.direction] == nil {
		return Vector2{x: 0, y: 0}
	}

	nextHeadPos := snake.headPosition.copy()
	nextHeadPos.add(DIRECTION_MAP[snake.direction])

	return nextHeadPos
}

func (snake *Snake) growByOne() {
	node := SnakeBodyNode{
		position: snake.headPosition,
	}

	snake.body = append(snake.body, node)
}

func (snake *Snake) moveToNextPosition(pos Vector2) {
	snake.headPosition = pos
	snake.pastPositions = append(snake.pastPositions, pos)
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
		if snake.headPosition.overlaps(node.position) {
			hasEatenItself = true
			break
		}
	}

	return hasEatenItself
}

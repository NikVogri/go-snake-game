package main

import (
	"math"
	"math/rand"
)

type FieldType string

const (
	FIELD_EMPTY      FieldType = "-"
	FIELD_SNAKE_HEAD FieldType = "O"
	FIELD_SNAKE_BODY FieldType = "#"
	FIELD_FRUIT      FieldType = "F"
)

type Board struct {
	fields [][]FieldType
	size   int
}

func createBoard(size int) *Board {
	board := &Board{
		fields: [][]FieldType{},
		size:   size,
	}

	board.init()

	return board
}

func (board *Board) update(snake *Snake, fruit *Fruit) {
	board.init()

	board.setField(fruit.position, FIELD_FRUIT)
	board.setField(snake.headPosition, FIELD_SNAKE_HEAD)

	for _, node := range snake.body {
		board.setField(node.position, FIELD_SNAKE_BODY)
	}
}

func (board *Board) init() {
	board.fields = make([][]FieldType, board.size)

	for i := 0; i < board.size; i++ {
		cols := make([]FieldType, board.size)

		for i := 0; i < board.size; i++ {
			cols[i] = FIELD_EMPTY
		}

		board.fields[i] = cols
	}
}

func (board *Board) setField(vector Vector2, fieldType FieldType) {
	board.fields[vector.y][vector.x] = fieldType
}

func (board *Board) isOutsideBorder(position Vector2) bool {
	board_wall := board.size - 1

	if position.x > board_wall || position.x < 0 {
		return true
	}

	if position.y > board_wall || position.y < 0 {
		return true
	}

	return false
}

func (board *Board) getRandomField() Vector2 {
	return Vector2{
		x: rand.Intn(board.size),
		y: rand.Intn(board.size),
	}
}

func (board *Board) getCenterPosition() Vector2 {
	return Vector2{
		x: int(math.Round(float64(board.size / 2))),
		y: int(math.Round(float64(board.size / 2))),
	}
}

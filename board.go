package main

import (
	"fmt"
	"math/rand"
)

const (
	FIELD_EMPTY = "-"
	FIELD_HEAD  = "O"
	FIELD_BODY  = "#"
	FIELD_FRUIT = "F"
)

type Board struct {
	fields [][]string
	size   int
}

func createBoard(size int) *Board {
	return &Board{
		fields: [][]string{},
		size:   size,
	}
}

func (board *Board) draw() {
	newBoard := make([][]string, board.size)

	for i := 0; i < board.size; i++ {
		cols := make([]string, board.size)

		for i := 0; i < board.size; i++ {
			cols[i] = FIELD_EMPTY
		}

		newBoard[i] = cols
	}

	board.fields = newBoard
}

func (board *Board) redraw() {
	board.draw()

	board.fields[fruit.position.y][fruit.position.x] = FIELD_FRUIT
	board.fields[snake.headPosition.y][snake.headPosition.x] = FIELD_HEAD

	for _, node := range snake.body {
		board.fields[node.position.y][node.position.x] = FIELD_BODY
	}
}

func (board *Board) print() {
	for i := 0; i < board.size; i++ {
		fmt.Println("")
	}

	for i := 0; i < board.size; i++ {
		fmt.Println(board.fields[i])
	}
}

func (board *Board) isOutsideBorder(point Point) bool {
	board_wall := board.size - 1

	if point.x > board_wall || point.x < 0 {
		return true
	}

	if point.y > board_wall || point.y < 0 {
		return true
	}

	return false
}

func (board *Board) pickRandomField() Point {
	return Point{
		x: rand.Intn(board.size),
		y: rand.Intn(board.size),
	}
}

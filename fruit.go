package main

type Fruit struct {
	position Vector2
}

func createFruit(board *Board) *Fruit {
	return &Fruit{
		position: board.getRandomField(),
	}
}

func (fruit *Fruit) respawn(board *Board) {
	fruit.position = board.getRandomField()
}

func (fruit *Fruit) hasBeenEaten(position Vector2) bool {
	return position.overlaps(fruit.position)
}

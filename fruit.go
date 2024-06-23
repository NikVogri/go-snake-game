package main

type Fruit struct {
	position Point
}

func createFruit() *Fruit {
	return &Fruit{
		position: board.pickRandomField(),
	}
}

func (fruit *Fruit) respawn() {
	fruit.position = board.pickRandomField()
}

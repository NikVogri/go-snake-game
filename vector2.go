package main

type Vector2 struct {
	x int
	y int
}

func (vector *Vector2) add(v *Vector2) {
	vector.x += v.x
	vector.y += v.y
}

func (vector *Vector2) overlaps(v Vector2) bool {
	return vector.x == v.x && vector.y == v.y
}

func (vector *Vector2) copy() Vector2 {
	return Vector2{x: vector.x, y: vector.y}
}

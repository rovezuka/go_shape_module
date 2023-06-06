package main

import "github.com/rovezuka/go_shape_module/shape"

func main() {
	c := shape.Circle{Radius: 5.0}
	r := shape.Rectangle{Width: 4.0, Height: 3.0}

	shape.PrintArea(c) // Выводит площадь круга
	shape.PrintArea(r) // Выводит площадь прямоугольника

	shape.PrintPerimeter(c) // Выводит периметр круга
	shape.PrintPerimeter(r) // Выводит периметр прямоугольника
}

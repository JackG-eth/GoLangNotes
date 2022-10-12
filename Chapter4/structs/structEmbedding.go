package main

import "fmt"

//go lets us declare a type but with no name, we say here that a point is embedded in a circle and so on. This way you don't need to give full sequences of names
//if you want to maintain using the shorthand these structs must be exported0

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	 Circle
	spokes int
}

func main() {
	var w Wheel
	w.x = 8
	w.y = 8
	w.Radius = 5
	w.spokes = 20

	var v = Wheel{Circle{Point{8,8},5},20}
	fmt.Printf("%#v\n",v)
}
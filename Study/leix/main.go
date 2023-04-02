package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	point := Point{
		X: 10,
		Y: 20,
	}

	var TmpA interface{}

	TmpA = point

	var TmpB Point

	TmpB = TmpA.(Point)
	fmt.Println(TmpB)
}

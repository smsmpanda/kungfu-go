package examples

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func Going411() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	//fmt.Println(screen)

	array := [3]float64{7.0, 8.5, 9.1}
	x := sum(&array)
	fmt.Println(x)
}

func sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

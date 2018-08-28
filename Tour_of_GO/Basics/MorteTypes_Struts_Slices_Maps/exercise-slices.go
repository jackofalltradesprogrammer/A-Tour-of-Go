package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		slice[i] = make([]uint8, dx)
	}

	for y, _ := range slice {
		img := 0
		for x, _ := range slice[y] {
			img++
			if img%3 == 0 {
				slice[y][x] = uint8(x ^ 2)

			}
			if img%3 == 1 {
				slice[y][x] = uint8((x + y) / 2)

			}
			if img%3 == 2 {
				slice[y][x] = uint8(x * y)
			}
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}

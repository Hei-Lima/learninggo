package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var slc [][]uint8
	for i := 0; i < dy; i++ {
		var slctwo []uint8
		for j := 0; j < dx; j++ {
			slctwo = append(slctwo, uint8(j^i))
		}
		slc = append(slc, slctwo)
	}

	return slc
}

func main() {
	pic.Show(Pic)
}

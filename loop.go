package main

import(
	"fmt"
)

func main(){
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	j := 0
	for j < 10 {
		//fmt.Println(j)
		j++
	}

	x := [2][2]int {{1,2},{3,4}}
	fmt.Println("Length of \"x\" slice", len(x))
	for y := 0; y < 2; y++ {
		//fmt.Println(x[y])
	}

	// z is a slice
	z := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13}
	fmt.Println("Length of \"z\" slice", len(z))
	for k := 0; k < len(z) ; k++{
		fmt.Println(z[k])
	}
}
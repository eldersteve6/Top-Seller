package main

import "fmt"

func main() {
	sellers := []Seller{
		{Account{"Alice", true}, 100, 5},
		{Account{"Bob", false}, 200, 10},
		{Account{"Charlie", true}, 150, 4},
		{Account{"Ojila", true}, 100, 8},
	}

	fmt.Println(topSeller(sellers))
}

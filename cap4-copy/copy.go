package main

import "fmt"

func main() {
	origem := []int{1, 2, 3, 4, 5}
	destino := make([]int, len(origem))

	copy(destino, origem)

	for i := range destino {
		destino[i] *= 2
	}

	fmt.Println(origem)
	fmt.Println(destino)
}

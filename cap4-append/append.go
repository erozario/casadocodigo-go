package main

import "fmt"

func main() {
	s := []int{23, 24, 25}
	n := []int{22}
	s = append(n, s...)

	fmt.Println(s)

	x := []int{23, 24, 25}
	y := append([]int{22}, x...)
	fmt.Println(y)

	a := []int{11, 12, 13, 16, 17, 18}
	b := []int{14, 15}
	a = append(a[:3], append(b, a[3:]...)...)
	fmt.Println(a)

	c := []int{10, 20, 30, 40, 50, 60}
	c = append(c[:2], c[4:]...)
	fmt.Println(c)
}

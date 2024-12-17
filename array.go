package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println("Length:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b = [...]int{1, 2: 3, 5, 6}
	fmt.Println(b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for k := 0; k < 3; k++ {
			twoD[i][k] = i + k
		}
	}
	fmt.Println(twoD)
	var twoD2 = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(twoD2)
}

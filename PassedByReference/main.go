package main

import "fmt"

func main() {
	sl := []int{
		1, 2, 3, 4, 5,
	}
	fmt.Println(sl) // 1, 2, 3, 4, 5,

	doubleTheSlice(sl)
	fmt.Println(sl) // 2, 4, 6, 8, 10,

	addOneInBack(sl) // it's doesn't work. why?? you can read README
	fmt.Println(sl)  // 2, 4, 6, 8, 10,

	num := 11
	double(&num)
	fmt.Println("\ndouble num :", num)
}

// This func doesn't have
// return value
func doubleTheSlice(s []int) {
	for i := range s {
		// double the value
		s[i] *= 2
	}
}

func addOneInBack(s []int) {
	s = append(s, 1)
}

func double(x *int) {
	*x = *x * 2
}

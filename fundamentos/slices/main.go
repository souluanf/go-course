package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	//slice = append(slice, 10, 2, 5, 6)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//for i := 0; i < 20; i++ {
	//	slice = append(slice, i)
	//	fmt.Println("Len: ", len(slice), " Cap: ", cap(slice))
	//}

	//sliceTest := slice
	//	//slice[0] = 10
	//	//fmt.Println(slice)
	//	//fmt.Println(sliceTest)

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
	}
	fmt.Println(sliceString[:2])
	fmt.Println(sliceString[2:3])
	fmt.Println(sliceString[2:4])
}

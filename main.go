package main

import (
	"fmt"
	"os"
	"strconv"
)

var Matrix [][]int

func main() {
	height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("height not valid")
	}
	width, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic("width not valid")
	}
	Matrix = make([][]int, height)
	for i := range Matrix {
		Matrix[i] = make([]int, width)
	}
	fmt.Println(Matrix)
}

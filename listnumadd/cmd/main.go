package main

import (
	lna "algorithm/listnumadd"
	"fmt"
)

func main() {
	a := []uint32{9}
	b := []uint32{4, 3, 0}
	c := lna.CaculateAdd(a, b)
	fmt.Println(c)
}

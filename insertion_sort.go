package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var slice []int
	slice = make([]int, r.Intn(32))

	fmt.Println("len(slice):", len(slice), ",slice:", slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = r.Intn(256)
	}

	fmt.Println("len(slice):", len(slice), ",slice:", slice)

	for i := 0; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				t := slice[j]
				slice[j] = slice[j-1]
				slice[j-1] = t
			}
		}
	}

	fmt.Println("len(slice):", len(slice), ",slice:", slice)
}

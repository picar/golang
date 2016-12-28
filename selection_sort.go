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

	for i := 0; i < len(slice); i++ {
		slice[i] = r.Intn(256)
	}

	fmt.Println("len(slice):", len(slice), ",slice:", slice)

	start := time.Now().UnixNano()

	for i := 0; i < len(slice); i++ {
		min := i
		for j := i+1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		if i != min {
			t := slice[i]
			slice[i] = slice[min]
			slice[min] = t
		}
	}

	time := time.Now().UnixNano() - start

	fmt.Println("len(slice):", len(slice), ",slice:", slice)
	fmt.Println("TimeNano:", time)
}

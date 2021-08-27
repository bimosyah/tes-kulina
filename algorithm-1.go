package main

import (
	"fmt"
)

func main(){
	// n := 9
	arr := []int{10,20,20,10,10,30,50,10,20}

	total := 0

	for i := 0; i < len(arr); i++ {
		temp := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				temp += 1
			}
		}
		if temp % 2 != 0 {
			total += 1
		}
	}

	fmt.Println(total)
}
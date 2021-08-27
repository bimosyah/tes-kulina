package main

import (
	"fmt"
)

func main(){
	// n := 9
	arr := []string{"U", "D", "D", "D", "U", "D", "U", "U"}

	total := 0
	post_level := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == "U" {
			post_level += 1
		}else if arr[i] == "D"{
			post_level -= 1
			if post_level == -1{
                total += 1
			}
		}
	}

	fmt.Println(total)
}
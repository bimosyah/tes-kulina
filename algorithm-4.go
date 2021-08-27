package main

import (
	"fmt"
)

func main(){
	switches := 0

	for i := 1; i <= 100; i++ {
		if i == 1 {
			switches = 100
			continue
		}
		for j := 1; j <= 100; j++ {

			if j % i == 0 {
				switches -= 1
			}
		}
	}

	if switches < 0 {
		switches = 0
	}

	fmt.Println(switches)
}
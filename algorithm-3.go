package main

import (
	"fmt"
	"strings"
)

func main(){
	// n := 9
	raw_input := "1.345.679"
	input := strings.Replace(raw_input, ".", "", -1)
	stringSlice := strings.Split(input, "")


	for i := 0; i < len(stringSlice); i++ {
		result := ""

		for j := 0; j < len(stringSlice) - i; j++ {
			if j == 0 {
				result += stringSlice[i]
			}else{
				result += "0"
			}
		}
		fmt.Println(result)
	}
}
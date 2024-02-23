package main

import "fmt"

func alpha(s string) int {
	str := []rune(s)
	j := 0
	i := 0
	for i < len(str) {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			j++
		}
		i++
	}
	return j
}

func main() {
	s := "Wello 78 World!    4455 /"
	nb := alpha(s)
	fmt.Println(nb) // Output: 10
}

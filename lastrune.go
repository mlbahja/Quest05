package main

import "fmt"

func lastrune(s string) rune {
	str := []rune(s)
	n := len(s)
	i := 0
	for n > i {
		i++
	}
	return str[n-1]
}
func main() {
	fmt.Println(string(lastrune("hello")))
}

package main

import "fmt"

func first(s string) rune {
	str := []rune(s)
	i := 0
	for i < len(str) {
		i++
	}
	return str[0]
}
func main() {
	fmt.Println(string(first("hello")))
	fmt.Println(string(first("salut")))
	fmt.Println(string(first("ola!")))
}

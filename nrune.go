package main

import "fmt"

func nrune(s string, n int) rune {
	str := []rune(s)
	if n >= len(str) || n < 0 {
		return 0
	}
	return str[n]
}
func main() {
	fmt.Println(string(nrune("hello!", 5)))

}

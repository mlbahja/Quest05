package main

import (
	"fmt"
)

/*func compare(s1 string, s2 string) int {
	var a int
	str1 := []rune(s1)
	str2 := []rune(s2)
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] > str2[j] {
				a = 1
			}
			if str1[i] < str2[j] {
				a = -1
			}
			if str1[i] == str2[j] {
				a = 0
			}
		}
	}
	return a
package main

import "fmt"*/

func compare(s1 string, s2 string) int {
	str1 := []rune(s1)
	str2 := []rune(s2)
	m := len(str1)
	if len(str2) < m {
		m = len(str2)
	}

	for i := 0; i < m; i++ {
		if str1[i] > str2[i] {
			return 1
		} else if str1[i] < str2[i] {
			return -1
		}
	}
	if len(str1) > len(str2) {
		return 1
	} else if len(str1) < len(str2) {
		return -1
	}

	return 0
}

func main() {
	fmt.Println(compare("salut!", "lut!"))
}
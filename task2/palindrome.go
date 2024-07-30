package main

import (
	// "fmt"
	"regexp"
	"strings"
)

func Is_palindrome(s string) bool {
	s = strings.ToLower(s)
	re := regexp.MustCompile("[^a-z0-9]")
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && re.MatchString(string(s[i])) {
			i++
		}
		for i < j && re.MatchString(string(s[j])) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// func main() {
// 	fmt.Println(is_palindrome("A man, a plan, a canal: Panama"))
// 	fmt.Println(is_palindrome("race a car"))
// }

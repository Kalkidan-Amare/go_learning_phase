package main

import (
	"fmt"
	"strings"

)
var panctuations = map[string]bool{".": true, "?": true, "-": true, ",": true}

func word_frequency(text string) map[string]int {
	words := strings.Fields(text)
	word_freq := make(map[string]int)
	for _, word := range words {
		if panctuations[strings.ToLower(string(word[len(word)-1]))] {
			word_freq[word[:len(word)-1]]++
		} else {
			word_freq[word]++
		}
	}
	return word_freq
}

func main(){
	text := "Hello, my name is Jah. What is your name?"
	word_freq := word_frequency(text)
	fmt.Println(word_freq)

	fmt.Println(Is_palindrome("A man, a plan, a canal: Panama"))
	fmt.Println(Is_palindrome("race a car"))
}
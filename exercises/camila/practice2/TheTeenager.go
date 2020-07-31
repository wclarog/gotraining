package practice2

import (
	"fmt"
	"strings"
)

func BobAnswers(text string) string {
	switch {
	case isAnything(text):
		return "Fine. Be that way!"
	case isYellingAQuestion(text):
		return "Calm down, I know what I'm doing!"
	case isANormalQuestion(text):
		return "Sure."
	case isYelling(text):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isYelling(text string) bool {
	return text == strings.ToUpper(text)
}

func isYellingAQuestion(text string) bool {
	return text == strings.ToUpper(text) && strings.Contains(text, "?")
}

func isANormalQuestion(text string) bool {
	return strings.Contains(text, "?")
}

func isAnything(text string) bool {
	return text == ""
}

func TestAnwers() {
	fmt.Println(BobAnswers("How are you?"))
	fmt.Println(BobAnswers("YELL AT HIM"))
	fmt.Println(BobAnswers("WHAT ARE YOU DOING?"))
	fmt.Println(BobAnswers(""))
	fmt.Println(BobAnswers("Good Morning"))
}

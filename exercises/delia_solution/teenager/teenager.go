package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func bobAnswer(question string) string {
	answer := ""
	if IsQuestion(question) {
		if IsUpperCase(question) {
			answer = "Calm down, I know what I'm doing!"
		} else {
			answer = "Sure."
		}
	} else if IsUpperCase(question) {
		answer = "Whoa, chill out!"
	} else if question == "" {
		answer = "Fine. Be that way!"
	} else {
		answer = "Whatever."
	}
	return answer
}

func IsQuestion(question string) bool {
	if len(question) > 0 {
		last := question[len(question)-1:]
		return last == "?"
	}
	return false
}

func IsUpperCase(question string) bool {
	for _, r := range question {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a question for Bob: \n")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)
	fmt.Println("Bob's answer is: ", bobAnswer(question))
}

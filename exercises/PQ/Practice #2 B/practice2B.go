package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isSpecificQuestion(statement string) bool {
	return statement == "How are you?"
}

func isAllCapitals(statement string) bool {
	return strings.ToUpper(statement) == statement
}

func isQuestion(statement string) bool {
	return statement != "" && statement[len(statement) - 1] == '?'
}

func isNothing(statement string) bool {
	return statement == ""
}

func bobAnswer(statement string) string {
	switch {
	case isNothing(statement):
		return "Fine. Be that way!."

	case isSpecificQuestion(statement):
		return "Sure."

	case isAllCapitals(statement):
		return "Whoa, chill out!"

	case isQuestion(statement):
		return "Calm down, I know what I'm doing!"

	default:
		return "Whatever."
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	statement, _ := reader.ReadString('\n')

	for ; statement != "Bye";  {
        answer := bobAnswer(statement)
        fmt.Printf("%s\n", answer)

		statement, _ = reader.ReadString('\n')
	}
}
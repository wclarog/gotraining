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

	case isQuestion(statement) && isAllCapitals(statement):
		return "Calm down, I know what I'm doing!"

	case isSpecificQuestion(statement):
		return "Sure."

	case isAllCapitals(statement):
		return "Whoa, chill out!"

	default:
		return "Whatever."
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	statement, _ := reader.ReadString('\n')
	statement = strings.TrimSpace(statement)

	for ; statement != "Bye";  {
        answer := bobAnswer(statement)
        fmt.Printf("%s\n", answer)

		statement, _ = reader.ReadString('\n')
		statement = strings.TrimSpace(statement)
	}
}
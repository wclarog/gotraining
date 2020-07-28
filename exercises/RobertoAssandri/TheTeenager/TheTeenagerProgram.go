package main

import (
	"fmt"
	"strings"
)

const questionMark = "?"

const question = "Sure."
const yell = "Whoa, chill out!"
const yellQuestion = "Calm down, I know what I'm doing!"
const addressHim = "Fine. Be that way!"
const everythingElse = "Whatever."

const questionType = 1
const yellType = 2
const yellQuestionType = 3
const addressHimType = 4
const everythingElseType = 5

func AnalyzePhrase(phrase string) int{
	if phrase == ""{
		return addressHimType
	}

	lastCharacterStr := string(phrase[len(phrase)-1])
	phraseToUpper := strings.ToUpper(phrase)

	//Yell
	if phraseToUpper == phrase{
		if lastCharacterStr == questionMark{
			return yellQuestionType
		}
		return yellType
	}

	if lastCharacterStr == questionMark{
		return questionType
	}


	return everythingElseType
}

func GetBobResponse(phraseType int) string{
	switch phraseType {
	default:
		return everythingElse
	case 1:
		return question
	case 2:
		return yell
	case 3:
		return yellQuestion
	case 4:
		return addressHim
	}
}

func main() {
	fmt.Println("Addressing Bob without saying anything")
	fmt.Println("Answer should be: 'Fine. Be that way!'")
	fmt.Println("Result is " + GetBobResponse(AnalyzePhrase("")))
	fmt.Println("")

	fmt.Println("Yelling at Bob")
	fmt.Println("Answer should be: 'Whoa, chill out!'")
	fmt.Println("Result is " + GetBobResponse(AnalyzePhrase("YELL")))
	fmt.Println("")

	fmt.Println("Yelling Bob a question")
	fmt.Println("Answer should be: 'Calm down, I know what I'm doing!'")
	fmt.Println("Result is " + GetBobResponse(AnalyzePhrase("WHAT?")))
	fmt.Println("")

	fmt.Println("Asking Bob a question")
	fmt.Println("Answer should be: 'Sure.'")
	fmt.Println("Result is " + GetBobResponse(AnalyzePhrase("How are you?")))
	fmt.Println("")

	fmt.Println("Saying anything else to Bob")
	fmt.Println("Answer should be: 'Whatever.'")
	fmt.Println("Result is " + GetBobResponse(AnalyzePhrase("Hello")))
}


//The teenager
//
//Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//
//Bob answers 'Sure.' if you ask him a question, such as "How are you?".
//
//He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
//
//He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
//
//He says 'Fine. Be that way!' if you address him without actually saying anything.
//
//He answers 'Whatever.' to anything else.
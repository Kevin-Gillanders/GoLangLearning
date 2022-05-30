package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)


var DEBUG bool = false
var minLen int = 4
var maxLen int = 6
var ApiKey string = "b3a32551a3367171d32c2694d1450ac337ee516eadf61a125"





func compareGuess(guess string, answer string) bool{
	if DEBUG{
		fmt.Printf("~~~DEBUG compareGuess~~~\n guess : %v \nanswer : %v\n", guess, answer)
	}
	guess = strings.ToLower(guess)
	answer = strings.ToLower(answer)
	
	incorrectLetter           := color.New(color.BgWhite).Add(color.FgBlack)
	correctLetter             := color.New(color.BgHiGreen)
	correctLetterWrongPlace   := color.New(color.BgHiYellow).Add(color.FgBlack)

	correctLetterCount := 0
	for pos, char := range guess {
		currentLetter := color.New()

		if char == rune(answer[pos]){
			currentLetter = correctLetter
			correctLetterCount++

		} else if strings.ContainsRune(answer, char) {
			currentLetter = correctLetterWrongPlace

		} else {
			currentLetter = incorrectLetter

		}

		currentLetter.Printf("[%c]", char)

	} 
	fmt.Println()
	return correctLetterCount == len(answer)

}

func GetNewWord() string{

	endPoint := fmt.Sprintf("https://api.wordnik.com/v4/words.json/randomWord?hasDictionaryDef=true&maxCorpusCount=-1&minDictionaryCount=1&maxDictionaryCount=-1&minLength=%v&maxLength=%v&api_key=%v",
						    minLen,
						    maxLen,
						    ApiKey)


	response, error := http.Get(endPoint)


	if error != nil {
		fmt.Println(error)
	}

	if DEBUG {
		
		fmt.Printf("This is the url being hit %v\n\n", endPoint)
		fmt.Println("~~~DEBUG GetNewWord~~~\n Writing out response :")
		fmt.Println(response)

	}


	defer response.Body.Close()

	jsonBody := json.NewDecoder(response.Body)

	var value Message

	if err := jsonBody.Decode(&value); err == io.EOF {	
		log.Fatal(err)			

	} else if err != nil {
		log.Fatal(err)

	}


	if DEBUG {
		fmt.Println("~~~DEBUG GetNewWord~~~\n Writing out value :")
		fmt.Println("JSON Decode")
		fmt.Println("=======\n", value) 
		fmt.Printf("Name : %v | value : %v |\n", value.Id, value.Word) 
		fmt.Println("=======") 

	}


	return value.Word
}


func PrintIntro(amountOfBoxes int){
	c := color.New(color.FgHiGreen).Add(color.BgBlack)
	c.Println("Please enter a guess")


	c = color.New(color.BgWhite).Add(color.FgBlack)
	c.Println(strings.Repeat("[ ]", amountOfBoxes))
}


type Message struct{
	Id int 
	Word string 
}



func main() {
	answer := GetNewWord()
	guesses := 0
	success := false

	reader := bufio.NewReader(os.Stdin)


	PrintIntro(len(answer))


	var guess string

	for{
		for (guesses < 6 && !success ){

			guess, _ = reader.ReadString('\n')
			guess = strings.ToLower(strings.Replace(guess, "\r\n", "", -1))

			if len(guess) != len(answer){
				fmt.Println("The word you entered is not the correct lenght, please guess again.")
				continue
			}
			correctGuess := compareGuess(guess, answer)

			guesses += 1
			if DEBUG{
				fmt.Printf("~~~DEBUG main~~~ \nguess : %v \nwas it right\n%v\n", guesses, answer)
			}

			if(correctGuess){
				fmt.Printf("Congratulations, you guessed the Gordle in %d \n", guesses)
				success = true
				break
			}
		}


		if !success {
			fmt.Printf("You have failed to guess todays Gordle, it was : %s\n", answer)
		}
		
		fmt.Println("Would you like to play again? [y/N]")

		playAgain, _ := reader.ReadString('\n')
		playAgain = strings.ToLower(strings.Replace(playAgain, "\r\n", "", -1))

		if(playAgain == "n"){
			break
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dictionary = []string{
	"Semih",
	"Abaza",
	"Kalem",
	"Sampanya",
	"Raki",
	"Pokemon",
	"Pikachu",
	"Latios",
	"Regice",
}

func main() {
	/*
		1) Choose a Word
		2) Print the UI
		3) Read the user input
		4) Validate the user input ( It has to be
	*/

	game := newGame()
	reader := bufio.NewReader(os.Stdin)
	isCompleted := false
	for game.faults < 9 && !isCompleted {
		isCompleted = true
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}
		game.newGuess(strings.ToLower(text))
		fmt.Println(game)

		for _, val := range game.currentLetters {
			if val[0] == "_"[0] {
				isCompleted = false
			}
		}
		if isCompleted {
			fmt.Println("Completed!")

		}
	}

}

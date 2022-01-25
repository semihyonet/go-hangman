package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type game struct {
	word           []string
	usedLetters    []string
	currentLetters []string
	faults         int
	corrects       int
}

func newGame() game {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	newWord := strings.ToLower(dictionary[r1.Intn(len(dictionary))])
	currentLetters := make([]string, len(newWord))
	for i := range currentLetters {
		currentLetters[i] = "_"
	}
	return game{
		word:           stringToArr(newWord),
		usedLetters:    []string{},
		currentLetters: currentLetters,
		faults:         0,
		corrects:       0,
	}
}

func (g game) validate(s string) bool {

	if len(s) > 2 {
		panic("TOO_LONG" + string(len(s)))
	}
	for _, letter := range g.usedLetters {
		if letter == s {
			return false
		}
	}

	return true
}

func (g *game) isLetterInWord(s string) bool {
	isInWord := false
	for i, val := range g.word {

		if val[0] == s[0] {
			if !isInWord {
				isInWord = true
			}
			g.currentLetters[i] = val
			g.corrects++
		}
	}
	return isInWord
}

// Checks if the word is in the letter or not
func (g *game) newGuess(s string) bool {
	if (*g).validate(s) {
		g.usedLetters = append(g.usedLetters, s)
		if !g.isLetterInWord(s) {
			g.faults++
			return false
		}
		return true
	} else {
		fmt.Println("Word is already guessed")
		return false
	}
}

func (g game) String() string {
	data, err := ioutil.ReadFile("./states/state" + strconv.Itoa(g.faults) + ".txt")
	if err != nil {
		panic(err)
	}

	s := "Game results\n "

	for _, val := range g.currentLetters {
		s += val
	}
	s += "\n"

	s += string(data)

	s += "\n"
	s += "Correct : " + strconv.Itoa(g.corrects) + "\n"
	s += "False : " + strconv.Itoa(g.faults)

	return s
}

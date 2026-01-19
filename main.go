package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./words.txt")
	if err != nil {
		panic(err)
	}

	wordArr := strings.Split(string(file), "\n")

	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(wordArr))))
	randIdx := int(n.Int64())
	answer := strings.TrimSpace(wordArr[randIdx])

	fmt.Println("Hangman🌳")

	word := strings.Repeat("_", len(answer))
	// fmt.Printf("Word: %s, %d\n", answer, len(answer))

	guessedRight := ""
	guessedWrong := ""
	wrongLetters := ""
	count := 0
	stickMan := ``

	for {
		guess := ""
		stickMan = displayStickMan(count)
		fmt.Printf("\nWord: %s\n", word)

		if count == 6 {
			fmt.Printf("You lost the game😒. The correct answer was: %s\n", answer)
			fmt.Println(stickMan)
			return
		}

		if word == answer && count < 6 {
			fmt.Println("You won the game🤩")
			fmt.Println(stickMan)
			return
		}

		fmt.Println(stickMan)
		fmt.Println("Guess a letter:")
		fmt.Scan(&guess)
		clearScreen()

		if len(guess) > 2 {
			fmt.Println("\nToo many letters.")
			if len(wrongLetters) > 0 {
				fmt.Printf("Wrong letters: %s\n", wrongLetters)
			}
			continue
		}

		if strings.Contains(guessedRight, guess) || strings.Contains(guessedWrong, guess) {
			fmt.Println("\nThis letter has already been guessed.")
			if len(wrongLetters) > 0 {
				fmt.Printf("Wrong letters: %s\n", wrongLetters)
			}
			continue
		}

		if strings.Contains(answer, guess) {
			fmt.Println("\nCorrect! ✅")
			guessedRight += guess
		} else {
			fmt.Println("\nWrong ❌")
			count++
			guessedWrong += guess
		}

		// Wrong letters
		wrongLetters = displayWrong(guessedWrong, wrongLetters)
		if len(wrongLetters) > 0 {
			fmt.Printf("Wrong letters: %s\n", wrongLetters)
		}

		// Correctly guessed letters
		word = displayRight(guess, answer, word)
		stickMan = displayStickMan(count)

	}

}

func displayRight(s string, answer string, oldWord string) string {
	arr := strings.Split(oldWord, "")
	for i := 0; i < len(answer); i++ {
		if string(answer[i]) == s {
			arr[i] = s
		}
	}
	newWord := strings.Join(arr, "")
	return newWord
}

func displayWrong(guessedWrong string, wrongLetters string) string {
	bytes := []byte(guessedWrong)
	arr := strings.Split(string(bytes), "")
	wrongLetters = strings.Join(arr, ",")
	return wrongLetters
}

func displayStickMan(count int) string {

	stickMan := []string{
		` 
                  
                 
                 
                 
        
		`,
		` 
                  
                 
                 
                 
        =========
		`,
		` 
                |
                |
                |
                |
        =========
		`,
		`
            +---+ 
                |
                |
                |
                |
        =========
		`,
		`
            +---+ 
            |   |		
                |
                |
                |
        =========
		`,
		`
            +---+ 
            |   | 
           🥲    |
           /|\  |
                |
                |
        =========
		`,
		`
            +---+ 
            |   |
           😵   |
           /|\	|
           / \  |
                |
        =========
		`,
	}

	return stickMan[count]
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

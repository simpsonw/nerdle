package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var solution = []rune("slice")

const emptySpace = " _ "
const maxAttempts = 5

func main() {
	var gameOver = false
	var numAttempts = 0
	var result []string
	for i := 0; i < len(solution); i++ {
		result = append(result, emptySpace)
	}
	fmt.Printf("Welcome to Nerdle!  Try to guess the word in %d attempts\n", maxAttempts)
	s := bufio.NewScanner(os.Stdin)
	for !gameOver {
		fmt.Printf("\n%s\n\n> ", strings.Join(result, ""))
		s.Scan()
		guess := strings.Trim(s.Text(), " ")
		if len(guess) != 5 {
			fmt.Println("Invalid guess")
			continue
		}

		if numAttempts == maxAttempts {
			fmt.Printf("Game over!  The correct answer was: %s\n", string(solution))
			gameOver = true
		} else {
			runeGuess := []rune(guess)
			for k, v := range runeGuess {
				if v == solution[k] {
					result[k] = fmt.Sprintf(" %c ", v)
				} else {
					result[k] = emptySpace
				}
			}
			for k, v := range runeGuess {
				if result[k] == emptySpace &&
					strings.Contains(string(solution), string(v)) &&
					!strings.Contains(strings.Join(result, ""), string(v)) {
					result[k] = fmt.Sprintf("[%c]", v)
				}
			}
		}

		if guess == string(solution) {
			fmt.Printf("\n%s\n\nCorrect!\n", strings.Join(result, ""))
			gameOver = true
		}

	}
}

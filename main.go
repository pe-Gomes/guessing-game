package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_GUESSES int64 = 100
	MAX_TRIES   int8  = 10
)

func main() {
	fmt.Println("Welcome to the guessing game!")
	fmt.Sprintln("A random number will the sorted and you have to guess in %i tries.Please pick a number between 1 and %i.", MAX_TRIES, MAX_GUESSES)

	x := rand.Int64N(MAX_GUESSES + 1)
	s := bufio.NewScanner(os.Stdin)

	guesses := [10]int64{}
	for i := range guesses {
		fmt.Println("What is your guess?")
		s.Scan()

		guess := s.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("Too low! The number is greater than", guessInt)
		case guessInt > x:
			fmt.Println("Too high! The number is less than", guessInt)
		case guessInt == x:
			fmt.Printf("\nCongratulations 🎉 You guessed the number %d\n"+
				"\nYou got it in %d tries.\n"+
				"\nYour guesses were: %v\n",
				x, i+1, guesses[:i])
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf("\nToo bad, you runned out of tries 😞. The number was %d\n"+
		"\nYou had %d tries.\n"+
		"\nYour guesses were: %v\n"+
		"\nBetter luck next time!\n",
		x, MAX_TRIES, guesses)
}

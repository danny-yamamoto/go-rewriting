package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Guess the number!")
	secretNumber := r.Intn(100)
	fmt.Printf("The secret number is %d\n", secretNumber)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please input your guess.")
		scanner.Scan()
		guessStr := scanner.Text()
		guess, err := strconv.Atoi(guessStr)
		if err != nil {
			continue
		}
		fmt.Printf("You guessed: %d\n", guess)
		if guess < secretNumber {
			fmt.Println("Too small")
		} else if guess > secretNumber {
			fmt.Println("Too big")
		} else {
			fmt.Println("You win")
			break
		}
	}
}

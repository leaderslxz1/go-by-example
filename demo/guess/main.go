package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(maxNum)
	fmt.Println(secretNum)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err in reading")
			//return
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		//fmt.Println(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("not a int")
			//return
			continue
		}
		//fmt.Println("Input guess:", guess)
		switch {
		case guess > secretNum:
			fmt.Println(guess, "> target. Plz try again.")
		case guess < secretNum:
			fmt.Println(guess, "< target. Plz try again.")
		default:
			fmt.Println(guess, "equals secret number. Congrats!")
			return
		}
	}
}

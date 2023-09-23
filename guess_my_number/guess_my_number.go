package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	var s string
	for {
		println("enter last number: ")
		var max int
		var err error
		for {
			fmt.Scanln(&s)
			max, err = strconv.Atoi(s)
			if err == nil {
				break
			}
			fmt.Printf("%T, %v\ninput number again:\n", s, s)
		}
		println("you will guess from 1 to " + s)
		randNumber := rand.Intn(max) + 1

		tryCount := 0
		for {
			tryCount++
			println("guess: ")
			fmt.Scanln(&s)
			guess, _ := strconv.Atoi(s)
			if randNumber == guess {
				fmt.Println("congratulation. you found it in", tryCount, "tries.")
				break
			} else if randNumber > guess {
				println("my number is bigger")
			} else {
				println("my number is lower")
			}

		}
		println("do you want play again (press y or 0 for yes): ")
		fmt.Scanln(&s)
		if s != "y" && s != "0" {
			break
		}

	}
	//2

}

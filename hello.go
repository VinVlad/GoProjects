package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100) + 1
	fmt.Println("Я выбираю число до 100.")
	fmt.Println("Число выбрано.")

	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("Осталось", 10-guesses, "попыток")
		fmt.Print("Введите число: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Загаданное число меньше)")
		} else if guess < target {
			fmt.Println("Загаданное число больше(")
		} else {
			success = true
			fmt.Println("Вы угадали!")
			break
		}
	}

	if !success {
		fmt.Println("Вы не угадали! Было загадано число", target)
	}
}

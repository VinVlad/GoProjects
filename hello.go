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
	fmt.Println("Я выбираю число до 100")
	fmt.Println("Число выбрано")

	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

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
}

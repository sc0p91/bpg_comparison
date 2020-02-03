package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	var input = ""
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())

	fmt.Println("Welcome to Type-Runner! Hit enter to begin the race...")
	startTime := time.Now().Unix()
	fmt.Scanln()

	fileBytes, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	for i := 1; i <= 5; i++ {
		word := strings.ToLower(sliceData[rand.Intn(len(sliceData))])
		for {
			if word == input {
				break
			}
			fmt.Print("Type ", word, ": ")
			input, _ = reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
		}

	}
	actualTime := time.Now().Unix() - startTime
	fmt.Println("You've done it in ", actualTime, "seconds")
}

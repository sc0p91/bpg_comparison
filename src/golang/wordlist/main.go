package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

//go:generate go run main.go
func main() {
	fileBytes, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := bytes.Buffer{}
	buffer.WriteString("package main\n")
	buffer.WriteString("var wordList = `")
	buffer.Write(fileBytes)
	buffer.WriteString("`")

	ioutil.WriteFile("../wordlist.go", buffer.Bytes(), 0644)
}

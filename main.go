package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Print("Enter text: ")

	// ? Read string from console
	text, err := reader.ReadString('\n')
	checkError(err)

	// ? Remove \n from string
	text = strings.TrimRight(text, "\n")

	// ? Convert string to int
	num, err := strconv.Atoi(text)
	checkError(err)


	fmt.Println(num + 10)
}
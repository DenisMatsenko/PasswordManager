package main

import (
	// "bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	// "strings"
	// "strconv"
	// "strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello!")

	// Test()

	CreatePassword()

	// // ? Read string from console
	// var text int
	// _, err := fmt.Scanf("%d", &text)
	// checkError(err)

	// fmt.Print(num)

	// fmt.Println(num + 10)
}

type PasswordParameters struct {
	passwordSecure int
	passwordLength int
}

type Password struct {
	word        string
	description string
}

func GeneratePassword(pp PasswordParameters) string {
	// * changable password chars
	arr := []string{"abcdefghijklmnopqrstuvwxyz", "0123456789", "!@#$%&*_+-,.?~"}

	// ? Set end point for random
	endPoint := 0
	switch pp.passwordSecure {
	case 1:
		endPoint = len(arr[0])
	case 2:
		endPoint = len(arr[0]) + len(arr[1])
	case 3:
		endPoint = len(arr[0]) + len(arr[1]) + len(arr[2])
	}

	// ? Randomly generate every char in password
	passwordChars := strings.Join(arr, "")
	var password string = ""
	for i := 0; i < pp.passwordLength; i++ {
		password = fmt.Sprintf("%s%c", password, passwordChars[rand.Intn(endPoint)])
	}

	return password
}

func CreatePassword() {
	// ? Init password parameters
	pp := PasswordParameters{passwordSecure: 3, passwordLength: 20}
	const countOfPasswords int = 5

	// ? Generate passwords
	var passwords [countOfPasswords]string
	for index := range passwords {
		passwords[index] = GeneratePassword(pp)
	}

	var password Password
	password.word = passwords[0]
	password.description = "pw description"

	SaveLocal(password)
}

func SaveLocal(password Password) {

	// ? Open or create file
	file, err := os.OpenFile("MyPasswords.csv", os.O_APPEND|os.O_CREATE, 0644)
	checkError(err)

	// ? Read data from file
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	checkError(err)
	file.Close()

	// ? Add new data
	data = append(data, []string{password.word, password.description})

	// ? Open file for writing
	file, err = os.OpenFile("MyPasswords.csv", os.O_WRONLY|os.O_CREATE, 0644)
	checkError(err)

	// ? Write data to file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.WriteAll(data)
	checkError(err)
}

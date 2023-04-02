package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type GeneratorParameters struct {
	passwordSecure int
	passwordLength int
	passwordsCount int
}
type Password struct {
	word        string
	description string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Hello!")

	for {
		fmt.Println("\n\n\nChoose option:\n 1. Generate password\n 2. Show passwords")
		var option int
		_, err := fmt.Scanf("%d", &option)
		checkError(err)

		// ? Choose option
		if option == 1 {
			CreatePassword()
			break
		} else if option == 2 {
			ShowPasswords()
			break
		} else {
			fmt.Println("Wrong option!")
		}
	}
}

// * Core functions
func CreatePassword() {
	passwordDescription := DescriptionInput()
	passwordSecure := SecureInput()
	passwordLength := LengthInput()
	passwordsCount := PassvordsCountInput()

	// ? Init password parameters
	gp := GeneratorParameters{passwordSecure: passwordSecure, passwordLength: passwordLength, passwordsCount: passwordsCount}

	// ? Generate passwords
	passwords := make([]string, gp.passwordsCount)
	for index := range passwords {
		passwords[index] = GeneratePassword(gp)
		fmt.Print(index+1, " ")
		fmt.Println(passwords[index])
	}

	// ? Choose password
	passwordNum := ChoosePasswordInput(gp)

	// ? Print new password
	fmt.Printf("\nCreated new password: %s\n", passwords[passwordNum-1])

	// ? Save new password
	var password Password
	password.word = passwords[passwordNum-1]
	password.description = passwordDescription
	SaveLocal(password)
}
func ShowPasswords() {
	// ? Get and Print data
	for _, row := range GetDataFromFile() {
		fmt.Println(row[0], " - ", row[1])
	}
}
func GeneratePassword(pp GeneratorParameters) string {
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

// * File functions
func GetDataFromFile() [][]string {
	// ? Open or create file
	file, err := os.OpenFile("MyPasswords.csv", os.O_APPEND|os.O_CREATE, 0644)
	checkError(err)
	defer file.Close()

	// ? Read data from file
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	checkError(err)

	return data
}
func SaveLocal(password Password) {
	// ? Read old data
	data := GetDataFromFile()

	// ? Add new data
	data = append(data, []string{password.word, password.description})

	// ? Open file for writing
	file, err := os.OpenFile("MyPasswords.csv", os.O_WRONLY|os.O_CREATE, 0644)
	checkError(err)

	// ? Write data to file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.WriteAll(data)
	checkError(err)
}

// * Interface"" functions
func DescriptionInput() string {
	var passwordDescription string
	// ? Read password description from console
	for {

		fmt.Print("Enter password description: ")
		_, err := fmt.Scanf("%s", &passwordDescription)
		checkError(err)

		if len(passwordDescription) > 0 {
			break
		} else {
			fmt.Println("Wrong option!")
		}
	}
	return passwordDescription
}
func SecureInput() int {
	var passwordSecure int
	// ? Read password secure from console
	for {
		fmt.Print("Choose password secure:\n 1. Low\n 2. Medium\n 3. High\nEnter: ")

		_, err := fmt.Scanf("%d", &passwordSecure)
		checkError(err)

		if passwordSecure >= 1 && passwordSecure <= 3 {
			break
		} else {
			fmt.Println("Wrong option!")
		}
	}
	return passwordSecure
}
func LengthInput() int {
	var passwordLength int
	// ? Read password length from console
	for {
		fmt.Print("Enter password length: ")
		_, err := fmt.Scanf("%d", &passwordLength)
		checkError(err)

		if passwordLength > 0 {
			break
		} else {
			fmt.Println("Wrong option!")
		}
	}
	return passwordLength
}
func PassvordsCountInput() int {
	var passwordsCount int
	// ? Read password count from console
	for {
		fmt.Print("Enter password count: ")
		_, err := fmt.Scanf("%d", &passwordsCount)
		checkError(err)

		if passwordsCount > 0 {
			break
		} else {
			fmt.Println("Wrong option!")
		}
	}
	return passwordsCount
}
func ChoosePasswordInput(gp GeneratorParameters) int {
	var passwordNum int
	// ? Choose password
	for {
		fmt.Print("Enter password number: ")
		_, err := fmt.Scanf("%d", &passwordNum)
		checkError(err)

		if passwordNum > 0 && passwordNum <= gp.passwordsCount {
			break
		} else {
			fmt.Println("Wrong option!")
		}
	}

	return passwordNum
}

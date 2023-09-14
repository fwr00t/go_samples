package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func generatePassword(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	source := rand.NewSource(time.Now().UnixNano())
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = characters[rand.New(source).Intn(len(characters))]
	}

	return string(password)
}

func main() {
	fmt.Println("Welcome to the password generator!")

	for {
		var lengthInput string
		fmt.Print("How long would you like your password to be? ")
		_, err := fmt.Scan(&lengthInput)

		if err != nil {
			fmt.Println("Error: An error occurred while reading your input.")
			return
		}

		length, err := strconv.Atoi(lengthInput)
		if err != nil || length <= 0 {
			fmt.Println("Error: Password length should be a positive number.")
			continue
		}

		password := generatePassword(length)

		fmt.Printf("\nYour password is: %s\n", password)
		fmt.Printf("The length of your password is: %d\n", len(password))

		for {
			var choice string
			fmt.Print("Would you like to generate another password? (yes/no): ")
			fmt.Scan(&choice)
			choice = strings.ToLower(choice)

			if choice == "yes" || choice == "y" {
				break
			} else if choice == "no" || choice == "n" {
				fmt.Println("Thank you for using this program!")
				return
			} else {
				fmt.Println("Error: Please enter 'yes' or 'no'.")
			}
		}
	}
}

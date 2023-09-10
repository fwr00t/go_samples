package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var length int
	var choice string
	var password string
	var characters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+")

	fmt.Println("Welcome to the password generator!")

	for {
		fmt.Println("How long would you like your password to be?")
		_, err := fmt.Scan(&length)
		if err != nil {
			fmt.Println("Error:", err)
			clearInputBuffer()
			continue // Continue the loop if numeric input is not provided
		}

		rand.Seed(time.Now().UnixNano())

		for i := 0; i < length; i++ {
			password += string(characters[rand.Intn(len(characters))])
		}

		fmt.Printf("\n")
		fmt.Println("Your password is: ", password)
		fmt.Println("The length of your password is: ", len(password))

		// Read the user's choice without using fmt.Scanln
		fmt.Print("Would you like to generate another password? (yes/no): ")
		_, _ = fmt.Scan(&choice)

		// Convert the choice to lowercase and check against accepted values
		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice == "yes" || choice == "y" {
			password = "" // Reset the password for the next iteration
		} else if choice == "no" || choice == "n" {
			fmt.Println("Thank you for using this program!")
			break // Exit the loop if the user does not want to generate another password
		} else {
			fmt.Println("Error: Please enter Y or N")
		}
	}
}

func clearInputBuffer() {
	var dummy string
	fmt.Scanln(&dummy)
}

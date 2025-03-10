package main

import (
	"fmt"
	"regexp"
)

var currentUser string

type Userdata struct {
	Username string
	Password string
}

var users = []Userdata{}

func greetUser() {
	var userLoginChoice int

	fmt.Println("Hello user, welcome to our Book Application.")
	fmt.Println("Please choose 1 for log-in or 2 for registration.")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&userLoginChoice)

	if userLoginChoice == 1 {
		handleLogin()
	} else if userLoginChoice == 2 {
		handleRegistration()
	} else {
		fmt.Println("Invalid choice. Please restart the program.")
	}
}

func handleLogin() {
	var username, password string

	for {
		fmt.Println("--------------------------------------------------")
		fmt.Print("Enter your username: ")
		fmt.Scan(&username)

		if len(username) == 0 {
			fmt.Println("Username cannot be empty. Please try again.")
			continue
		}

		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		if len(password) < 6 {
			fmt.Println("Password must be at least 6 characters long. Please try again.")
			continue
		}

		for _, user := range users {
			if user.Username == username && user.Password == password {
				fmt.Println("Login successful!")

				currentUser = username
				runAPP(currentUser)
				return
			}
		}

		fmt.Println("Invalid username or password. Please try again.")
	}
}

func handleRegistration() {
	var newUsername, newPassword, confirmedPassword string

	for {
		fmt.Println("--------------------------------------------------")
		fmt.Print("Enter your username: ")
		fmt.Scan(&newUsername)

		if len(newUsername) == 0 {
			fmt.Println("Username cannot be empty. Please try again.")
			continue
		}

		if !isValidUsername(newUsername) {
			fmt.Println("Username must be at least 3 characters long and cannot contain spaces. Please try again.")
			continue
		}

		fmt.Print("Enter your new password: ")
		fmt.Scan(&newPassword)

		if len(newPassword) < 6 {
			fmt.Println("Password must be at least 6 characters long. Please try again.")
			continue
		}

		fmt.Print("Confirm your new password: ")
		fmt.Scan(&confirmedPassword)

		if newPassword != confirmedPassword {
			fmt.Println("Passwords do not match. Please try again.")
		} else {

			newUser := Userdata{
				Username: newUsername,
				Password: newPassword,
			}
			users = append(users, newUser)
			fmt.Println("Registration successful!")

			handleLogin()
			break
		}
	}
}

func isValidUsername(username string) bool {
	regex := "^[a-zA-Z0-9_]+$"
	matched, _ := regexp.MatchString(regex, username)
	return matched && len(username) >= 3
}

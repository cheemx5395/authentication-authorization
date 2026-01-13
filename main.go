package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	name     string
	password []byte
}

var users = make([]user, 0)

func main() {
	// Take User input for username and password
	var username string
	var password string
	fmt.Println("Enter your Username: ")
	fmt.Scan(&username)
	fmt.Println("Enter your password")
	fmt.Scan(&password)

	// Hash the password and store in in-memory array
	cost := bcrypt.DefaultCost
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Printf("Error hashing the password: %v\n", err)
	}

	users = append(users, user{name: username, password: hashedPass})

	for _, user := range users {
		fmt.Printf("Username: %s, Password: %s\n", user.name, user.password)
	}

	// Ask the user to log-in using username and password
	var name string
	var pass string
	fmt.Println("Enter your username:")
	fmt.Scan(&name)
	fmt.Println("Enter your password:")
	fmt.Scan(&pass)

	err = bcrypt.CompareHashAndPassword(users[0].password, []byte(pass))
	if err != nil {
		fmt.Println("Invalid Credentials")
		return
	}

	// Adding MFA layer
	fmt.Println("Login Successful")
	otp, err := GenerateOTP(6)
	if err != nil {
		fmt.Printf("Error while creating the otp: %v", err)
		return
	}
	fmt.Printf("Generated OTP: %s\n", otp)

	var userEnteredOTP string
	fmt.Println("Enter your OTP: ")
	fmt.Scan(&userEnteredOTP)

	if userEnteredOTP == otp {
		fmt.Println("Welcome, your entering Nirvana!")
	} else {
		fmt.Println("Your access is restricted in Nirvana!")
	}
}

func GenerateOTP(length int) (string, error) {
	const otpChars = "0123456789"
	otpCharsLength := len(otpChars)

	buffer := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(otpCharsLength)))
		if err != nil {
			return "", err
		}

		buffer[i] = otpChars[num.Int64()]
	}

	return string(buffer), nil
}

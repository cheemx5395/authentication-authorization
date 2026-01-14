package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
)

type userRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type user struct {
	Name     string `json:"name"`
	Password []byte `json:"password"`
}

type response struct {
	Message string `json:"message"`
}

var User = &user{}
var OTP = ""

const port = ":8080"

func main() {
	mux := &http.ServeMux{}
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	mux.HandleFunc("POST /signup", signupHandler)
	mux.HandleFunc("POST /login", loginHandler)
	mux.HandleFunc("POST /otp", otpHandler)

	fmt.Printf("Server running at port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error in running server: %v", err)
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

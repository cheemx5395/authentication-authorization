package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	var res response
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding the request body: %v\n", err)
		res.Message = "Bad Request"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	if req.Name == "" || req.Password == "" {
		log.Println("Fields left empty by user")
		res.Message = "Both Fields are Necessary!"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	cost := bcrypt.DefaultCost
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), cost)
	if err != nil {
		log.Printf("Error hashing the Error in hashing the passwordpassword: %v\n", err)
		res.Message = "Internal Server Error"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}

	User.Name = req.Name
	User.Password = hashedPass

	body, _ := json.Marshal(User)
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	var res response
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error parsing the request: %v\n", err)
		res.Message = "Bad Request"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	if req.Name == "" || req.Password == "" {
		log.Println("Fields left empty by user")
		res.Message = "Both Fields are Necessary!"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	err = bcrypt.CompareHashAndPassword(User.Password, []byte(req.Password))
	if err != nil {
		log.Println("Name and Password doesn't match!")
		res.Message = "Unauthorized Access!"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(body)
		return
	}

	res.Message = "User Logged in Successfully!"

	// send otp
	otp, _ := GenerateOTP(6)
	log.Printf("Generated OTP: %s\n", otp)
	OTP = otp

	body, _ := json.Marshal(res)
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func otpHandler(w http.ResponseWriter, r *http.Request) {
	var res response
	req := struct {
		OTP string `json:"otp"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error parsing the request: %v", err)
		res.Message = "Bad Request"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	if req.OTP == "" {
		log.Println("No OTP Entered in request")
		res.Message = "Bad Request"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	if req.OTP != OTP {
		log.Println("Your access is restricted in Nirvana!")
		res.Message = "Unauthorized Access"
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(body)
		return
	}

	log.Println("Welcome, you are entering Nirvana!")
	res.Message = "Welcome to Nirvana"
	body, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

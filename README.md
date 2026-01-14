````markdown
# Assignment 01 – Solution in Golang

This project is a simple authentication flow implemented in **Golang**, demonstrating user signup, login, and OTP-based verification using **in-memory storage**.

---

## How to Run

```bash
go run .
````

The server starts on `http://localhost:8080` and maintains user data **in memory** for the duration of its execution.

## Request–Response Handling

### 1. Signup User

**POST `/signup`**

**Request**

```json
{
  "name": "cheems",
  "password": "cheems@123"
}
```

**Response**

```json
{
  "name": "cheems",
  "password": "JDJhJDEwJHQ3OGFsNUdLODhLTzVHQ0g0djFWOC43bFF3cm9PRTlvUnZXalZieXRyVWtOUFlSVEhuME5p"
}
```

### 2. Login User

**POST `/login`**

**Request**

```json
{
  "name": "cheems",
  "password": "cheems@123"
}
```

**Response**

```json
{
  "message": "User Logged in Successfully!"
}
```

During login, a **One-Time Password (OTP)** is generated and logged to the server console.
This OTP is required for the next step.

### 3. OTP Verification

**POST `/otp`**

**Request**

```json
{
  "otp": "467448"
}
```

**Response**

```json
{
  "message": "Welcome to Nirvana"
}
```

## Notes

* User data is stored in memory while the server is running.
* Passwords are securely hashed before storage.
* Accurate HTTP status codes are returned for all responses.
* Responses are properly encapsulated and structured.
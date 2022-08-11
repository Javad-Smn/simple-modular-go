# simple-modular-go
Simple modular server app for golang with JWT authentication

## Installation
```git clone https://github.com/Javad-Smn/simple-modular-go.git```

## Run
```go run main.go``` (At root of project)

## Active routes
### 1. Sign up
>http://localhost:8080/v1/auth/sign-up
#### Body of request
{
    "name": "Karim",
    "age": 28,
    "password": "4321"
}


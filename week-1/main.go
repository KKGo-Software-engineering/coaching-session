package main

import (
	"fmt"
	"github.com/KKGo-Software-engineering/coaching-session/week-1/calculator"
)

func main() {
	result := calculator.Add(1, 2)
	fmt.Println(result)
}

// Uncomment the following code to run golangci
//func Login(username string, password string) bool {
//	secret := "Nwu0FElpq2%12LqjRALc"
//	if username == "admin" && password == secret {
//		return true
//	}
//
//	return false
//}

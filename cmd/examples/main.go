package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateUser()
	fmt.Println(err)
	err = CreateOrg()
	fmt.Println(err)
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		fmt.Printf("integer: %d string: %s any-value: %v", 123, "a-string",
			"another string\n")
		return fmt.Errorf("create user: %w", err)
	}
	// ...continue on
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}

	return nil
}

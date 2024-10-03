package model

import "fmt"

type User struct{}

func (user *User) CreateUser(username string, password string) (string, error) {
	fmt.Printf("create user %v with password %v\n", username, password)
	commit()

	// return "", fmt.Errorf("database error")
	return "inserted", nil
}

func commit() {
	fmt.Println("commit")
}
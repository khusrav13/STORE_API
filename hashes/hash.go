package hashes

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashIt(Password string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.MinCost)
	if err != nil {
		fmt.Printf("Error can't hash password %e", err)
	}
	return string(password)
}

//////////////Check it then :TODO.
func CheckHash(Password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	if err != nil {
		fmt.Printf("Error can't hash password %e", err)
	}
	return err == nil
}

package utilities

import "golang.org/x/crypto/bcrypt"

func CompareHashAndPassowrd(hashedPassword ,parsedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(parsedPassword))
	if err != nil {
		return err
	}else{
		return nil
	}
}
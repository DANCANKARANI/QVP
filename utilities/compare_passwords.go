package utilities

import "golang.org/x/crypto/bcrypt"

/*compares the parsed password with the password from the database
@params hashedPassword
@params ParsedPassword
*/
func CompareHashAndPassowrd(hashedPassword ,parsedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(parsedPassword))
	if err != nil {
		return err
	}else{
		return nil
	}
}
package password_generator

import (
	"log"

	"github.com/sethvargo/go-password/password"
)

func PasswordGenerator() (GeneratedPW string){
	password, error := password.Generate(16, 8, 8, false, false)
	if error != nil {
	  log.Fatal(error)
	}
	// fmt.Println("This is the random generated password:", password)
	GeneratedPW = password
	return
}
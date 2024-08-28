package auth

import "fmt"

// auth package with below function
// if we have to export this funtion to be used in another package then we have to name the functions to be started with Capital lettter.
func LoginWithCredentials(username string, password string) {
	fmt.Println("Lgging user using", username, password)
}

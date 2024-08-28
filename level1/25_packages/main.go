package main

import (
	"fmt"

	"github.com/SAHIL-Sharma21/music_player/auth"
	"github.com/SAHIL-Sharma21/music_player/user"
	"github.com/fatih/color"
)

func main() {

	auth.LoginWithCredentials("Sahil", "23445")

	userSession := auth.GetSession()
	fmt.Println("user session", userSession)

	//using struct which is defined in the user package
	newUser := user.User{
		Email: "sahil@dev.com",
		Name:  "Sahil",
	}

	// fmt.Println(newUser.Email, newUser.Name)

	//using color pkg -> 3rd party pakg in go
	//go mod tidy cmd -> used to format and correct the indirect dependency, it is the most poweful cmd

	color.Red(newUser.Email)
	color.Green(newUser.Name)
}

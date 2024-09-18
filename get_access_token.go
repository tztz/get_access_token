package main

import (
	"fmt"
	"os"

	"github.com/tztz/get_access_token/internal/environment"
	"github.com/tztz/get_access_token/pkg/accesstoken"
)

// main prints an access token (JWT) for the environment passed via command-line argument.
func main() {
	var env string
	var verboseFlag bool

	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if len(os.Args) > 2 && os.Args[2] == "-v" {
		verboseFlag = true
	}

	url, basicAuthSecret, err := environment.Data(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	token, details, err := accesstoken.New(url, basicAuthSecret)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if verboseFlag {
		fmt.Println(details)
		fmt.Printf("\nAccess token for %s: %s\n", env, token)
	} else {
		fmt.Println(token)
	}
}

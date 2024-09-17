package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// main prints an access token (JWT) for the environment passed via command-line argument.
func main() {
	var env string
	var verboseFlag bool = false

	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if len(os.Args) > 2 && os.Args[2] == "-v" {
		verboseFlag = true
	}

	token, err := accessToken(env, verboseFlag)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if verboseFlag {
		fmt.Printf("\nAccess token for %s: %s\n", env, token)
	} else {
		fmt.Println(token)
	}
}

// accessToken returns an access token (JWT) for the passed environment.
// If verboseFlag is true, then further details are printed to console.
func accessToken(env string, verboseFlag bool) (string, error) {
	authStrings, err := readFile(".env")
	if err != nil {
		return "", err
	}

	var url string
	var basicAuth string
	switch env {
	case "int":
		url = authStrings["UrlInt"]
		basicAuth = authStrings["BasicAuthInt"]
	case "pre":
		url = authStrings["UrlPre"]
		basicAuth = authStrings["BasicAuthPre"]
	case "prod":
		url = authStrings["UrlProd"]
		basicAuth = authStrings["BasicAuthProd"]
	default:
		fmt.Println("Pass environment. One of: int, pre, prod")
		fmt.Println("Pass -v as second argument to get a verbose output.")
		return "", err
	}

	payload := strings.NewReader("grant_type=client_credentials&scope=read")

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Basic "+basicAuth)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)
	token := data["access_token"].(string)

	if verboseFlag {
		fmt.Println("---- Response body ----")
		fmt.Println(string(body))
		fmt.Println("---- End of response body ----")
	}

	return token, nil
}

// readFile reads the environment file and returns the content as map.
func readFile(filename string) (map[string]string, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	rawStr := string(dat)
	rawAuthStrings := strings.Split(rawStr, "\n")

	authStrings := map[string]string{}

	for _, rawAuthString := range rawAuthStrings {
		s := strings.TrimSpace(rawAuthString)
		if s != "" {
			keyValue := strings.SplitN(s, "=", 2)
			authStrings[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		}
	}

	return authStrings, nil
}

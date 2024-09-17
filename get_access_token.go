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
	var verboseFlag bool

	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if len(os.Args) > 2 && os.Args[2] == "-v" {
		verboseFlag = true
	}

	url, basicAuthSecret, err := envData(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	token, err := accessToken(url, basicAuthSecret, verboseFlag)
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

// envData returns URL and basic auth secret for the given environment.
func envData(env string) (string, string, error) {
	authStrings, err := readFile()
	if err != nil {
		return "", "", err
	}

	var url string
	var basicAuthSecret string
	switch env {
	case "int":
		url = authStrings["UrlInt"]
		basicAuthSecret = authStrings["BasicAuthInt"]
	case "pre":
		url = authStrings["UrlPre"]
		basicAuthSecret = authStrings["BasicAuthPre"]
	case "prod":
		url = authStrings["UrlProd"]
		basicAuthSecret = authStrings["BasicAuthProd"]
	default:
		return "", "", fmt.Errorf("environment missing, pass one of int, pre, prod")
	}

	return url, basicAuthSecret, nil
}

// accessToken returns an access token (JWT) for the passed url and basic auth secret.
// If verboseFlag is true then further details are printed to console.
func accessToken(url string, basicAuthSecret string, verboseFlag bool) (string, error) {
	payload := strings.NewReader("grant_type=client_credentials&scope=read")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Basic "+basicAuthSecret)

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
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	if data == nil {
		return "", fmt.Errorf("received data is nil, maybe due to a missing VPN connection")
	}
	token := data["access_token"].(string)

	if verboseFlag {
		fmt.Println("---- Response body ----")
		fmt.Println(string(body))
		fmt.Println("---- End of response body ----")
	}

	return token, nil
}

// readFile reads the environment file and returns the content as map.
func readFile() (map[string]string, error) {
	dat, err := os.ReadFile(".env")
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

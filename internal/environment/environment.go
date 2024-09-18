package environment

import (
	"fmt"
	"os"
	"strings"
)

// Data returns URL and basic auth secret for the given environment.
func Data(env string) (string, string, error) {
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

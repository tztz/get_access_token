package environment

import (
	"fmt"
	"os"
	"strings"
)

// Data returns URL and basic auth secret for the given environment.
func Data(env string) (url string, basicAuthSecret string, err error) {
	authStrings, err := readEnvFile()
	if err != nil {
		return "", "", err
	}

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

// readEnvFile reads the environment file and returns the content as map.
func readEnvFile() (content map[string]string, err error) {
	dat, err := os.ReadFile(".env")
	if err != nil {
		return nil, err
	}
	rawStr := string(dat)
	rawAuthStrings := strings.Split(rawStr, "\n")

	content = map[string]string{}

	for _, rawAuthString := range rawAuthStrings {
		s := strings.TrimSpace(rawAuthString)
		if s != "" {
			keyValue := strings.SplitN(s, "=", 2)
			content[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		}
	}

	return content, nil
}

package environment

import (
	"fmt"
	"os"
	"strings"
)

const URL_INT = "UrlInt"
const URL_PRE = "UrlPre"
const URL_PROD = "UrlProd"
const BASIC_AUTH_INT = "BasicAuthInt"
const BASIC_AUTH_PRE = "BasicAuthPre"
const BASIC_AUTH_PROD = "BasicAuthProd"

// Data returns URL and basic auth secret for the given environment.
func Data(env string) (url string, basicAuthSecret string, err error) {
	authStrings, err := readEnv()
	if err != nil {
		return "", "", err
	}

	switch env {
	case "int":
		url = authStrings[URL_INT]
		basicAuthSecret = authStrings[BASIC_AUTH_INT]
	case "pre":
		url = authStrings[URL_PRE]
		basicAuthSecret = authStrings[BASIC_AUTH_PRE]
	case "prod":
		url = authStrings[URL_PROD]
		basicAuthSecret = authStrings[BASIC_AUTH_PROD]
	default:
		return "", "", fmt.Errorf("environment missing, pass one of int, pre, prod")
	}

	return url, basicAuthSecret, nil
}

// readEnv reads the environment file (.env), if present, and the resp. environment variables.
// Returns the content as map.
// If necessary properties are missing then an error is returned.
func readEnv() (content map[string]string, err error) {
	authStrings1 := readEnvVariables()
	authStrings2, _ := readEnvFile()

	// Merge both maps: override environment variables with values from the environment file
	for k, v := range authStrings2 {
		authStrings1[k] = v
	}

	if err := checkAllProperties(authStrings1); err != nil {
		return nil, fmt.Errorf("at least one necessary property is missing: %w", err)
	}

	return authStrings1, nil
}

// readEnvVariables reads the necessary environment variables and returns the content as map.
func readEnvVariables() map[string]string {
	return map[string]string{
		URL_INT:         os.Getenv(URL_INT),
		URL_PRE:         os.Getenv(URL_PRE),
		URL_PROD:        os.Getenv(URL_PROD),
		BASIC_AUTH_INT:  os.Getenv(BASIC_AUTH_INT),
		BASIC_AUTH_PRE:  os.Getenv(BASIC_AUTH_PRE),
		BASIC_AUTH_PROD: os.Getenv(BASIC_AUTH_PROD),
	}
}

// readEnvFile reads the environment file (.env) in the root folder of this module
// and returns the content as map.
// If the file cannot be read then an error is returned.
func readEnvFile() (content map[string]string, err error) {
	dat, err := os.ReadFile(".env")
	if err != nil {
		return nil, fmt.Errorf("environment file cannot be read: %w", err)
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

// checkAllProperties returns an error if at least one necessary property is missing.
func checkAllProperties(content map[string]string) error {
	props := []string{
		URL_INT,
		URL_PRE,
		URL_PROD,
		BASIC_AUTH_INT,
		BASIC_AUTH_PRE,
		BASIC_AUTH_PROD,
	}
	if content == nil {
		return fmt.Errorf("content is nil")
	}
	for _, key := range props {
		if content[key] == "" {
			return fmt.Errorf("property '%s' is missing", key)
		}
	}
	return nil
}

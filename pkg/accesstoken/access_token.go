package accesstoken

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// New returns an access token (JWT) together with the detailed response body
// for the passed url and basic auth secret.
func New(url string, basicAuthSecret string) (string, string, error) {
	payload := strings.NewReader("grant_type=client_credentials&scope=read")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Basic "+basicAuthSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", fmt.Errorf("received body cannot be unmarshalled: %s\nReceived body:\n%s", err, body)
	}
	token := data["access_token"].(string)

	details := "---- Response body ----\n" + string(body) + "\n---- End of response body ----"

	return token, details, nil
}

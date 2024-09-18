package accesstoken

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestShouldSuccessfullyReturnAccessTokenAndDetails(t *testing.T) {
	// given

	expectedTokenStr := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJLdGxrSE5WY3E0VnpKWTI1N3dNMjFTY2JGWmlYT1ItVUN5dk13WTZfd0FvIn0.eyJleHAiOjE3MjY2Nzk3NjUsImlhdCI6MTcyNjY3OTQ2NSwianRpIjoiNmI2NWNlOWEtYjgwZC00NmUxLWJlNGQtMzQxYmJlOWM2ZGQ0IiwiaXNzIjoiaHR0cHM6Ly9hdXRoLXRlc3QucmV3ZS5jbG91ZC9yZWFsbXMvZWNvbS1pbnQiLCJzdWIiOiIxYzUwNTZlNi1mYWMwLTQxYjYtYmU3My02MTJiZWUzMDI2M2MiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ0ZXN0LWNsaWVudC1ETy1OT1QtREVMRVRFIiwic2NvcGUiOiJyZWFkIHByb2ZpbGUgZW1haWwiLCJjbGllbnRJZCI6InRlc3QtY2xpZW50LURPLU5PVC1ERUxFVEUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiI4NC4zOC4xOTIuOSIsInByZWZlcnJlZF91c2VybmFtZSI6InNlcnZpY2UtYWNjb3VudC10ZXN0LWNsaWVudC1kby1ub3QtZGVsZXRlIiwiY2xpZW50QWRkcmVzcyI6Ijg0LjM4LjE5Mi45In0.RNzHRws6hEF9kDOlMljkp9NDdh6LAUv7fRGEC29M621HnE84Fy_SzaX0FGJPL32-ver3wWeRCV8OI0nFzDjjLBUyRaj8Jgg7hn01WDQo9Tgdhk3u72VsOa74tqLbtPaRkWjWv7TcZfwuIRHcDGxbsRHUMos_zlo3cCRj3QL2HDfXXv2GXMm7q06t7HEWQ-gNNgLBR0NW1RxdsYUFk6WuTnipgA-Lu3RswqpUS6Ayg4hhz1WnncOUV6sM8Lgw-o-JRfhWwL5a8sCBno_75hg4-IkPOvn0CN46mPlZ7-CsJPMnv10nzlA7WZruUy1eLxOm6uZy8Sk8CzTkpWVFGvyB2g"
	expectedResponseStr := fmt.Sprintf(`{"access_token":"%s","expires_in":300,"refresh_expires_in":0,"token_type":"Bearer","not-before-policy":0,"scope":"read profile email"}`, expectedTokenStr)
	expectedDetails := fmt.Sprintf("---- Response body ----\n%s\n---- End of response body ----", expectedResponseStr)

	// and

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://example.com/get-token",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("content-type") != "application/x-www-form-urlencoded" {
				t.Errorf("expected 'content-type: application/x-www-form-urlencoded' header, got: %s", req.Header.Get("Accept"))
			}
			response := httpmock.NewBytesResponse(200, []byte(expectedResponseStr))
			return response, nil
		},
	)

	// and

	url := "https://example.com/get-token"
	basicAuthSecret := "secret-1234"

	// when

	token, details, err := New(url, basicAuthSecret)

	// then

	if token != expectedTokenStr {
		t.Errorf("expected access token with value\n%s\ngot:\n%s", expectedTokenStr, token)
	}

	if details != expectedDetails {
		t.Errorf("expected details with value\n%s\ngot:\n%s", expectedDetails, details)
	}

	if err != nil {
		t.Errorf("expected err to be nil, but was nil")
	}
}

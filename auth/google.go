package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGoogleUserInfo(t string) (UserInfo, error) {
	url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", t)
	var ui UserInfo
	err := getJSON(url, &ui)
	if err != nil {
		return ui, err
	}
	return ui, nil
}

// UserInfo represents a response from
// google to a token verification request.
type UserInfo struct {
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
}

// getJSON requests google's api url and
// attempts to convert the response into
// a UserInfo struct.
func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

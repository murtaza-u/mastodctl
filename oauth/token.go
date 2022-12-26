package oauth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type token struct {
	Access string `json:"access_token"`
}

func (o O) Token(code string) (string, error) {
	uri := o.Server + TokenEnd

	params := make(url.Values, 5)
	params.Set("client_id", o.ClientID)
	params.Set("client_secret", o.ClientSecret)
	params.Set("redirect_uri", RedirectURI)
	params.Set("grant_type", "authorization_code")
	params.Set("scope", o.scope)
	params.Set("code", code)

	resp, err := http.PostForm(uri, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	tkn := new(token)
	err = json.Unmarshal(body, tkn)
	if err != nil {
		return "", err
	}

	return tkn.Access, nil
}

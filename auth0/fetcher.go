package auth0

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TokenFetcher implementers can fetch an auth0 token.
type TokenFetcher interface {
	Token(audience string) (string, error)
}

// NewTokenFetcher creates a tokenFetcher that can get an access
// token for a client_credentials grant from Auth0.
//
// The 3DSIM prod token endpoint is: https://3dsim.auth0.com/oauth/token
// The 3DSIM token endpoint is: https://3dsim-qa.auth0.com/oauth/token
func NewTokenFetcher(httpClient *http.Client, tokenURL, clientID, clientSecret string) TokenFetcher {
	return &tokenFetcher{
		httpClient:   httpClient,
		clientID:     clientID,
		clientSecret: clientSecret,
		grantType:    "client_credentials",
		tokenURL:     tokenURL,
	}
}

type tokenFetcher struct {
	clientID     string
	clientSecret string
	grantType    string
	tokenURL     string
	httpClient   *http.Client
}

func (a *tokenFetcher) Token(audience string) (string, error) {
	request := &request{
		Audience:     audience,
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
		GrantType:    a.grantType,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	resp, err := a.httpClient.Post(a.tokenURL, "application/json", bytes.NewBuffer(requestBytes))
	if err != nil {
		return "", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	var tokenResponse response
	err = json.Unmarshal(respBytes, &tokenResponse)
	if err != nil {
		return "", err
	}
	return tokenResponse.AccessToken, nil
}

type request struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (r *response) String() string {
	return r.TokenType + " " + r.AccessToken
}

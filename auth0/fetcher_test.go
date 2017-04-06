package auth0

import (
	"github.com/stretchr/testify/assert"

	"net/http"
	"testing"
)

func TestNewClientCredentialsFlowExpectsObjectInitializedSuccessfully(t *testing.T) {
	// arrange
	clientID := "PV2AvGcMjOFErV6QpaqKnfrUdt8yPuHI"
	clientSecret := "9oXvXvWHfQaaAiWr-wBfS5Vtyp3aGyMuwIwqYs2NuRtmV7-1XEXXNJ1ZA97jLo6J"
	tokenURL := "https://3dsim.auth0.com/oauth/token"
	grantType := "client_credentials"

	// act
	f := NewTokenFetcher(nil, tokenURL, clientID, clientSecret)

	// assert
	a := f.(*tokenFetcher)
	assert.Equal(t, clientID, a.clientID, "expected client id to match")
	assert.Equal(t, clientSecret, a.clientSecret, "expected client secret to match")
	assert.Equal(t, tokenURL, a.tokenURL, "expected tokenURL to match")
	assert.Equal(t, grantType, a.grantType, "expected grant type to match")
}

// Integration test
func _TestAccessTokenFromGlobalCredentials(t *testing.T) {
	// arrange
	clientID := "zO2mKgrhEA6kcI23E0lRHutHBd1AX8ht"
	clientSecret := "W20FgCL0FZ_ZMFmpdrh13Y49WuflWrJswPqJjDtaXtcaUAYD2x0ETsiPQ1xh8xez"
	tokenURL := "https://3dsim-qa.auth0.com/oauth/token"

	// act
	f := NewTokenFetcher(http.DefaultClient, tokenURL, clientID, clientSecret)
	accessToken, err := f.Token("")

	// assert
	assert.Nil(t, err, "Expected no errors")
	assert.NotEmpty(t, accessToken, "Expected a non-empty access token")
}

package auth0_client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	structObject "backend/internal/4_domain/struct_object"
	valueObject "backend/internal/4_domain/value_object"
)

// ...
// FetchAccessToken ...
func (receiver *Auth0Client) FetchAccessToken(
	ctx context.Context,
	credential structObject.Credential,
) (
	accessToken valueObject.AccessToken,
) {
	payload := map[string]string{
		"client_id":     credential.ClientID.Content.GetValue(),
		"client_secret": credential.ClientSecret.Content.GetValue(),
		"audience":      "https://auth0my-yayuji.com",
		"grant_type":    "client_credentials",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		accessToken.SetError(err)
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"https://dev-fe00zeb23uke8zls.us.auth0.com/oauth/token",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		accessToken.SetError(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		accessToken.SetError(err)
		return
	}
	defer resp.Body.Close()

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		accessToken.SetError(err)
		return
	}

	accessToken = valueObject.NewAccessToken(
		&tokenResponse.AccessToken,
	)

	return
}

package auth0

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"backend/internal/4_domain/struct_object"
	"backend/internal/4_domain/value_object"
)

// ...
// FetchAccessToken ...
func (receiver *Auth0) FetchAccessToken(
	ctx context.Context,
	credential struct_object.Credential,
) (
	accessToken value_object.AccessToken,
	err error,
) {
	payload := map[string]string{
		"client_id":     credential.ClientID.Content.GetValue(),
		"client_secret": credential.ClientSecret.Content.GetValue(),
		"audience":      "https://auth0my-yayuji.com",
		"grant_type":    "client_credentials",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return accessToken, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"https://dev-fe00zeb23uke8zls.us.auth0.com/oauth/token",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return accessToken, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return accessToken, err
	}
	defer resp.Body.Close()

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return accessToken, err
	}

	accessToken, err = value_object.NewAccessToken(
		&tokenResponse.AccessToken,
	)
	if err != nil {
		return accessToken, err
	}

	return
}

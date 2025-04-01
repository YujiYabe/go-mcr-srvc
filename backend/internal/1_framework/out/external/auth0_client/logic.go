package auth0_client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

// ...
// FetchAccessToken ...
func (receiver *Auth0Client) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken typeObject.AccessToken,
) {
	return accessToken

	payload := map[string]string{
		"client_id":     credential.ClientID.GetValue(),
		"client_secret": credential.ClientSecret.GetValue(),
		"audience":      "https://auth0my-yayuji.com",
		"grant_type":    "client_credentials",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		accessToken.SetError(ctx, err)
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"https://dev-fe00zeb23uke8zls.us.auth0.com/oauth/token",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		accessToken.SetError(ctx, err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		accessToken.SetError(ctx, err)
		return
	}
	defer resp.Body.Close()

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		accessToken.SetError(ctx, err)
		return
	}

	accessToken = typeObject.NewAccessToken(
		ctx,
		&tokenResponse.AccessToken,
	)

	return
}

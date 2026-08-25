package auth0_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	err error,
) {
	payload := map[string]string{
		"client_id":     credential.ClientID().GetValue(),
		"client_secret": credential.ClientSecret().GetValue(),
		"audience":      receiver.audience,
		"grant_type":    receiver.grantType,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		receiver.tokenURL,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		err = fmt.Errorf("auth0 token request failed: status %d", resp.StatusCode)
		return
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return
	}

	accessToken, err = typeObject.NewAccessToken(
		&tokenResponse.AccessToken,
	)

	return
}

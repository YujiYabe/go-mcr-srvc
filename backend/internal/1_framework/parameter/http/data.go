package http_parameter

type V1User struct {
	ID    *int    `json:"id,omitempty" query:"id"`
	Name  *string `json:"name,omitempty" query:"name"`
	Email *string `json:"email,omitempty" query:"email"`
}

type V1Credential struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
}

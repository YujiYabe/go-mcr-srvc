package http_parameter

type V1Person struct {
	ID          *int    `json:"id,omitempty" query:"id"`
	Name        *string `json:"name,omitempty" query:"name"`
	MailAddress *string `json:"mailAddress,omitempty" query:"mailAddress"`
}

type V1Credential struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
}

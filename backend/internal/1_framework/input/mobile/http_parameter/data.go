package http_parameter

type V1PersonParameter struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	MailAddress *string `json:"mail_address,omitempty"`
}

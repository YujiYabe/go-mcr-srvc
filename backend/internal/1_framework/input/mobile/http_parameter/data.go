package http_parameter

type V1PersonParameter struct {
	ID          *int    `json:"id,omitempty" query:"id"`
	Name        *string `json:"name,omitempty" query:"name"`
	MailAddress *string `json:"mail_address,omitempty" query:"mail_address"`
}

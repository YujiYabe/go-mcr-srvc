package http_parameter

type V1PersonParameter struct {
	ID          *int    `json:"id,omitempty" query:"id"`
	Name        *string `json:"name,omitempty" query:"name"`
	MailAddress *string `json:"mailAddress,omitempty" query:"mailAddress"`
}

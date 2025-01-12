package http_parameter

type V1CommonParameter struct {
	V1Error          *V1Error         `json:"v1_error,omitempty"`
	V1RequestContext V1RequestContext `json:"v1_request_context"`
}

type V1Error struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Detail  string    `json:"detail"`
	Field   string    `json:"field"`
	Errors  []V1Error `json:"errors"`
}

type V1RequestContext struct {
	TraceID          string   `json:"trace_id"`
	RequestStartTime int64    `json:"request_start_time"`
	ClientIP         string   `json:"client_ip"`
	UserAgent        string   `json:"user_agent"`
	UserID           string   `json:"user_id"`
	Permissions      []string `json:"permissions"`
	AccessToken      string   `json:"access_token"`
	TenantID         string   `json:"tenant_id"`
	Locale           string   `json:"locale"`
	Timezone         string   `json:"timezone"`
}

type GetPersonListByConditionRequest struct {
	V1CommonParameter *V1CommonParameter `protobuf:"bytes,1,opt,name=v1_common_parameter,json=v1CommonParameter" json:"v1_common_parameter,omitempty"`
	V1PersonParameter *V1PersonParameter `protobuf:"bytes,2,req,name=v1_person_parameter,json=v1PersonParameter" json:"v1_person_parameter"`
}

type GetPersonListByConditionResponse struct {
	V1CommonParameter      *V1CommonParameter      `protobuf:"bytes,1,req,name=v1_common_parameter,json=v1CommonParameter" json:"v1_common_parameter"`
	V1PersonParameterArray *V1PersonParameterArray `protobuf:"bytes,2,req,name=v1_person_parameter_array,json=v1PersonParameterArray" json:"v1_person_parameter_array"`
}

type V1PersonParameter struct {
	ID          *uint32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Name        *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	MailAddress *string `protobuf:"bytes,4,opt,name=mail_address,json=mailAddress" json:"mail_address,omitempty"`
}

type V1PersonParameterArray struct {
	Persons []*V1PersonParameter `protobuf:"bytes,2,rep,name=persons" json:"persons"`
}

type V1Person struct {
	ID          *int    `json:"id,omitempty" query:"id"`
	Name        *string `json:"name,omitempty" query:"name"`
	MailAddress *string `json:"mailAddress,omitempty" query:"mailAddress"`
}

type V1Credential struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
}

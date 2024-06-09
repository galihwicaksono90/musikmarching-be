package response

type Meta struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func Response(code uint, message string, data interface{}) *APIResponse {
	response := APIResponse{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
	return &response
}

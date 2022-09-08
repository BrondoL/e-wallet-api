package utils

type SResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type EResponse struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func SuccessResponse(message string, code int, data interface{}) SResponse {
	return SResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "success",
		},
		Data: data,
	}
}

func ErrorResponse(message string, code int, err interface{}) EResponse {
	return EResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "error",
		},
		Error: err,
	}
}

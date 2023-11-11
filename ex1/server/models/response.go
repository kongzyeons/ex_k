package models

type messageResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(result interface{}, code int, massage string) messageResponse {
	return messageResponse{
		Status:  true,
		Code:    code,
		Message: massage,
		Data:    result,
	}
}

func Err_response(err error, code int) messageResponse {
	var data interface{}
	return messageResponse{
		Status:  false,
		Code:    code,
		Message: err.Error(),
		Data:    data}
}

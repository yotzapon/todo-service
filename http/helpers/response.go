package helpers

import "net/http"

// Response standard response struct
type Response struct {
	ResponseCode        int16       `json:"responseCode" example:"200"`
	ResponseDescription string      `json:"responseDescription" example:"OK"`
	Data                interface{} `json:"data"`
}

// ResponseError error response struct
type ResponseError struct {
	ResponseCode        int16  `json:"responseCode" example:"4001"`
	ResponseDescription string `json:"responseDescription" example:"INVALID_BODY_OR_PARAM_REQUEST"`
}

// SuccessJSON compose response success JSON
func SuccessJSON(status int16, data interface{}) Response {
	return Response{
		ResponseCode:        status,
		ResponseDescription: http.StatusText(int(status)),
		Data:                data,
	}
}

// ErrorJSON compose response error JSON
func ErrorJSON(errorCode int, errText string) ResponseError {
	return ResponseError{
		ResponseCode:        int16(errorCode),
		ResponseDescription: errText,
	}
}

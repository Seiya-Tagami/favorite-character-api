package errors

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ToResponse(code int, message string) Response {
	return Response{
		code,
		message,
	}
}

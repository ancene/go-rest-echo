package serialization

type Error struct {
	Code  int         `json:"code" xml:"code"`
	Stack interface{} `json:"stack" xml:"stack"`
}

type Response struct {
	Success bool        `json:"success" xml:"success"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
	Error   *Error      `json:"error" xml:"error"`
}

func OK(msg string, data interface{}) Response {
	return Response{
		Success: true,
		Message: msg,
		Data:    data,
		Error:   nil,
	}
}

func BadRequest() Response {
	return Response{
		Success: false,
		Message: "Bad Request",
		Data:    nil,
		Error: &Error{
			Code:  400,
			Stack: []string{},
		},
	}
}

func Unauthorized() Response {
	return Response{
		Success: false,
		Message: "Unauthorized",
		Data:    nil,
		Error: &Error{
			Code:  401,
			Stack: []string{},
		},
	}
}

func Forbidden() Response {
	return Response{
		Success: false,
		Message: "Forbidden",
		Data:    nil,
		Error: &Error{
			Code:  403,
			Stack: []string{},
		},
	}
}

func NotFound() Response {
	return Response{
		Success: false,
		Message: "Not Found",
		Data:    nil,
		Error: &Error{
			Code:  404,
			Stack: []string{},
		},
	}
}

func MethodNotAllowed() Response {
	return Response{
		Success: false,
		Message: "Method Not Allowed",
		Data:    nil,
		Error: &Error{
			Code:  405,
			Stack: []string{},
		},
	}
}

func RequestTimeout() Response {
	return Response{
		Success: false,
		Message: "Request Timeout",
		Data:    nil,
		Error: &Error{
			Code:  408,
			Stack: []string{},
		},
	}
}

func PayloadTooLarge() Response {
	return Response{
		Success: false,
		Message: "Payload Too Large",
		Data:    nil,
		Error: &Error{
			Code:  413,
			Stack: []string{},
		},
	}
}

func UnprocessableEntity() Response {
	return Response{
		Success: false,
		Message: "Unprocessable Entity",
		Data:    nil,
		Error: &Error{
			Code:  422,
			Stack: []string{},
		},
	}
}

func TooManyRequests() Response {
	return Response{
		Success: false,
		Message: "Too Many Requests",
		Data:    nil,
		Error: &Error{
			Code:  429,
			Stack: []string{},
		},
	}
}

func InternalServerError() Response {
	return Response{
		Success: false,
		Message: "Internal Server Error",
		Data:    nil,
		Error: &Error{
			Code:  500,
			Stack: []string{},
		},
	}
}

func ServiceUnavailable() Response {
	return Response{
		Success: false,
		Message: "Service Unavailable",
		Data:    nil,
		Error: &Error{
			Code:  503,
			Stack: []string{},
		},
	}
}

package helper

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseSuccessWithoutData struct {
	Message string `json:"message"`
}

type ResponseFailed struct {
	Message string `json:"message"`
}

func APIResponseSuccess(message string, data interface{}) ResponseSuccess {
	jsonResponse := ResponseSuccess{
		Message: message,
		Data:    data,
	}
	return jsonResponse
}

func APIResponseSuccessWithouData(message string) ResponseSuccessWithoutData {
	jsonResponse := ResponseSuccessWithoutData{
		Message: message,
	}
	return jsonResponse
}

func APIResponseFailed(message string) ResponseFailed {
	jsonResponse := ResponseFailed{
		Message: message,
	}
	return jsonResponse
}

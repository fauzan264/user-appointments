package helper

type Response struct {
	Status		bool 		`json:"status"`
	Message 	string 		`json:"message"`
	Data		interface{}	`json:"data"`
}

func APIResponse(status bool, message string, data interface{}) Response {
	jsonResponse := Response{
		Status : status,
		Message: message,
		Data: data,
	}

	return jsonResponse
}
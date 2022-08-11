package libs

import structs "github.com/Javad-Smn/simple-modular-go/src/structs"

func Response(status int, message string, data interface{}) structs.Response {
	response := structs.Response{
		Message: message,
		Status:  status,
		Data:    data,
	}
	return response
}

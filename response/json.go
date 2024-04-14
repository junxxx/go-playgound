package response

import "encoding/json"

type Response struct {
	Status int
	Data   interface{}
	Msg    string
}

func Json(status int, data interface{}, msg string) ([]byte, error) {
	res := Response{
		Status: status,
		Data:   data,
		Msg:    msg,
	}
	return json.Marshal(res)
}

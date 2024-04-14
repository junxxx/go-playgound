package response

import "encoding/json"

type Response struct {
	Status int
	Data   interface{}
	Msg    string
}

func Json(data interface{}) ([]byte, error) {
	res := Response{Data: data}
	return json.Marshal(res)
}

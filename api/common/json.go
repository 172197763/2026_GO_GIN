package common

import "encoding/json"

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func JsonEncode(data any) string {
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}
func JsonDecode(str string, data any) error {
	return json.Unmarshal([]byte(str), data)
}

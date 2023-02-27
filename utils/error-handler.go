package utils

import "encoding/json"

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandlerToJson(ErrMsg map[string]interface{}) []byte {
	errStr, err := json.Marshal(ErrMsg)
	ErrorHandler(err)
	return errStr
}

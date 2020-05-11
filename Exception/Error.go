package Exception

import ()

type ErrorHandler struct {
	Errorcode string "json:errorcode"
	ErrorDesc string "json:errordesc"
}

func SetErrorDetails(Code string, Desc string) *ErrorHandler {
	err := &ErrorHandler{Code, Desc}
	return err
}

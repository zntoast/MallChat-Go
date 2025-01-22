package ercode

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ResponseErrHandler(err error) (int, any) {
	switch err.(type) {
	case *errorEntry:
		code := GetCode(err)
		return code, err.Error()
	default:
		return 400, err.Error()
	}
}

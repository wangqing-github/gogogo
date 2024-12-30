package msg

type Response struct {
	status int32
	msg    string
	data   any
}

func Success(data any) Response {
	resp := Response{
		status: 200,
		data:   data,
		msg:    "ok",
	}
	return resp
}

func Fail(msg string) Response {
	if len(msg) == 0 {
		msg = "error"
	}
	resp := Response{
		status: 100,
		msg:    msg,
	}
	return resp
}

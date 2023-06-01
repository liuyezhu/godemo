package response

const (
	ErrnoSuccess = 0
	ErrnoError   = 1
	ErrnoUnKnown = 2
	ErrnoTest    = 3
)

var ErrNoToMsgMap = map[int32]string{
	ErrnoSuccess: "success",
	ErrnoError:   "failed",
	ErrnoUnKnown: "unknown",
	ErrnoTest:    "test",
}

func GetErrMsg(errNo int32) string {

	if errMsg, ok := ErrNoToMsgMap[errNo]; ok {
		return errMsg
	}
	return "unknown error"
}

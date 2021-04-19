package resp

type ErrorCode int32

const (
	RESP_CODE_SUCCESS         ErrorCode = 0
	RESP_CODE_PARAMS_ERROR    ErrorCode = 4001
	RESP_CODE_SEVER_EXCEPTION ErrorCode = 5000
	RESP_CODE_DB_ERROR        ErrorCode = 5001
	RESP_CODE_REDIS_ERROR     ErrorCode = 5002
	RESP_CODE_RPC_ERROR       ErrorCode = 5003
)

type IErrorCode interface {
	Prompts() string
	Message() string
	Status() int32
}

func (p ErrorCode) Prompts() string {
	switch p {
	case RESP_CODE_SUCCESS:
		return ""
	case RESP_CODE_PARAMS_ERROR:
		return "请求参数错误"
	case RESP_CODE_SEVER_EXCEPTION, RESP_CODE_DB_ERROR,
		RESP_CODE_REDIS_ERROR, RESP_CODE_RPC_ERROR:
		return "服务器内部错误，请稍后重试"
	}
	return "unkown error"
}

func (p ErrorCode) Message() string {
	switch p {
	case RESP_CODE_SUCCESS:
		return "success"
	case RESP_CODE_PARAMS_ERROR:
		return "params error"
	case RESP_CODE_SEVER_EXCEPTION, RESP_CODE_DB_ERROR,
		RESP_CODE_REDIS_ERROR, RESP_CODE_RPC_ERROR:
		return "server exception"
	}
	return "unkown error"
}

func (p ErrorCode) Status() int32 {
	return int32(p)
}

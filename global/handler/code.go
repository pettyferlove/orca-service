package handler

type ResultCode int

const (
	Success                      ResultCode = 0
	Failure                      ResultCode = 1
	InvalidArgument              ResultCode = 1001
	UserAuthenticateError        ResultCode = 2001
	InterfaceAddressInvalid      ResultCode = 3001
	InterfaceRequestTimeout      ResultCode = 3002
	InterfaceExceedLoad          ResultCode = 3003
	PermissionNoAccess           ResultCode = 4001
	PermissionAuthenticateFailed ResultCode = 4002
	PermissionTokenInvalid       ResultCode = 4003
	PermissionTokenExpired       ResultCode = 4004
	PermissionSignatureError     ResultCode = 4005
	DataNotFound                 ResultCode = 5001
	DataIsWrong                  ResultCode = 5002
	DataAlreadyExist             ResultCode = 5003
	BusinessErr                  ResultCode = 6001
	ServiceException             ResultCode = 9999
)

var ResultCodeMessage = map[ResultCode]string{
	Success:                      "成功",
	Failure:                      "失败",
	InvalidArgument:              "参数无效",
	UserAuthenticateError:        "用户认证失败",
	InterfaceAddressInvalid:      "接口地址无效",
	InterfaceRequestTimeout:      "接口请求超时",
	InterfaceExceedLoad:          "接口负载过高",
	PermissionNoAccess:           "无访问权限",
	PermissionAuthenticateFailed: "认证失败",
	PermissionTokenInvalid:       "Token无效",
	PermissionTokenExpired:       "Token过期",
	PermissionSignatureError:     "签名错误",
	DataNotFound:                 "数据未找到",
	DataIsWrong:                  "数据有误",
	DataAlreadyExist:             "数据已存在",
	BusinessErr:                  "业务错误",
	ServiceException:             "服务异常",
}

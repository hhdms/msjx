package models

// Response 通用响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(data interface{}) Response {
	return Response{
		Code: 1,
		Msg:  "success",
		Data: data,
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(msg string) Response {
	return Response{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
}

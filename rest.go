package structs

// 定义rest 请求/返回数据通用结构

// Response 返回
type Response struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

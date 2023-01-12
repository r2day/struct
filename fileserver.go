package structs

// FileServerUploadRequest 文件服务器上传请求
type FileServerUploadRequest struct {
}

// FileServerUploadSuccessResponse 文件服务器上传成功返回
type FileServerUploadSuccessResponse struct {
	Url string `json:"url"`
}
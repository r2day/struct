package structs

// 前端上传单个图片编码
type UploadBase64ImageRequest struct {
	// 图片转换为base64后的字符串
	Src string `json:"src"`
	// 图片标题
	Title string `json:"title"`
}
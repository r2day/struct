package structs

// ReactImageRawFile 定义react-admin/image-input 上传图片的数据类型
type ReactImageRawFile struct {
	Path string `json:"path"`
}

// ReactImageInputPicture 图片信息
type ReactImageInputPicture struct {
	RawFile  ReactImageRawFile `json:"rawFile"`
	Src string `json:"src"`
	Title string `json:"title"`
}

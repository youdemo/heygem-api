package filein

type UploadInp struct {
	File     string `json:"file"   dc:"文件"`
	FileName string `json:"fileName" dc:"文件名"`
	FileType string `json:"fileType" dc:"文件类型"`
	Path     string `json:"path" dc:"文件路径"`
}

type DownloadInp struct {
	FileType string `json:"fileType" dc:"文件类型"`
	Path     string `json:"path" dc:"文件路径"`
}

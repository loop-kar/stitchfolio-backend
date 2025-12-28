package responseModel

type FileResponse struct {
	FileUrl  string `json:"fileUrl,omitempty"`
	FileName string `json:"fileName,omitempty"`
}

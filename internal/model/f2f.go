package model

type FaceResponse struct {
	Code    int       `json:"code" dc:"状态码"`
	Data    *FaceData `json:"data"`
	Msg     string    `json:"msg"`
	Success bool      `json:"success"`
}

type FaceData struct {
	Code     string `json:"code" dc:"任务编码"`
	Msg      string `json:"msg"`
	Progress int    `json:"progress"`
	Result   string `json:"result"`
	Status   int    `json:"status"`
}

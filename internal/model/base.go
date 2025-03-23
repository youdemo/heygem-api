package model

type PageReq struct {
	Page     int `json:"page" d:"1"`
	PageSize int `json:"pageSize" d:"10"`
}

type PageRes struct {
	Total int `json:"total" example:"0" dc:"数据总行数"`
}

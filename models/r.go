package models

type R struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Total   int32       `json:"total"`
}

package models

type ApiSearch struct {
	Id         int64  `json:"id"`
	QueryParam string `json:"query_param"`
	ApiStatus  int16   `json:"api_status"`
	Response   string `json:"response"`
}

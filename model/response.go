package models

type ResponseView struct {
	ResponseCode int         `json:"ResponseCode"`
	ResponseDesc string      `json:"ResponseDesc"`
	ResponseTime string      `json:"ResponseTime"`
	Message      interface{} `json:"Message"`
	Result       interface{} `json:"Result"`
}

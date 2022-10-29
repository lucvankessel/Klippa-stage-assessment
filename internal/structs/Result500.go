package structs

type Result500 struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Request_id string `json:"request_id"`
	Result string `json:"result"`
}
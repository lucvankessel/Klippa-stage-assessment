package structs

type Result400 struct {
	Code int `json:"code"`
	Errors []struct {
		Code int `json:"code"`
		Message string `json:"message"`
		Request_id string `json:"request_id"`
		Result string `json:"result"`
	} `json:"errors"`
	Request_id string `json:"request_id"`
	Result string `json:"result"`
}
package structs


type Result200 struct {

	// Because this project can facilitate multiple templates i didnt make a detailed struct for them all, i store the raw 'data' object we get from a response in the Data field and need to loop over it in any
	Data map[string]interface{} `json:"data"`
	Request_id string `json:"request_id"`
	Result string `json:"result"`

}

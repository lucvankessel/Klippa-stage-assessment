package structs


type Result200 struct {

	// Because this project can facilitate multiple templates i didnt make a detailed struct, i store the common fields and store the data object into this Data field. from there i made a function that can loop over any of the given results and display it in a custom way.
	Data map[string]interface{} `json:"data"`
	Request_id string `json:"request_id"`
	Result string `json:"result"`

}

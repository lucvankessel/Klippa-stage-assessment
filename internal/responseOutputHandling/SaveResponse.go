package responseoutputhandling

import (
	"assessment/internal/structs"
	"encoding/json"
	"fmt"
	"os"
)

// saves the response body of a request response to a given filename.
func SaveResponse(bodyData []byte, requestconfig structs.RequestConfig) {
	filename := requestconfig.SavefileName + ".json"
	
	// from the body byte array we create the jsonmap.
	var jsonmap map[string]*json.RawMessage
	if err := json.Unmarshal(bodyData, &jsonmap); err != nil {
		fmt.Println("JsonUnmarshal error: ", err)
		os.Exit(0)
	}

	var jsonIndent []byte
	// indenting to make it also a bit more human readable.
	jsonIndent, err := json.MarshalIndent(jsonmap, "", " "); 
	if err != nil {
		fmt.Println("JsonIndent Error: ", err)
		os.Exit(0)
	}

	_ = os.WriteFile(filename, jsonIndent, 0644)
}
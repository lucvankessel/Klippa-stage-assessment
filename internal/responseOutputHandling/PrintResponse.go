package responseoutputhandling

import (
	"assessment/internal/structs"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Because in some cases i dont know how deep the responses can go i needed to make a function that figures that out for me and prints everything that way.
func recursePrint(input map[string]interface{}, depth int, fulloutput bool) {

	for key, value := range input {
		// if there is no value i dont want to print it, for a cli output i see it as useless output.
		if !fulloutput {
			if value == nil || value == "" || value == float64(0) || key == "raw_text" {
				continue
			}
		}

		tabsDivider := strings.Repeat("\t", depth-1)
		tabsValues := strings.Repeat("\t", depth)
		if rec, ok := value.(map[string]interface{}); ok {
			divider := strings.Repeat("=", depth)
			fmt.Printf("%[3]v%[1]s %[2]s %[1]s \n", divider, key, tabsDivider)
			recursePrint(rec, depth+1, fulloutput)
		} else {
			fmt.Printf("%[3]v %[1]s: %[2]v \n", key, value, tabsValues)
		}

	}

}

// This function will pretty print the result in the console.
func PrintResponse(bodyData []byte, statusCode int, fulloutput bool) {

	var jsonmap map[string]interface{}
	if err := json.Unmarshal(bodyData, &jsonmap); err != nil {
		fmt.Println("print JsonUnmarshal error: ", err)
		os.Exit(0)
	}

	fmt.Println("=== PARSE RESULTS ===")
	fmt.Println("Status: ",	jsonmap["result"])

	// Because i want to give the output some styling i decided to manually print the results instead of using marshalindent.
	if statusCode == 200 {
		var result200 structs.Result200
		json.Unmarshal(bodyData, &result200)

		recursePrint(result200.Data, 1, fulloutput)
		fmt.Println("Request ID: ", result200.Request_id)

	} else if statusCode == 400 {
		var result400 structs.Result400
		json.Unmarshal(bodyData, &result400)

		fmt.Println("Error code: ", result400.Code)

		fmt.Println("=Errors=")
		for i, error := range result400.Errors {
			fmt.Println("Error ", i)
			fmt.Println("Code: ", error.Code)
			fmt.Println("Message: ", error.Message)
			fmt.Println("Request id: ", error.Request_id)
			fmt.Println("Result: ", error.Result)
		}

		fmt.Println("Request id: ", result400.Request_id)


	} else if statusCode == 500 {
		var result500 structs.Result500
		json.Unmarshal(bodyData, &result500)

		fmt.Println("Error code: ", result500.Code)
		fmt.Println("Error message: ", result500.Message)
		fmt.Println("Request id: ", result500.Request_id)

	}

}
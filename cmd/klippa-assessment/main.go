package main

import (
	"assessment/internal/klippa-api"
	"assessment/internal/structs"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Initialize the flags, this is done through the flag package provided by GO
	ApiKey := flag.String("api", "[REQUIRED]", "a api key")
	Template := flag.String("template", "financial_full", "which default provided template to use")
	ExtractionType := flag.String("textextraction", "fast", "what thype of text extraction you want to use, needs to be fast or full")
	FilePath := flag.String("file", "[REQUIRED]", "a document or image file path")
	SaveFile := flag.String("save", "filename", "name of how you want to save the result")

	flag.Parse()

	// check if the file mentioned in the filepath exists
	if _, err := os.Stat(string(*FilePath)); errors.Is(err, os.ErrNotExist) {
		fmt.Println("This file does not exist: ", *FilePath)
		os.Exit(0)
	}

	// create construct from the input flags.
	requestconfig := new(structs.RequestConfig)
	requestconfig.ApiKey = string(*ApiKey)
	requestconfig.Template = string(*Template)
	requestconfig.ExtractionType = string(*ExtractionType)
	requestconfig.FilePath = string(*FilePath)
	requestconfig.SavefileName = string(*SaveFile)

	// give the request config to the ParseDocument function, this will execute the api call.
	response := klippaApi.ParseDocument(requestconfig)


	// TODO: voor wat voor reden kan ik niet 2 keer op dezelfde response een readall doen, de laatste van de 2 zal een fout krijgen bij het lezen (namelijk dat deze leeg is.)
	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Save ReadAll error: ", err)
		os.Exit(0)
	}

	// check if the given savefile name isnt the default one. if it is different we save it to a file.
	if requestconfig.SavefileName != "filename" {
		SaveResponse(bodyData, *requestconfig)
	}
	
	PrintResponse(bodyData, response.StatusCode)

}


// saves the response body of a request response to a given filename.
func SaveResponse(bodyData []byte, requestconfig structs.RequestConfig) {
	filename := requestconfig.SavefileName + ".json"
	
	var jsonmap map[string]*json.RawMessage
	if err := json.Unmarshal(bodyData, &jsonmap); err != nil {
		fmt.Println("JsonUnmarshal error: ", err)
		os.Exit(0)
	}

	var jsonIndent []byte
	jsonIndent, err := json.MarshalIndent(jsonmap, "", " "); 
	if err != nil {
		fmt.Println("JsonIndent Error: ", err)
		os.Exit(0)
	}

	_ = os.WriteFile(filename, jsonIndent, 0644)
}


// This function will pretty print the result in the console.
func PrintResponse(bodyData []byte, statusCode int) {

	var jsonmap map[string]interface{}
	if err := json.Unmarshal(bodyData, &jsonmap); err != nil {
		fmt.Println("print JsonUnmarshal error: ", err)
		os.Exit(0)
	}

	fmt.Println("=== PARSE RESULTS ===")
	fmt.Println("Status: ",	jsonmap["result"])

	// Because i want to give the output some styling i decided to manually print the results instead of using marshalindent.
	if statusCode == 200 {

		fmt.Println("= Raw Data =")
		fmt.Println(jsonmap["data"])

	} else if statusCode == 400 {

	} else if statusCode == 500 {
		fmt.Println("Error code: ", jsonmap["code"])
		fmt.Println("Error message: ", jsonmap["message"])
		fmt.Println("Request id: ", jsonmap["request_id"])
	}

}
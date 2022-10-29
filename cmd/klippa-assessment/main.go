package main

import (
	"assessment/internal/klippa-api"
	"assessment/internal/structs"
	// "encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	// "net/http/httputil"
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
	response, err := klippaApi.ParseDocument(requestconfig)

	// check if the given savefile name isnt the default one. if it is different we save it to a file.
	if requestconfig.SavefileName != "filename" {
		SaveResponse(response, *requestconfig)
	}

	// print the response to the console.
	fmt.Println(err)

}

// saves the response body of a request response to a given filename.
func SaveResponse(response *http.Response, requestconfig structs.RequestConfig) {
	filename := requestconfig.SavefileName + ".json"
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	_ = os.WriteFile(filename, body, 0644)
}

// This function will pretty print the result in the console.
func PrintResponse(response *http.Response) {

}
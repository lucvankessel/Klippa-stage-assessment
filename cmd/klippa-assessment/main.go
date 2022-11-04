package main

import (
	"assessment/internal/klippa-api"
	ROH "assessment/internal/responseOutputHandling"
	"assessment/internal/structs"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Initialize the flags, this is done through the flag package provided by GO
	apiKey := flag.String("api", "[REQUIRED]", "a api key")
	template := flag.String("template", "financial_full", "which default provided template to use")
	extractionType := flag.String("textextraction", "fast", "what thype of text extraction you want to use, needs to be fast or full")
	filepath := flag.String("file", "[REQUIRED]", "a document or image file path")
	savefile := flag.String("save", "filename", "name of how you want to save the result")
	debug := flag.Bool("debug", false, "enable the debugmode, this wont send any requests to the api and will work with a static json file in this folder")
	fullOutput := flag.Bool("fulloutput", false, "get the full output nicely formated, or have it only show the data that is filled in.")

	flag.Parse()

	if *debug {
		// If debug is turned on we load from an exampleResponse.json, this is done to not use credits while testing the output functionality of this program.
		file, _ := os.ReadFile("testResponse/exampleResponse.json")
		ROH.PrintResponse(file, 200, *fullOutput)
	} else {
		// Normal flow.
	
		// Check if the user did fill in the required flags.
		if string(*filepath) == "[REQUIRED]" {
			fmt.Println("Please give a file to parse, use -file=[path/to/filename.extension] to do so.")
			os.Exit(0)
		}
		if string(*apiKey) == "[REQUIRED]" {
			fmt.Println("Please fill in a api key to use in the request, use -api=[apikey] to do so. ")
			os.Exit(0)
		}

		// create construct from the input flags.
		requestconfig := new(structs.RequestConfig)
		requestconfig.ApiKey = string(*apiKey)
		requestconfig.Template = string(*template)
		requestconfig.ExtractionType = string(*extractionType)
		requestconfig.FilePath = string(*filepath)
		requestconfig.SavefileName = string(*savefile)

		// check if the file mentioned in the filepath exists
		if _, err := os.Stat(string(*filepath)); errors.Is(err, os.ErrNotExist) {
			fmt.Println("This file does not exist: ", *filepath)
			os.Exit(0)
		}

		// give the request config to the ParseDocument function, this will execute the api call with all the config stored in the requestConfig struct.
		response := klippaApi.ParseDocument(requestconfig)

		// Read the body.
		bodyData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Save ReadAll error: ", err)
			os.Exit(0)
		}

		// check if the given savefile name isnt the default one. if it is different we save it to a file.
		if requestconfig.SavefileName != "filename" {
			ROH.SaveResponse(bodyData, *requestconfig)
		}

		// print the response to the console.
		ROH.PrintResponse(bodyData, response.StatusCode, *fullOutput)

	}

}
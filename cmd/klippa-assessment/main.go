package main

import (
	"assessment/internal/structs"
	"assessment/internal/klippa-api"
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {

	ApiKey := flag.String("api", "[REQUIRED]", "a api key")
	Template := flag.String("template", "financial_full", "which default provided template to use")
	ExtractionType := flag.String("textextraction", "fast", "what thype of text extraction you want to use, needs to be fast or full")
	FilePath := flag.String("file", "[REQUIRED]", "a document or image file path")
	SaveFile := flag.String("save", "output", "name of how you want to save the result")

	flag.Parse()

	// print all flags into the console.
	// fmt.Println("ApiKey", *ApiKey)
	// fmt.Println("Template", *Template)
	// fmt.Println("ExtractionType", *ExtractionType)
	// fmt.Println("FilePath", *FilePath)
	// fmt.Println("SaveFile", *SaveFile)

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

	result, err := klippaApi.ParseDocument(requestconfig)

	fmt.Println(err)
	fmt.Println(result)

}

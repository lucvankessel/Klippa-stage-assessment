package main

import (
	// "assessment/internal/structs"
	"flag"
	"fmt"
)

func main() {

	ApiKey := flag.String("api", "[key provided by klippa]", "a api key")
	Template := flag.String("template", "financial_full", "which default provided template to use")
	TextExtractionType := flag.String("textextraction", "fast", "what thype of text extraction you want to use, needs to be fast or full")
	FilePath := flag.String("file", "file/to/document.pdf", "a document or image file path")
	SaveFile := flag.String("save", "filename", "name of how you want to save the result")

	flag.Parse()

	fmt.Println("ApiKey", *ApiKey)
	fmt.Println("Template", *Template)
	fmt.Println("TextExtractionType", *TextExtractionType)
	fmt.Println("FilePath", *FilePath)
	fmt.Println("SaveFile", *SaveFile)
}

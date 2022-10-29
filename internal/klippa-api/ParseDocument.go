package klippaApi

import (
	"assessment/internal/structs"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Calls the klippa parsedocument api
func ParseDocument( requestconfig *structs.RequestConfig ) (*http.Response, error) {

	// Opens the provided file that needs parsing
	file, _ := os.Open(requestconfig.FilePath)
	defer file.Close()

	// Writing all the body elements of the request.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("document", filepath.Base(file.Name()))
	io.Copy(part, file)
	extraction, err := writer.CreateFormField("pdf_text_extraction")
	if err != nil {
		return nil, err
	}
	extraction.Write([]byte(requestconfig.ExtractionType))
	writer.Close()

	// Creating a request and giving it some headers that are needed.
	r,_ := http.NewRequest(
		"POST",
		fmt.Sprintf("https://custom-ocr.klippa.com/api/v1/parseDocument/%s", requestconfig.Template),
		body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Auth-Key", requestconfig.ApiKey)

	// Execute the api call.
    resp, err := http.DefaultClient.Do(r)
    if err != nil {
        fmt.Println(err)
		os.Exit(0)
    }

	return resp, err
}
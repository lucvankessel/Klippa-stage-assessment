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

// this function will make the parsedocument request to the klippa api
func ParseDocument( requestconfig *structs.RequestConfig ) (*http.Response, error) {


	file, _ := os.Open(requestconfig.FilePath)

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("document", filepath.Base(file.Name()))
	io.Copy(part, file)
	
	extraction, err := writer.CreateFormField("pdf_text_extraction")
	if err != nil {
		return nil, err
	}
	extraction.Write([]byte(fmt.Sprintf(requestconfig.ExtractionType)))
	writer.Close()

	r,_ := http.NewRequest(
		"POST",
		fmt.Sprintf("https://custom-ocr.klippa.com/api/v1/parseDocument/%s", requestconfig.Template),
		body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Auth-Key", requestconfig.ApiKey)

    resp, err := http.DefaultClient.Do(r)
    if err != nil {
        fmt.Println(err)
    }

	return resp, err
}
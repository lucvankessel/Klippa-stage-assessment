package structs


type Result200 struct {

	Data struct {
		Hash string `json:"hash"`
		HashDuplicate bool `json:"hash_duplicate"`
		Parsed *string `json:"parsed"`
		Quality struct {
			Blurriness int `json:"blurriness"`
			Characters int `json:"characters"`
			NonTextCharacters int `json:"non_text_chars"`
			OverExposed int `json:"over_exposed"`
			UnderExposed int `json:"under_exposed"`
			UnknownDocumentType int `json:"unknown_document_type"`
			UnsupportedChars int `json:"unsupported_chars"`
		} `json:"quality"`
		RawText string `json:"raw_text"`
	} `json:"data"`

	Request_id string `json:"request_id"`
	Result string `json:"result"`

}

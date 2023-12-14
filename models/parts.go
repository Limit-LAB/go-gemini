package models

type TextParts struct {
	Text string `json:"text"`
}

func (_ TextParts) parts() {}

type InlineDataParts struct {
	InlineData struct {
		MimeType string `json:"mime_type"`
		Data     string `json:"data"`
	} `json:"inline_data"`
}

func (_ InlineDataParts) parts() {}

type IParts interface {
	parts()
}

var _ IParts = TextParts{}
var _ IParts = InlineDataParts{}

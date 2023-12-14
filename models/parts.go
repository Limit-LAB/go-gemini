package models

type TextPart struct {
	Text string `json:"text"`
}

func (_ TextPart) part() {}

type InlineDataPart struct {
	InlineData InlineData `json:"inline_data"`
}

type InlineData struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

func (_ InlineDataPart) part() {}

type IPart interface {
	part()
}

var _ IPart = TextPart{}
var _ IPart = InlineDataPart{}

type Parts []IPart

func NewParts(parts []IPart) Parts {
	return parts
}

func (p Parts) AppendPart(part IPart) Parts {
	return append(p, part)
}

func NewTextPart(text string) TextPart {
	return TextPart{
		Text: text,
	}
}

func NewInlineDataPart(mimeType, data string) InlineDataPart {
	return InlineDataPart{
		InlineData: InlineData{
			MimeType: mimeType,
			Data:     data,
		},
	}
}

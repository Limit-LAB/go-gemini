package models

type Part struct {
	Text       *string     `json:"text,omitempty"`
	InlineData *InlineData `json:"inline_data,omitempty"`
}

func (p Part) IsText() bool {
	return p.Text != nil
}

func (p Part) IsInlineData() bool {
	return p.InlineData != nil
}

func (p Part) GetText() string {
	if p.Text == nil {
		return ""
	}
	return *p.Text
}

func (p Part) GetInlineData() InlineData {
	if p.InlineData == nil {
		return InlineData{}
	}
	return *p.InlineData
}

type InlineData struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type Parts []Part

func NewParts(parts ...Part) Parts {
	return parts
}

func (p Parts) AppendPart(part ...Part) Parts {
	return append(p, part...)
}

func NewTextPart(text string) Part {
	return Part{
		Text: &text,
	}
}

func NewInlineDataPart(mimeType, data string) Part {
	return Part{
		InlineData: &InlineData{
			MimeType: mimeType,
			Data:     data,
		},
	}
}

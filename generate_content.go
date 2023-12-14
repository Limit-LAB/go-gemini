package gemini

import (
	"errors"
	"github.com/Limit-Lab/gemini/models"
)

func (c *Client) GenerateContent(model models.GeminiModel, req models.GenerateContentRequest) (models.GenerateContentResponse, error) {
	for _, content := range req.Contents {
		err := validateGenerateContentRequest(model, content)
		if err != nil {
			return models.GenerateContentResponse{}, err
		}
	}
	url := c.url(string(model), "generateContent")
	return post[models.GenerateContentResponse](c.hc, url, req)
}

var ErrTextOnlyModel = errors.New("this model only supports text input")

func validateGenerateContentRequest(model models.GeminiModel, req models.Content) error {
	if model != models.GeminiPro {
		return nil
	}
	for _, c := range req.Parts {
		switch c.(type) {
		case models.TextParts:
			continue
		default:
			return ErrTextOnlyModel
		}
	}
	return nil
}

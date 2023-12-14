package gemini

import (
	"errors"
	"github.com/Limit-LAB/go-gemini/models"
)

func (c *Client) GenerateContent(model models.GeminiModel, req models.GenerateContentRequest) (models.GenerateContentResponse, error) {
	for _, content := range req.Contents {
		err := validateGenerateContentRequest(model, content)
		if err != nil {
			return models.GenerateContentResponse{}, err
		}
	}
	url := c.url(string(model), "generateContent")
	rst, err := unjson[models.GenerateContentResponse](c.post(url, req))
	return rst, err
}

var ErrTextOnlyModel = errors.New("this model only supports text input")

func validateGenerateContentRequest(model models.GeminiModel, req models.Content) error {
	if model != models.GeminiPro {
		return nil
	}
	for _, c := range req.Parts {
		if c.Text != nil {
			continue
		}
		return ErrTextOnlyModel
	}
	return nil
}

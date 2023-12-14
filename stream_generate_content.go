package gemini

import "github.com/Limit-Lab/go-gemini/models"

func (c *Client) GenerateContentStream(model models.GeminiModel, req models.GenerateContentRequest) (models.GenerateContentResponse, error) {
	for _, content := range req.Contents {
		err := validateGenerateContentRequest(model, content)
		if err != nil {
			return models.GenerateContentResponse{}, err
		}
	}
	url := c.url(string(model), "streamGenerateContent")
	return unjson[models.GenerateContentResponse](c.post(url, req))
	// TODO: implement stream
}

type GenerateContentStream struct {
}

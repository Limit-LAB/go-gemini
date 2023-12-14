package gemini

import "github.com/Limit-Lab/go-gemini/models"

func (c *Client) GetModelInfo(model models.GeminiModel) (models.ModelInfo, error) {
	url := c.url(string(model), "")
	return get[models.ModelInfo](c.hc, url)
}

func (c *Client) GetModelList() ([]models.ModelInfo, error) {
	url := c.url("models", "")
	lst, err := get[models.ModelListResponse](c.hc, url)
	if err != nil {
		return nil, err
	}
	return lst.Models, nil
}

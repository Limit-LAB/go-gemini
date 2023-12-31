package gemini

import "github.com/Limit-LAB/go-gemini/models"

func (c *Client) GetModelInfo(model models.GeminiModel) (models.ModelInfo, error) {
	url := c.url(string(model), "")
	return unjson[models.ModelInfo](c.get(url))
}

func (c *Client) GetModelList() ([]models.ModelInfo, error) {
	lst, err := unjson[models.ModelListResponse](c.get(c.baseUrl))
	return lst.Models, err
}

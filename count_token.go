package gemini

import (
	"github.com/Limit-LAB/go-gemini/models"
)

func (c *Client) CountToken(model models.GeminiModel, req models.CountTokenRequest) (rst models.CountTokenResponse, err error) {
	url := c.url(string(model), "countToken")
	rst, err = unjson[models.CountTokenResponse](c.post(url, req))
	return
}

package gemini

import (
	"github.com/Limit-Lab/gemini/models"
)

func (c *Client) CountToken(model models.GeminiModel, req models.CountTokenRequest) (rst models.CountTokenResponse, err error) {
	url := c.url(string(model), "countToken")
	rst, err = post[models.CountTokenResponse](c.hc, url, req)
	return
}

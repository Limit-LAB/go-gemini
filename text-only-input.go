package gemini

import "github.com/Limit-Lab/gemini/models"

func (c *Client) TextOnlyInput(model models.GeminiModel) {
	_ = c.baseUrl + string(model) + ":generateContent?key=" + c.key

}

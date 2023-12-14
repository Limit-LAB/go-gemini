package gemini

import "net/http"

type Client struct {
	hc      *http.Client
	key     string
	baseUrl string
}

func NewClient(key string) *Client {
	return &Client{
		hc:      &http.Client{},
		key:     key,
		baseUrl: "https://generativelanguage.googleapis.com/v1beta/models/",
	}
}

func (c *Client) SetBaseUrl(url string) *Client {
	c.baseUrl = url
	return c
}

func (c *Client) keyParam() string {
	return "?key=" + c.key
}

func (c *Client) url(model, action string) string {
	if action == "" {
		return c.baseUrl + model + c.keyParam()
	}
	return c.baseUrl + model + ":" + action + c.keyParam()
}

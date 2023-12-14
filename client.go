package gemini

import "net/http"

type AuthBy string

const (
	AuthByHttpHeader AuthBy = "HttpHeader"
	AuthByUrlQuery   AuthBy = "UrlQuery"
)

type Client struct {
	hc      *http.Client
	key     string
	baseUrl string
	auth    AuthBy // Vertex AI uses HTTP HEAD instead of Query
}

func NewClient(key string) *Client {
	return &Client{
		hc:      &http.Client{},
		key:     key,
		baseUrl: "https://generativelanguage.googleapis.com/v1beta/models/",
		auth:    AuthByUrlQuery,
	}
}

func (c *Client) SetBaseUrl(url string) *Client {
	c.baseUrl = url
	return c
}

func (c *Client) SetAuthWay(auth AuthBy) *Client {
	c.auth = auth
	return c
}

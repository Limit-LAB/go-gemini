package gemini

import (
	"github.com/Limit-LAB/go-gemini/models"
	"net/http"
)

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

func (c *Client) VertexAI(region models.VertexRegion, projId string) *Client {
	rg := string(region)
	c.auth = AuthByHttpHeader
	c.baseUrl = "https://" + rg + "-aiplatform.googleapis.com/v1/projects/" + projId + "/locations/" + rg + "/publishers/google/models/"
	return c
}

func (c *Client) AIStudio() *Client {
	c.auth = AuthByUrlQuery
	c.baseUrl = "https://generativelanguage.googleapis.com/v1beta/models/"
	return c
}

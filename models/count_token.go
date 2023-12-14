package models

type CountTokenRequest struct {
	Contents []Content `json:"contents"`
}

type CountTokenResponse struct {
	TotalTokens int `json:"totalTokens"`
}

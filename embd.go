package gemini

import (
	"github.com/Limit-LAB/go-gemini/models"
)

func (c *Client) EmbedContent(model models.EmbeddingModel, parts []models.Part) (models.EmbeddingValue, error) {
	url := c.url(string(model), "embedContent")
	req := models.EmbeddingContentRequest{
		Model: "models/" + string(model),
		Content: models.Content{
			Parts: parts,
		},
	}
	rst, err := unjson[models.EmbeddingContentResponse](c.post(url, req))
	return rst.Embedding, err
}

func (c *Client) BatchEmbedContent(model models.EmbeddingModel, parts [][]models.Part) ([]models.EmbeddingValue, error) {
	url := c.url(string(model), "batchEmbedContents")
	req := models.BatchEmbeddingContentsRequest{}
	modelStr := "models/" + string(model)
	for _, content := range parts {
		req.Requests = append(req.Requests, models.EmbeddingContentRequest{
			Model: modelStr,
			Content: models.Content{
				Parts: content,
			},
		})
	}
	rst, err := unjson[models.BatchEmbeddingContentsResponse](c.post(url, req))
	return rst.Embeddings, err
}

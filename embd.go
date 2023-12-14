package gemini

import "github.com/Limit-Lab/gemini/models"

func (c *Client) EmbedContent(model models.EmbeddingModel, content []models.IParts) (models.EmbeddingValue, error) {
	url := c.url(string(model), "embedContent")
	req := models.EmbeddingContentRequest{
		Model: "models/" + string(model),
		Content: models.Content{
			Parts: content,
		},
	}
	resp, err := post[models.EmbeddingContentResponse](c.hc, url, req)
	if err != nil {
		return models.EmbeddingValue{}, err
	}
	return resp.Embedding, nil
}

func (c *Client) BatchEmbedContent(model models.EmbeddingModel, contents [][]models.IParts) ([]models.EmbeddingValue, error) {
	url := c.url(string(model), "batchEmbedContents")
	req := models.BatchEmbeddingContentsRequest{}
	modelStr := "models/" + string(model)
	for _, content := range contents {
		req.Requests = append(req.Requests, models.EmbeddingContentRequest{
			Model: modelStr,
			Content: models.Content{
				Parts: content,
			},
		})
	}

	resp, err := post[models.BatchEmbeddingContentsResponse](c.hc, url, req)
	if err != nil {
		return nil, err
	}
	return resp.Embeddings, nil
}

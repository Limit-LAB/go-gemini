package models

type EmbeddingContentRequest struct {
	Model   string  `json:"model"`
	Content Content `json:"content"`
}

type EmbeddingValue struct {
	Values []float32 `json:"values"`
}

type EmbeddingContentResponse struct {
	Embedding EmbeddingValue `json:"embedding"`
}

type BatchEmbeddingContentsRequest struct {
	Requests []EmbeddingContentRequest `json:"requests"`
}

type BatchEmbeddingContentsResponse struct {
	Embeddings []EmbeddingValue `json:"embeddings"`
}

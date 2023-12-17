package models

type ModelListResponse struct {
	Models []ModelInfo `json:"models"`
}

type ModelInfo struct {
	Name                       string   `json:"name"`
	Version                    string   `json:"version"`
	DisplayName                string   `json:"displayName"`
	Description                string   `json:"description"`
	InputTokenLimit            int      `json:"inputTokenLimit"`
	OutputTokenLimit           int      `json:"outputTokenLimit"`
	SupportedGenerationMethods []string `json:"supportedGenerationMethods"`
	Temperature                float32  `json:"temperature,omitempty"`
	TopP                       float32  `json:"topP,omitempty"`
	TopK                       int      `json:"topK,omitempty"`
}

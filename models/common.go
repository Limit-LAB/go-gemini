package models

type PromptFeedback struct {
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  Role   `json:"role,omitempty"`
}

type GenerationConfig struct {
	StopSequences   []string `json:"stopSequences,omitempty"`
	Temperature     *float32 `json:"temperature,omitempty"`
	MaxOutputTokens *int     `json:"maxOutputTokens,omitempty"`
	TopP            *float32 `json:"topP,omitempty"`
	TopK            *float32 `json:"topK,omitempty"`
}

func NewGenerationConfig() *GenerationConfig {
	return &GenerationConfig{}
}

func (g *GenerationConfig) WithStopSequences(stopSequences ...string) *GenerationConfig {
	g.StopSequences = stopSequences
	return g
}

func (g *GenerationConfig) WithTemperature(temperature float32) *GenerationConfig {
	g.Temperature = &temperature
	return g
}

func (g *GenerationConfig) WithMaxOutputTokens(maxOutputTokens int) *GenerationConfig {
	g.MaxOutputTokens = &maxOutputTokens
	return g
}

func (g *GenerationConfig) WithTopP(topP float32) *GenerationConfig {
	g.TopP = &topP
	return g
}

func (g *GenerationConfig) WithTopK(topK float32) *GenerationConfig {
	g.TopK = &topK
	return g
}

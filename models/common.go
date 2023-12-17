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
	StopSequences   []string `json:"stopSequences"`
	Temperature     float32  `json:"temperature"`
	MaxOutputTokens int      `json:"maxOutputTokens"`
	TopP            float32  `json:"topP"`
	TopK            float32  `json:"topK"`
}

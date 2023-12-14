package models

type PromptFeedback struct {
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type Content struct {
	Parts []IPart `json:"parts"`
	Role  Role    `json:"role,omitempty"`
}

type GenerationConfig struct {
	StopSequences   []string `json:"stopSequences"`
	Temperature     float64  `json:"temperature"`
	MaxOutputTokens int      `json:"maxOutputTokens"`
	TopP            float64  `json:"topP"`
	TopK            int      `json:"topK"`
}

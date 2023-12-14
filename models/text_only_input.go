package models

type TextOnlyInputRequest struct {
	Candidates       []TextOnlyCandidate `json:"candidates"`
	PromptFeedback   PromptFeedback      `json:"promptFeedback"`
	GenerationConfig *GenerationConfig   `json:"generationConfig"`
}

type TextOnlyCandidate struct {
	Content       Content        `json:"content"`
	FinishReason  string         `json:"finishReason"`
	Index         int            `json:"index"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type PromptFeedback struct {
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type Content struct {
	Parts []IParts `json:"parts"`
	Role  Role     `json:"role"`
}

type GenerationConfig struct {
	StopSequences   []string `json:"stopSequences"`
	Temperature     float64  `json:"temperature"`
	MaxOutputTokens int      `json:"maxOutputTokens"`
	TopP            float64  `json:"topP"`
	TopK            int      `json:"topK"`
}

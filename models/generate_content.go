package models

type GenerateContentRequest struct {
	Contents         []Content
	GenerationConfig *GenerationConfig `json:"generationConfig"`
}

type GenerateContentResponse struct {
	Candidates     []GenerateContentCandidate `json:"candidates"`
	PromptFeedback PromptFeedback             `json:"promptFeedback"`
}

type GenerateContentCandidate struct {
	Content       Content        `json:"content"`
	FinishReason  string         `json:"finishReason"`
	Index         int            `json:"index"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

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

func NewGenerateContentRequest(contents []Content) GenerateContentRequest {
	return GenerateContentRequest{
		Contents: contents,
	}
}

func (r GenerateContentRequest) WithGenerationConfig(config GenerationConfig) GenerateContentRequest {
	r.GenerationConfig = &config
	return r
}

func (r GenerateContentRequest) AppendContent(role Role, parts []IPart) GenerateContentRequest {
	r.Contents = append(r.Contents, Content{
		Role:  role,
		Parts: parts,
	})
	return r
}

package models

type GenerateContentRequest struct {
	Contents         []Content         `json:"contents"`
	GenerationConfig *GenerationConfig `json:"generationConfig,omitempty"`
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

func NewGenerateContentRequest(contents ...Content) GenerateContentRequest {
	return GenerateContentRequest{
		Contents: contents,
	}
}

func NewGenerateContentRequestWithConfig(cfg *GenerationConfig, contents ...Content) GenerateContentRequest {
	return GenerateContentRequest{
		Contents:         contents,
		GenerationConfig: cfg,
	}
}

func (r GenerateContentRequest) WithGenerationConfig(config GenerationConfig) GenerateContentRequest {
	r.GenerationConfig = &config
	return r
}

func NewContent(role Role, parts ...Part) Content {
	return Content{
		Role:  role,
		Parts: parts,
	}
}

func (r GenerateContentRequest) AppendContent(role Role, parts ...Part) GenerateContentRequest {
	r.Contents = append(r.Contents, Content{
		Role:  role,
		Parts: parts,
	})
	return r
}

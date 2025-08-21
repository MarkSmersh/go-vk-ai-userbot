package openai

type Response struct {
	ID                 string            `json:"id"`
	Object             string            `json:"object"`
	CreatedAt          int64             `json:"created_at"`
	Status             string            `json:"status"`
	Error              *ErrorResponse    `json:"error,omitempty"`
	IncompleteDetails  map[string]any    `json:"incomplete_details,omitempty"`
	Instructions       any               `json:"instructions,omitempty"` // can be string or []string
	MaxOutputTokens    *int              `json:"max_output_tokens,omitempty"`
	MaxToolCalls       *int              `json:"max_tool_calls,omitempty"`
	Model              string            `json:"model"`
	Output             []Output          `json:"output"`
	OutputText         *string           `json:"output_text,omitempty"`
	ParallelToolCalls  bool              `json:"parallel_tool_calls"`
	PreviousResponseID *string           `json:"previous_response_id,omitempty"`
	Prompt             map[string]any    `json:"prompt,omitempty"`
	Reasoning          map[string]any    `json:"reasoning,omitempty"`
	ServiceTier        *string           `json:"service_tier,omitempty"`
	Store              bool              `json:"store"`
	Temperature        *float64          `json:"temperature,omitempty"`
	Text               map[string]any    `json:"text,omitempty"`
	ToolChoice         any               `json:"tool_choice,omitempty"` // string or object
	Tools              []map[string]any  `json:"tools"`
	TopLogprobs        *int              `json:"top_logprobs,omitempty"`
	TopP               *float64          `json:"top_p,omitempty"`
	Truncation         *string           `json:"truncation,omitempty"`
	Usage              *TokenUsage       `json:"usage,omitempty"`
	User               *string           `json:"user,omitempty"`
	Metadata           map[string]string `json:"metadata"`
}

type Output struct {
	Type    string        `json:"type"`
	ID      string        `json:"id"`
	Status  string        `json:"status"`
	Role    string        `json:"role"`
	Content []OutputBlock `json:"content"`
}

type OutputBlock struct {
	Type        string `json:"type"`
	Text        string `json:"text,omitempty"`
	Annotations []any  `json:"annotations,omitempty"`
}

type ErrorResponse struct {
	// Define if needed based on error fields (code, message, etc.)
}

type TokenUsage struct {
	InputTokens         int            `json:"input_tokens"`
	InputTokensDetails  map[string]int `json:"input_tokens_details"`
	OutputTokens        int            `json:"output_tokens"`
	OutputTokensDetails map[string]int `json:"output_tokens_details"`
	TotalTokens         int            `json:"total_tokens"`
}

package deepseek

type Response struct {
	ID                string   `json:"id"`                 // Unique identifier for the chat completion
	Choices           []Choice `json:"choices"`            // List of chat completion choices
	Created           int64    `json:"created"`            // Unix timestamp of creation
	Model             string   `json:"model"`              // The model used
	SystemFingerprint string   `json:"system_fingerprint"` // Backend configuration fingerprint
	Object            string   `json:"object"`             // Always "chat.completion"
	Usage             Usage    `json:"usage"`              // Usage statistics
}

// Choice represents a single completion choice
type Choice struct {
	FinishReason string        `json:"finish_reason"`      // Reason the model stopped (stop, length, etc.)
	Index        int           `json:"index"`              // Index of this choice
	Message      ChoiceMessage `json:"message"`            // Generated message
	Logprobs     *Logprobs     `json:"logprobs,omitempty"` // Optional. Log probability information
}

// ChoiceMessage represents the generated assistant message
type ChoiceMessage struct {
	Content          string     `json:"content"`                     // Nullable. Message content
	ReasoningContent string     `json:"reasoning_content,omitempty"` // Nullable. Reasoning content (deepseek-reasoner only)
	ToolCalls        []ToolCall `json:"tool_calls,omitempty"`        // Optional. Tool calls
	Role             string     `json:"role"`                        // Always "assistant"
}

// ToolCall represents a tool call object
type ToolCall struct {
	// Define fields according to API schema for tool calls
}

// Logprobs contains log probability details for tokens
type Logprobs struct {
	Content []LogprobContent `json:"content,omitempty"` // Nullable. List of token logprobs
}

// LogprobContent represents log probability info for a token
type LogprobContent struct {
	Token       string       `json:"token"`        // Token text
	Logprob     float64      `json:"logprob"`      // Log probability (or -9999.0 if very unlikely)
	Bytes       []int        `json:"bytes"`        // Nullable. UTF-8 byte representation
	TopLogprobs []TopLogprob `json:"top_logprobs"` // Top likely tokens
}

// TopLogprob represents one of the most likely tokens
type TopLogprob struct {
	Token   string  `json:"token"`   // Candidate token
	Logprob float64 `json:"logprob"` // Log probability
}

// Usage contains statistics for token usage
type Usage struct {
	CompletionTokens        int                     `json:"completion_tokens"`         // Tokens in generated completion
	PromptTokens            int                     `json:"prompt_tokens"`             // Tokens in prompt
	PromptCacheHitTokens    int                     `json:"prompt_cache_hit_tokens"`   // Prompt tokens that hit cache
	PromptCacheMissTokens   int                     `json:"prompt_cache_miss_tokens"`  // Prompt tokens that missed cache
	TotalTokens             int                     `json:"total_tokens"`              // Total tokens used
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"` // Breakdown of completion tokens
}

// CompletionTokensDetails provides breakdown of completion token usage
type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens,omitempty"` // Tokens used for reasoning
}

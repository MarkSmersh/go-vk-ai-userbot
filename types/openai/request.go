package openai

type Request struct {
	Background        bool              `json:"background,omitempty"`
	Include           []string          `json:"include,omitempty"`
	Input             []Input           `json:"input,omitempty"`    // can be string, []any, or other
	Messages          []Input           `json:"messages,omitempty"` // used by Deepseek
	Instructions      string            `json:"instructions,omitempty"`
	MaxOutputTokens   int               `json:"max_output_tokens,omitempty"`
	MaxToolCalls      int               `json:"max_tool_calls,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	Model             string            `json:"model,omitempty"`
	ParallelToolCalls bool              `json:"parallel_tool_calls,omitempty"`
	PreviousResponse  string            `json:"previous_response_id,omitempty"`
	Prompt            map[string]any    `json:"prompt,omitempty"`    // structure depends on template
	Reasoning         map[string]any    `json:"reasoning,omitempty"` // o-model specific
	ServiceTier       string            `json:"service_tier,omitempty"`
	Store             bool              `json:"store,omitempty"`
	Stream            bool              `json:"stream,omitempty"`
	Temperature       float64           `json:"temperature,omitempty"`
	Text              map[string]any    `json:"text,omitempty"`        // can also be a concrete struct
	ToolChoice        any               `json:"tool_choice,omitempty"` // string or object
	Tools             []map[string]any  `json:"tools,omitempty"`       // can be concrete tool type
	TopLogprobs       int               `json:"top_logprobs,omitempty"`
	TopP              float64           `json:"top_p,omitempty"`
	Truncation        string            `json:"truncation,omitempty"`
	User              string            `json:"user,omitempty"`
}

type Input struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

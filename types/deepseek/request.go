package deepseek

type Request struct {
	Messages         []Message       `json:"messages"`                    // Required. List of messages in the conversation
	Model            string          `json:"model"`                       // Required. ID of the model to use. E.g., "deepseek-chat"
	FrequencyPenalty float64         `json:"frequency_penalty,omitempty"` // Optional. Number between -2.0 and 2.0
	MaxTokens        int             `json:"max_tokens,omitempty"`        // Optional. Maximum number of tokens to generate
	PresencePenalty  float64         `json:"presence_penalty,omitempty"`  // Optional. Number between -2.0 and 2.0
	ResponseFormat   *ResponseFormat `json:"response_format,omitempty"`   // Optional. Response format object
	Stop             *Stop           `json:"stop,omitempty"`              // Optional. Stop object
	Stream           bool            `json:"stream,omitempty"`            // Optional. Whether to stream responses
	StreamOptions    *StreamOptions  `json:"stream_options,omitempty"`    // Optional. Stream options
	Temperature      float64         `json:"temperature,omitempty"`       // Optional. Sampling temperature (0–2)
	TopP             float64         `json:"top_p,omitempty"`             // Optional. Nucleus sampling (≤ 1)
	Tools            []Tool          `json:"tools,omitempty"`             // Optional. List of tools
	ToolChoice       *ToolChoice     `json:"tool_choice,omitempty"`       // Optional. Tool choice object
	Logprobs         bool            `json:"logprobs,omitempty"`          // Optional. Whether to return log probabilities
	TopLogprobs      int             `json:"top_logprobs,omitempty"`      // Optional. Number of top tokens to return (requires logprobs = true)
}

// Message represents a single message in the chat
type Message struct {
	Role    string `json:"role"`           // "system", "user", "assistant", etc.
	Content string `json:"content"`        // The content of the message
	Name    string `json:"name,omitempty"` // An optional name for the participant.
}

// ResponseFormat defines how the model should format responses
type ResponseFormat struct {
	Type string `json:"type,omitempty"` // Optional. Format type
}

// Stop defines stopping criteria for the model
type Stop struct {
	Sequence  string   `json:"sequence,omitempty"`  // Optional. Stop sequence
	Sequences []string `json:"sequences,omitempty"` // Optional. Multiple stop sequences
}

// StreamOptions defines options for streaming responses
type StreamOptions struct {
	// Add fields if the API specifies any streaming options
}

// Tool represents a tool definition
type Tool struct {
	// Define tool fields according to API
}

// ToolChoice represents a specific tool selection
type ToolChoice struct {
	// Define fields according to API
}

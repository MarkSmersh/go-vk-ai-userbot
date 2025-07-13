package events

type Event struct {
	Ts                      int     `json:"ts,omitempty"`
	Updates                 [][]any `json:"updates,omitempty"`
	Failed                  int     `json:"failed,omitempty"`
	Error                   string  `json:"error,omitempty"`
	RedirectURI             string  `json:"redirect_uri,omitempty"`               // VK captcha redirect URL
	CaptchaSID              string  `json:"captcha_sid,omitempty"`                // Captcha session ID
	IsRefreshEnabled        bool    `json:"is_refresh_enabled,omitempty"`         // Whether captcha refresh is enabled
	CaptchaImg              string  `json:"captcha_img,omitempty"`                // Captcha image URL
	CaptchaTS               string  `json:"captcha_ts,omitempty"`                 // Timestamp for captcha generation
	CaptchaAttempt          int     `json:"captcha_attempt,omitempty"`            // Attempt count for solving captcha
	CaptchaRatio            float64 `json:"captcha_ratio,omitempty"`              // Display ratio or scale for captcha
	IsSoundCaptchaAvailable bool    `json:"is_sound_captcha_available,omitempty"` // If sound captcha option is available
}

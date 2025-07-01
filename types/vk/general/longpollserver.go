package general

type LongPollServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
	Pts    string `json:"pts,omitempty"`
}

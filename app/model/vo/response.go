package vo

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

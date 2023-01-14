package schema

type Like struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	VideoId    int    `json:"video_id"`
	ActionType string `json:"action_type"`
}

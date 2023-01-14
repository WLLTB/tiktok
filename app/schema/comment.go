package schema

import "time"

type Comment struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	VideoId     int    `json:"video_id"`
	CommentText string `json:"comment_text"`
	CreateDate  time.Time
	ActionType  string `json:"action_type"`
}

package schema

type Follow struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	FollowerId int    `json:"follower_id"`
	ActionType string `json:"action_type"`
}

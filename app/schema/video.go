package schema

import "time"

type Video struct {
	Id          int    `json:"id"`
	AuthorId    int    `json:"author_id"`
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	PublishTime time.Time
	Title       string `json:"title"`
}

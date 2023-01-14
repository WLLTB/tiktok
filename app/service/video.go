package service

import (
	"strconv"
	"tiktok/app/repository"
	. "tiktok/app/schema"
	"tiktok/app/vo"
	. "time"
)

func GetVideos(time string, count int) ([]Video, error) {
	lastTime := parseTime(time)

	videos := repository.PageVideoListByTime(lastTime, count)
	return videos, nil
}

func ConvertVideoList(userId int, lastTime string, count int) ([]vo.Video, error) {
	rawVideos, err := GetVideos(lastTime, count)
	if err != nil {
		return nil, err
	}
	likeList := repository.GetLikeListByUserId(userId)

	var videos []vo.Video

	for _, rawVideo := range rawVideos {
		authorInfo, err := GetAuthorInfo(userId, int(rawVideo.AuthorId))
		if err != nil {
			return nil, err
		}

		videos = append(videos, vo.Video{
			Id:            rawVideo.Id,
			Author:        authorInfo,
			PlayUrl:       rawVideo.PlayUrl,
			CoverUrl:      rawVideo.CoverUrl,
			FavoriteCount: repository.GetLikedCountByVideoId(rawVideo.Id),
			CommentCount:  repository.GetCommentCountByVideoId(rawVideo.Id),
			IsFavorite:    isFavorite(likeList, int(rawVideo.Id)),
			Title:         rawVideo.Title,
		})
	}
	return videos, nil
}

func parseTime(time string) Time {
	if time == "" {
		return Now()
	}
	me, _ := strconv.ParseInt(time, 10, 64)
	return Unix(me, 0)
}

func isFavorite(likeList []Like, videoId int) bool {
	for _, like := range likeList {
		if like.VideoId == videoId {
			return true
		}
	}
	return false
}

package service

import (
	"strconv"
	"sync"
	"tiktok/app/repository"
	. "tiktok/app/schema"
	"tiktok/app/vo"
	. "time"
)

func SupplementFeedVideoList(userId int64, lastTime string, count int64) ([]vo.Video, error) {
	rawVideos := repository.PageVideoListByTime(parseTime(lastTime), count)

	videoList, err := buildVideos(userId, rawVideos)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}

func SupplementLikeVideoList(userId int64) ([]vo.Video, error) {
	rawVideos := repository.GetLikeVideoList(userId)

	videoList, err := buildVideos(userId, rawVideos)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}

func buildVideos(userId int64, rawVideos []Video) ([]vo.Video, error) {
	// 查出用户喜欢列表
	likeList := repository.GetLikeListByUserId(userId)

	var wg sync.WaitGroup
	wg.Add(len(rawVideos))

	videoChan := make(chan vo.Video, len(rawVideos))

	for _, rawVideo := range rawVideos {
		go func(rawVideo Video) {
			defer wg.Done()
			video := vo.Video{
				Id:            rawVideo.Id,
				Author:        SupplementTargetUserInfo(userId, rawVideo.AuthorId),
				PlayUrl:       rawVideo.PlayUrl,
				CoverUrl:      rawVideo.CoverUrl,
				FavoriteCount: repository.CountLikedByVideoId(rawVideo.Id),
				CommentCount:  repository.CountCommentByVideoId(rawVideo.Id),
				IsFavorite:    checkIsFavorite(likeList, rawVideo.Id),
				Title:         rawVideo.Title,
			}
			videoChan <- video
		}(rawVideo)
	}
	wg.Wait()
	close(videoChan)

	var videos []vo.Video
	for video := range videoChan {
		videos = append(videos, video)
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

func checkIsFavorite(likeList []Like, videoId int64) bool {
	for _, like := range likeList {
		if like.VideoId == videoId {
			return true
		}
	}
	return false
}

package service

import (
	"strconv"
	"sync"
	"tiktok/app/repository"
	. "tiktok/app/schema"
	"tiktok/app/vo"
	. "time"
)

func SupplementVideoList(userId int64, lastTime string, count int64) ([]vo.Video, error) {
	// 数据库中查出来的数据，后续要转换成所需要的字段
	rawVideos := repository.PageVideoListByTime(parseTime(lastTime), count)

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
				Author:        supplementTargetUserInfo(userId, rawVideo.AuthorId),
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

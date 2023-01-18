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
	likeList := repository.GetLikeListByUserId(userId)
	likeMap := make(map[int64]bool)
	for _, like := range likeList {
		likeMap[like.VideoId] = true
	}

	videoChan := make(chan vo.Video, len(rawVideos))

	for _, rawVideo := range rawVideos {
		video := vo.Video{
			Id:            rawVideo.Id,
			Author:        SupplementTargetUserInfo(userId, rawVideo.AuthorId),
			PlayUrl:       rawVideo.PlayUrl,
			CoverUrl:      rawVideo.CoverUrl,
			FavoriteCount: repository.CountLikedByVideoId(rawVideo.Id),
			CommentCount:  repository.CountCommentByVideoId(rawVideo.Id),
			IsFavorite:    checkIsFavorite(likeMap, rawVideo.Id),
			Title:         rawVideo.Title,
		}
		videoChan <- video
	}

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

func checkIsFavorite(likeList map[int64]bool, videoId int64) bool {
	_, ok := likeList[videoId]
	return ok
}

func SupplementCommentList(userId int64, videoId int64) ([]vo.Comment, error) {
	rawComments := repository.GetCommentList(videoId)

	commentList, err := buildComments(userId, videoId, rawComments)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func buildComments(userId int64, videoId int64, rawComments []Comment) ([]vo.Comment, error) {
	var wg sync.WaitGroup
	wg.Add(len(rawComments))

	commentChan := make(chan vo.Comment, len(rawComments))

	for _, rawComment := range rawComments {
		go func(rawComment Comment) {
			defer wg.Done()
			comment := vo.Comment{
				Id:         rawComment.Id,
				User:       SupplementTargetUserInfo(userId, rawComment.UserId),
				Content:    rawComment.CommentText,
				CreateDate: rawComment.CreateDate.Format("2006-01-02 15:04:05"),
			}
			commentChan <- comment
		}(rawComment)
	}
	wg.Wait()
	close(commentChan)

	var comments []vo.Comment
	for comment := range commentChan {
		comments = append(comments, comment)
	}
	return comments, nil
}

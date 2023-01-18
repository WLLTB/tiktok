package service

import (
	"sync"
	"tiktok/app/repository"
	. "tiktok/app/schema"
	"tiktok/app/vo"
)

func SupplementCommentList(userId int64, videoId int64) ([]vo.Comment, error) {
	rawComments := repository.GetCommentList(videoId)

	commentList, err := buildComments(userId, rawComments)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func buildComments(userId int64, rawComments []Comment) ([]vo.Comment, error) {
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

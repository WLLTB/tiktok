package service

import (
	. "tiktok/app/repository"
	. "tiktok/app/schema"
	"tiktok/app/vo"
)

func SupplementCommentList(userId int64, videoId int64) ([]vo.Comment, error) {
	rawComments := GetCommentList(videoId)

	commentList, err := buildCommentList(userId, rawComments)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func buildCommentList(userId int64, rawComments []Comment) ([]vo.Comment, error) {
	commentList := make([]vo.Comment, len(rawComments))
	for i, rawComment := range rawComments {
		comment := vo.Comment{
			Id:         rawComment.Id,
			User:       SupplementTargetUserInfo(userId, rawComment.UserId),
			Content:    rawComment.CommentText,
			CreateDate: rawComment.CreateDate.Format("2006-01-02 15:04:05"),
		}
		commentList[i] = comment
	}
	return commentList, nil
}

func HandlerCommentAction(actionType int64, comment Comment, userId int64, authorId int64) vo.Comment {
	switch actionType {
	case 1:
		commentId := InsertComment(comment)
		return vo.Comment{
			Id:         commentId,
			User:       SupplementTargetUserInfo(userId, authorId),
			Content:    comment.CommentText,
			CreateDate: comment.CreateDate.Format("2006-01-02 15:04:05"),
		}
	case 2:
		DeleteComment(comment)
		return vo.Comment{}
	default:
		return vo.Comment{}
	}
}

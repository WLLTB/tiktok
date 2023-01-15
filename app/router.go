package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	controller "tiktok/app/controller"
	"tiktok/app/utils"
	"tiktok/app/vo"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	authRouter := r.Group("/douyin")
	authRouter.Use(tokenAuth())
	unAuthRouter := r.Group("/douyin")

	// 基础接口
	unAuthRouter.GET("/feed/", controller.Feed)
	authRouter.GET("/user/", controller.UserInfo)
	unAuthRouter.POST("/user/register/", controller.Register)
	unAuthRouter.POST("/user/login/", controller.Login)
	authRouter.POST("/publish/action/", controller.Publish)
	authRouter.GET("/publish/list/", controller.PublishList)

	// 互动接口
	authRouter.POST("/favorite/action/", controller.FavoriteAction)
	authRouter.GET("/favorite/list/", controller.FavoriteList)
	authRouter.POST("/comment/action/", controller.CommentAction)
	authRouter.GET("/comment/list/", controller.CommentList)

	// 社交接口
	authRouter.POST("/relation/action/", controller.RelationAction)
	authRouter.GET("/relation/follow/list/", controller.FollowList)
	authRouter.GET("/relation/follower/list/", controller.FollowerList)
	authRouter.GET("/relation/friend/list/", controller.FriendList)
	authRouter.GET("/message/chat/", controller.MessageChat)
	authRouter.POST("/message/action/", controller.MessageAction)

	// 样例
	authRouter.GET("/demo", controller.GetTableUserList)

	r.Run(":9999")
}

func tokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if c.Request.Method == http.MethodGet {
			token = c.Query("token")
		} else {
			token = c.PostForm("token")
		}

		_, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}

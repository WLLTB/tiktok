package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	. "tiktok/app/constant"
	"tiktok/app/controller"
	"tiktok/app/utils"
	. "tiktok/app/vo"
)

func InitRouter(r *gin.Engine) {
	r.Static(StaticPath, PublicPath)
	setupCommonRoutes(r.Group(DefaultRouter))
	setupAuthRoutes(r.Group(DefaultRouter))

	r.Run(PORT)
}

func setupCommonRoutes(r *gin.RouterGroup) {
	r.Use(recoveryMiddleware())
	setUpCors(r)

	// 基础接口
	r.GET(FeedPath, controller.Feed)
	r.POST(RegisterPath, controller.Register)
	r.POST(LoginPath, controller.Login)
}

func setupAuthRoutes(r *gin.RouterGroup) {
	r.Use(tokenAuth())
	r.Use(recoveryMiddleware())
	setUpCors(r)

	// 基础接口
	r.GET(UserInfoPath, controller.UserInfo)
	r.POST(PublishPath, controller.Publish)
	r.GET(PublishListPath, controller.PublishList)

	// 互动接口
	r.POST(FavoriteActionPath, controller.FavoriteAction)
	r.GET(FavoriteListPath, controller.FavoriteList)
	r.POST(CommentActionPath, controller.CommentAction)
	r.GET(CommentListPath, controller.CommentList)

	// 社交接口
	r.POST(RelationActionPath, controller.RelationAction)
	r.GET(FollowListPath, controller.FollowList)
	r.GET(FollowerListPath, controller.FollowerList)
	r.GET(FriendListPath, controller.FriendList)
	r.GET(MessageChatPath, controller.MessageChat)
	r.POST(MessageActionPath, controller.MessageAction)

	// 样例
	r.GET(DemoPath, controller.GetTableUserList)
}

func tokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token = c.Query(TOKEN)
		if token == "" {
			token = c.PostForm(TOKEN)
		}

		_, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: InvalidMessage})
			c.Abort()
			return
		}
		c.Next()
	}
}

func recoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusOK, Response{
					StatusCode: 1,
					StatusMsg:  ServerError,
				})
			}
		}()
		c.Next()
	}
}

func setUpCors(r *gin.RouterGroup) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowMethods("PUT", "PATCH", "GET", "POST", "DELETE")
	r.Use(cors.New(config))

}

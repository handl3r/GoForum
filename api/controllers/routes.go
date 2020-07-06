package controllers

import (
	"github.com/handl3r/GoForum/api/middlewares"
)

func (server *Server) initializeRoutes() {
	v1 := server.Router.Group("/api/v1")
	{
		v1.GET("/users", server.GetUsers)
		v1.GET("/users/:id", server.GetUser)
		v1.POST("/users", server.CreateUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), server.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), server.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), server.DeleteUser)

		v1.GET("/posts", server.GetPosts)
		v1.GET("/posts/:id", server.GetPost)
		v1.GET("user_posts/:id", server.GetPostsOfUser)
		v1.POST("/posts", middlewares.TokenAuthMiddleware(), server.CreatePost)
		v1.PUT("/posts/:id", middlewares.TokenAuthMiddleware(), server.UpdatePost)
		v1.DELETE("/posts/:id", middlewares.TokenAuthMiddleware(), server.DeletePost)

		v1.GET("/likes/:id", server.GetLikes)
		v1.POST("/likes/:id",middlewares.TokenAuthMiddleware(), server.LikePost)
		v1.DELETE("/likes/:id",middlewares.TokenAuthMiddleware(), server.UnlikePost)

		v1.GET("/comments/:id", server.GetComments)
		v1.POST("/comments/:id", middlewares.TokenAuthMiddleware(), server.CreatePost)
		v1.PUT("/comments/:id", middlewares.TokenAuthMiddleware(), server.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.TokenAuthMiddleware(),server.DeleteComment)

		v1.POST("/login", server.Login)

		v1.POST("/password/forgot", server.ForgotPassword)
		v1.POST("/password/reset", server.ResetPassword)
	}
}

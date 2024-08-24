package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zedann/realtime_chat_app/server/internal/user"
	"github.com/zedann/realtime_chat_app/server/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createroom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)

}

func Start(addr string) error {
	return r.Run(addr)
}

package api

import (
	"log"

	"github.com/ari1021/websocket/controller"
	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/server/db"
	"github.com/ari1021/websocket/server/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func NewEcho(hub *websocket.Hub) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "./view/rooms.html")
	e.File("/rooms/create", "./view/create_room.html")
	e.File("/chat", "./view/chat.html")

	e.GET("/ws/:id", controller.ServeRoomWs)
	e.Validator = &customValidator{Validator: validator.New()}
	e.GET("/rooms", controller.GetRooms)
	rh := controller.NewRoomHandler(db.DB, &model.Room{})
	e.POST("/rooms", rh.CreateRoom)
	e.DELETE("/rooms/:id", controller.DeleteRoom)
	e.GET("/rooms/:id/chats", controller.GetChats)
	e.POST("/rooms/:id/chats", controller.CreateChat)
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.DB.Conn = conn
	return e
}

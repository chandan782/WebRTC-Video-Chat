package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(ctx *fiber.Ctx) error {
	return ctx.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func Room(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	if uuid == "" {
		ctx.Status(400)
		return nil
	}
	return nil
}

func RoomWebsocket(ctx *websocket.Conn) {
	uuid := ctx.Params("uuid")
	if uuid == "" {
		return
	}

}

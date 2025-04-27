package users

import (
	"github.com/bachhieu/fountain/baselib/v_net/v_api"
	"github.com/bachhieu/fountain/proto/v_proto"
	"github.com/gofiber/fiber/v2"
)

func (api *UsersAPI) IsLogin(ctx *fiber.Ctx) error {
	userID := v_api.GetContextDataString(ctx, v_api.KContextKeyUserID)
	if userID == "" {
		return v_api.WriteError(ctx, v_proto.NewRpcError(403, "You need login or register before!"))
	}
	return ctx.Next()
}

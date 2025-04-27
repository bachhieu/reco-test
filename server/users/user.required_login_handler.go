package users

import (
	"fmt"

	"github.com/bachhieu/fountain/baselib/auth_token"
	"github.com/bachhieu/fountain/baselib/v_net/v_api"
	"github.com/bachhieu/fountain/proto/v_proto"
	"github.com/gofiber/fiber/v2"
)

func (api *UsersAPI) RequiredLogin(ctx *fiber.Ctx) error {
	_, _, accountID, _, _, _, _, err := auth_token.ExtractToken(ctx, false)

	if err != nil {
		fmt.Println("le::::", err)
		return v_api.WriteError(ctx, v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), "Token invalid or expired!"))
	}

	ctx.Locals(v_api.KContextKeyUserID, accountID)
	return ctx.Next()
}

/* !!
 * Code generated by fkit
 * If you need to write additional functions, should create a separate file.
 * File Created: Sunday, 27 April 2025 16:10:50 +07:00
 * Author: VietLD (leduyviet2612@gmail.com)
 * -----
 * Last Modified: Sunday, 27 April 2025 16:10:50 +07:00
 * Modified By: HIEUBV\hieubv
 * -----
 * Copyright 2025. All rights reserved.
 */

package products

import (
	"github.com/bachhieu/fountain/baselib/base"
	"github.com/bachhieu/fountain/baselib/v_log"
	"github.com/bachhieu/fountain/baselib/v_net/v_api"
	"github.com/bachhieu/fountain/proto/v_proto"
	"github.com/bachhieu/test/biz/dal/do/product_do"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Create Product
// @Description  Create a new Product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string   true  "Insert your access token" default(access_token)
// @Param        request       body      product_do.CreateProductReq  true  "Create Product request"
// @Success      200          {object}  v_api.Response[product_do.ProductResponse]
// @Failure      400          {object}  v_proto.VolioRpcError
// @Failure      500          {object}  v_proto.VolioRpcError
// @Router       /admin/products [post]
func (api *ProductsAPI) Create(ctx *fiber.Ctx) error {
	input := &product_do.CreateProductReq{}

	if err := ctx.BodyParser(input); err != nil {
		err := v_proto.NewRpcError(int32(v_proto.VolioRpcErrorCodes_BAD_REQUEST), err.Error())
		v_log.V(1).WithError(err).Errorf("ProductsAPI::Create - Error: %+v", err)
		return v_api.WriteError(ctx, err)
	}

	input.CreatedBy = v_api.GetContextDataString(ctx, v_api.KContextKeyUserID)
	if err := api.productsController.Create(input); err != nil {
		v_log.V(1).WithError(err).Errorf("ProductsAPI::Create - Error: %+v", err)
		return v_api.WriteError(ctx, err)
	}

	v_log.V(3).Infof("ProductsAPI::Create - Reply: %s", base.JSONDebugDataString(input))

	return v_api.WriteSuccess(ctx, input)
}

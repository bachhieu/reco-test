/* !!
 * Code generated by fkit
 * If you need to write additional functions, should create a separate file.
 * File Created: Sunday, 27 April 2025 14:28:41 +07:00
 * Author: VietLD (leduyviet2612@gmail.com)
 * -----
 * Last Modified: Sunday, 27 April 2025 14:28:41 +07:00
 * Modified By: HIEUBV\hieubv
 * -----
 * Copyright 2025. All rights reserved.
 */

package users

import (
	"github.com/bachhieu/fountain/baselib/v_net/v_api"
	"github.com/bachhieu/fountain/biz/f_core"
	"github.com/bachhieu/test/biz/core/users"
)

type UsersAPI struct {
	*v_api.VolioAPI
	usersController *users.UsersController
}

func NewUsersAPI(controllers []f_core.CoreController) *UsersAPI {
	impl := &UsersAPI{VolioAPI: v_api.GetVolioAPIInstance()}

	for _, ctrl := range controllers {
		switch x := ctrl.(type) {
		case *users.UsersController:
			impl.usersController = x
		}
	}

	return impl
}

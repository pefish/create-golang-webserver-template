package controller

import (
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
)

type UserControllerClass struct {
}

var UserController = UserControllerClass{}

type LoginParams struct {
	Username string `json:"username" validate:"required" desc:"用户名"`
	Password string `json:"password" validate:"required" desc:"密码"`
}

func (lc *UserControllerClass) Login(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	params := LoginParams{}
	apiSession.MustScanParams(&params)

	return params, nil
}

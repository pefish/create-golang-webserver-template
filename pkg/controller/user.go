package controller

import (
	i_core "github.com/pefish/go-interface/i-core"
	t_error "github.com/pefish/go-interface/t-error"
)

type UserControllerType struct {
}

var UserController = UserControllerType{}

type LoginParams struct {
	Username string `json:"username" validate:"required" desc:"用户名"`
	Password string `json:"password" validate:"required" desc:"密码"`
}

func (lc *UserControllerType) Login(apiSession i_core.IApiSession) (interface{}, *t_error.ErrorInfo) {
	var params LoginParams
	err := apiSession.ScanParams(&params)
	if err != nil {
		apiSession.Logger().ErrorF("Read params error. %+v", err)
		return nil, t_error.INTERNAL_ERROR
	}

	return params, nil
}

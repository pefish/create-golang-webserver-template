package controller

import (
	_type "github.com/pefish/go-core/api-session/type"
	go_error "github.com/pefish/go-error"
)

type LoginControllerClass struct {
}

var LoginController = LoginControllerClass{}


type LoginParam struct {
	Username string `json:"username" validate:"required" desc:"用户名"`
}
func (this *LoginControllerClass) Login(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	params := LoginParam{}
	apiSession.ScanParams(&params)



	return params, nil
}
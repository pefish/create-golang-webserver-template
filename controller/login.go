package controller

import (
	"github.com/pefish/go-core/api-session"
)

type LoginControllerClass struct {
}

var LoginController = LoginControllerClass{}


type LoginParam struct {
	Username string `json:"username" validate:"required" desc:"用户名"`
}
func (this *LoginControllerClass) Login(apiSession *api_session.ApiSessionClass) interface{} {
	params := LoginParam{}
	apiSession.ScanParams(&params)



	return params
}
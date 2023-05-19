package route

import (
	"_template_/pkg/controller"
	"github.com/pefish/go-core/api"
	"github.com/pefish/go-http/gorequest"
)

var LoginRoute = []*api.Api{
	{
		Description: "登录",
		Path:        "/v1/login",
		Method:      gorequest.POST,
		Params: controller.LoginParam{
			Username: `test`,
		},
		Controller: controller.LoginController.Login,
	},
}

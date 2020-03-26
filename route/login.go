package route

import (
	"_template_/controller"
	"github.com/pefish/go-core/api"
)

var LoginRoute = []*api.Api{
	{
		Description: "登录",
		Path:        "/v1/login",
		Method:      "POST",
		Params: controller.LoginParam{
			Username: `test`,
		},
		Controller: controller.LoginController.Login,
	},
}

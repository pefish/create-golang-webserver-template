package route

import (
	"github.com/pefish/go-core/api"
	"github.com/pefish/go-http/gorequest"
	"package-name/pkg/controller"
)

var UserRoute = []*api.Api{
	{
		Description: "登录",
		Path:        "/v1/login",
		Method:      gorequest.POST,
		Params: controller.LoginParams{
			Username: `test`,
			Password: `test`,
		},
		Controller: controller.UserController.Login,
	},
}

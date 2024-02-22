package route

import (
	"package-name/pkg/controller"

	"github.com/pefish/go-core/api"
	"github.com/pefish/go-http/gorequest"
)

var UserRoute = []*api.Api{
	api.NewApi(&api.NewApiParamsType{
		Description: "登录",
		Path:        "/v1/login",
		Method:      gorequest.POST,
		Params: controller.LoginParams{
			Username: `test`,
			Password: `test`,
		},
		ControllerFunc: controller.UserController.Login,
	}),
}

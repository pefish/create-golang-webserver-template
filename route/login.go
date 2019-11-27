package route

import (
	"github.com/pefish/go-core/api-channel-builder"
	"github.com/pefish/go-core/api-strategy"
	"template/controller"
)

var LoginRoute = map[string]*api_channel_builder.Route{
	`login post`: {
		Description: "登录",
		Path:        "/v1/login",
		Method:      "POST",
		Strategies: []api_channel_builder.StrategyRoute{

		},
		ParamType: api_strategy.ALL_TYPE,
		Params: controller.LoginParam{
			Username: `test`,
		},
		Controller: controller.LoginController.Login,
	},
}

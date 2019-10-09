package service

import (
	"github.com/pefish/go-core/api-channel-builder"
	"github.com/pefish/go-core/api-strategy"
	"github.com/pefish/go-core/service"
	"oauth-login-consent/constant"
	"oauth-login-consent/controller"
)

type LoginSvcClass struct {
	service.BaseServiceClass
}

var LoginSvc = LoginSvcClass{}

func (this *LoginSvcClass) Init(opts ...interface{}) service.InterfaceService {
	this.SetName(`storm钱包服务api`)
	this.SetPath(`/api/storm-wallet`)
	api_strategy.ParamValidateStrategy.SetErrorCode(constant.PARAM_ERROR)

	this.SetRoutes(map[string]*api_channel_builder.Route{
		`get_new_deposit_address`: {
			Description: "获取新充值地址",
			Path:        "/v1/new-address",
			Method:      "POST",
			Strategies: []api_channel_builder.StrategyRoute{

			},
			ParamType: api_strategy.ALL_TYPE,
			Params: controller.LoginParam{
				Username: `test`,
			},
			Controller: controller.LoginController.Login,
		},
	})
	return this
}

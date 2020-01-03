package main

import (
	"_template_/constant"
	"_template_/route"
	"fmt"
	"github.com/pefish/go-application"
	"github.com/pefish/go-config"
	"github.com/pefish/go-core/api-strategy"
	"github.com/pefish/go-core/logger"
	"github.com/pefish/go-core/service"
	"github.com/pefish/go-logger"
	"github.com/pefish/go-mysql"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			os.Exit(1)
		}
		os.Exit(0)
	}()

	go_config.Config.MustLoadYamlConfig(go_config.Configuration{
		ConfigEnvName: `GO_CONFIG`,
		SecretEnvName: `GO_SECRET`,
	})

	service.Service.SetName(`_template_`)
	// 处理日志
	env := go_config.Config.GetString(`env`)
	go_application.Application.Debug = env == `local` || env == `dev`
	go_logger.Logger.Init(service.Service.GetName(), ``)
	logger.LoggerDriver.Register(go_logger.Logger)

	// 初始化数据库连接
	mysqlConfig := go_config.Config.MustGetMap(`mysql`)
	go_mysql.MysqlHelper.ConnectWithMap(mysqlConfig)
	defer go_mysql.MysqlHelper.Close()

	service.Service.SetHost(go_config.Config.GetString(`host`))
	service.Service.SetPort(go_config.Config.GetUint64(`port`))
	service.Service.SetPath(`/api/oauth`)
	api_strategy.ParamValidateStrategy.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.LoginRoute)
	service.Service.Run()
}

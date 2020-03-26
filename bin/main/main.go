package main

import (
	"_template_/constant"
	"_template_/route"
	"fmt"
	"github.com/pefish/go-application"
	"github.com/pefish/go-config"
	go_core "github.com/pefish/go-core"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
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

	go_core.Service.SetName(`_template_`)
	// 处理日志
	go_application.Application.SetEnv(go_config.Config.GetString(`env`))
	go_logger.Logger = go_logger.NewLogger(go_logger.WithName(go_core.Service.GetName()), go_logger.WithIsDebug(go_application.Application.Debug))
	logger.LoggerDriver.Register(go_logger.Logger)

	// 初始化数据库连接
	mysqlConfig := go_config.Config.MustGetMap(`mysql`)
	go_mysql.MysqlHelper.SetLogger(go_logger.Logger)
	go_mysql.MysqlHelper.MustConnectWithMap(mysqlConfig)
	defer go_mysql.MysqlHelper.Close()

	go_core.Service.SetHost(go_config.Config.GetString(`host`))
	go_core.Service.SetPort(go_config.Config.GetUint64(`port`))
	go_core.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategy.SetErrorCode(constant.PARAM_ERROR)

	go_core.Service.SetRoutes(route.LoginRoute)
	go_core.Service.Run()
}

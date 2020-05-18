package main

import (
	"_template_/constant"
	"_template_/route"
	"github.com/pefish/go-application"
	"github.com/pefish/go-config"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
	"github.com/pefish/go-core/service"
	"github.com/pefish/go-logger"
	"github.com/pefish/go-mysql"
	"github.com/pefish/go-task-driver"
	"log"
)

func main() {
	go_config.Config.MustLoadYamlConfig(go_config.Configuration{
		ConfigEnvName: `GO_CONFIG`,
		SecretEnvName: `GO_SECRET`,
	})

	service.Service.SetName(`_template_`)
	// 处理日志
	go_application.Application.SetEnv(go_config.Config.MustGetString(`env`))
	go_logger.Logger = go_logger.NewLogger(go_logger.WithIsDebug(go_application.Application.Debug))
	logger.LoggerDriver.Register(go_logger.Logger)

	// 初始化数据库连接
	mysqlConfig, err := go_config.Config.GetMap(`mysql`)
	if err != nil {
		log.Fatal(err)
	}
	go_mysql.MysqlHelper.SetLogger(go_logger.Logger)
	go_mysql.MysqlHelper.MustConnectWithMap(mysqlConfig)
	defer go_mysql.MysqlHelper.Close()

	service.Service.SetHost(go_config.Config.MustGetString(`host`))
	service.Service.SetPort(go_config.Config.MustGetUint64Default(`port`, 8000))
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategy.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.LoginRoute)

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.SetLogger(go_logger.Logger)
	taskDriver.Register(service.Service)

	taskDriver.RunWait()
}

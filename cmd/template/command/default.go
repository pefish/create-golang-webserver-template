package command

import (
	"_template_/pkg/constant"
	"_template_/pkg/route"
	"_template_/version"
	"flag"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
	"github.com/pefish/go-core/service"
	go_logger "github.com/pefish/go-logger"
	task_driver "github.com/pefish/go-task-driver"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	service.Service.SetName(version.AppName)
	// 处理日志
	logger.LoggerDriverInstance.Register(go_logger.Logger)

	// 初始化数据库连接
	//mysqlConfig, err := go_config.ConfigManagerInstance.GetMap(`mysql`)
	//if err != nil {
	//	return err
	//}
	//go_mysql.MysqlInstance.SetLogger(go_logger.Logger)
	//go_mysql.MysqlInstance.MustConnectWithMap(mysqlConfig)
	//defer go_mysql.MysqlInstance.Close()

	service.Service.SetHost(go_config.ConfigManagerInstance.MustGetString(`host`))
	service.Service.SetPort(go_config.ConfigManagerInstance.MustGetUint64Default(`port`, 8000))
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategyInstance.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.LoginRoute)

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.SetLogger(go_logger.Logger)
	taskDriver.Register(service.Service)

	taskDriver.RunWait(data.ExitCancelCtx)

	return nil
}

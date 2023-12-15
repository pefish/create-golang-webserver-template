package command

import (
	"_package-name_/pkg/constant"
	"_package-name_/pkg/global"
	"_package-name_/pkg/route"
	"_package-name_/version"
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
	flagSet.String("host", "127.0.0.1", "The host of web server.")
	flagSet.Int("port", 8000, "The port of web server.")
	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	//go_mysql.MysqlInstance.Close()
	return nil
}

func (dc *DefaultCommand) Init(data *commander.StartData) error {
	service.Service.SetName(version.AppName)
	logger.LoggerDriverInstance.Register(go_logger.Logger)

	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	//go_mysql.MysqlInstance.SetLogger(go_logger.Logger)
	//go_mysql.MysqlInstance.MustConnectWithConfiguration(go_mysql.Configuration{
	//	Host:            global.GlobalConfig.Db.Host,
	//	Username:        global.GlobalConfig.Db.User,
	//	Password:        global.GlobalConfig.Db.Pass,
	//	Database:        global.GlobalConfig.Db.Db,
	//})

	service.Service.SetHost(global.GlobalConfig.Host)
	service.Service.SetPort(global.GlobalConfig.Port)
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategyInstance.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.UserRoute)

	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(service.Service)

	taskDriver.RunWait(data.ExitCancelCtx)

	return nil
}

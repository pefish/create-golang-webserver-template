package command

import (
	"package-name/pkg/constant"
	"package-name/pkg/global"
	"package-name/pkg/route"
	"package-name/version"

	"github.com/pefish/go-commander"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
	"github.com/pefish/go-core/service"
	task_driver "github.com/pefish/go-task-driver"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *DefaultCommand) Data() interface{} {
	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	//go_mysql.MysqlInstance.Close()
	return nil
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	service.Service.SetName(version.AppName)
	logger.LoggerDriverInstance.Register(command.Logger)

	// go_mysql.MysqlInstance.SetLogger(command.Logger)
	// err = go_mysql.MysqlInstance.ConnectWithConfiguration(go_mysql.Configuration{
	// 	Host:     global.GlobalConfig.DbHost,
	// 	Port:     global.GlobalConfig.DbPort,
	// 	Username: global.GlobalConfig.DbUser,
	// 	Password: global.GlobalConfig.DbPass,
	// 	Database: global.GlobalConfig.DbDb,
	// })
	// if err != nil {
	// 	return err
	// }

	service.Service.SetHost(global.GlobalConfig.ServerHost)
	service.Service.SetPort(uint64(global.GlobalConfig.ServerPort))
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategyInstance.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.NewUserRoute())

	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(service.Service)

	taskDriver.RunWait(command.Ctx)

	return nil
}

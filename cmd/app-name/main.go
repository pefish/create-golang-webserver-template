package main

import (
	"fmt"
	"package-name/cmd/app-name/command"
	"package-name/version"

	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(
		version.AppName,
		version.Version,
		fmt.Sprintf("%s is a template. Author: pefish", version.AppName),
	)
	commanderInstance.RegisterDefaultSubcommand(&commander.SubcommandInfo{
		Desc:       "Use this command by default if you don't set subcommand.",
		Args:       nil,
		Subcommand: command.NewDefaultCommand(),
	})
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}

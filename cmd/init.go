package cmd

import (
	"PortalClient/configs"
	"PortalClient/pkg/tracing"
)

func Init() error {
	if err := configs.InitConfigFile(); err != nil {
		return err
	}
	if err := configs.InitMySql(); err != nil {
		return err
	}
	if err := tracing.InitTracer(); err != nil {
		return err
	}

	return nil
}

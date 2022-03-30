package engine

import (
	"cli/internal/clients"
	"cli/internal/config"
)

//go:generate mockgen -source=./interface.go -destination=./interface_mock.go -package=engine

type Interface interface {
	//Run will get the template from server
	Run()
	//Init will init any environment
	Init() error
	Save()
}

type Engine struct {
	Config config.Config
	//ETDClient is used for any admin connection
	ETDClient clients.ETDInterface
	//ZipClient is used for unzip operation
	ZipClient clients.ZipInterface
	//InstallClient is used for installation local dependencies
	InstallClient clients.InstallInterface
}

package engine

import (
	"cli/internal/clients"
	"cli/internal/config"
)

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
	ETDClient clients.ETD
	//ZipClient is used for unzip operation
	ZipClient clients.Zip
	//InstallClient is used for installation local dependencies
	InstallClient clients.Install
}

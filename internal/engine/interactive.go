package engine

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"

	"cli/internal/constants"
)

type InteractiveEngine struct {
	Engine
}

//NewInteractiveEngine will create a new interactive command line
func NewInteractiveEngine() InteractiveEngine {
	return InteractiveEngine{}
}

//Init initialize selection client
func (e *InteractiveEngine) Init() {
	// read saved Config from file
	viper.SetConfigName(constants.ConfigName)
	viper.AddConfigPath(constants.ConfigUnixGlobalPath)
	viper.AddConfigPath(constants.ConfigUnixHomePath)
	viper.AddConfigPath(constants.ConfigLocalPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err == nil {
		fmt.Printf("%s", err)
		// Found configuration
		e.Config.DefaultNodeId = viper.GetString(constants.NodeIDKey)
		e.Config.DefaultEndpoint = viper.GetString(constants.EndpointKey)
		e.Config.DefaultPassword = viper.GetString(constants.PasswordKey)
	}
}

func (e *InteractiveEngine) Run() {

}

//Save will save configurations to a Config file
func (e *InteractiveEngine) Save() {
	viper.Set(constants.NodeIDKey, e.Config.DefaultNodeId)
	viper.Set(constants.EndpointKey, e.Config.DefaultEndpoint)
	viper.Set(constants.PasswordKey, e.Config.DefaultPassword)

	err := viper.WriteConfigAs(path.Join(constants.ConfigLocalPath, constants.ConfigName))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

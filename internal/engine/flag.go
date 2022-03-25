package engine

import (
	"flag"
	"os"

	"github.com/joho/godotenv"

	"cli/internal/constants"
	"cli/internal/errors"
)

type FlagEngine struct {
	Engine
}

func NewFlagEngine() FlagEngine {
	return FlagEngine{}
}

func (e *FlagEngine) Init() error {
	environment := flag.String("environment", "production", "Which environment you are running?")
	template := flag.String("template", "", "Which template id you want to use?")

	flag.Parse()

	if *environment != constants.Production && *environment != constants.Beta && *environment != constants.Local {
		return errors.NewInvalidEnvironmentError(*environment)
	}

	if len(*template) == 0 {
		return errors.NewInvalidTemplateIdError(*template)
	}

	if err := e.readEnvironmentFile(); err != nil {
		return err
	}
	e.Config.DefaultTemplateId = *template
	e.setupEndpoint(*environment)
	return nil
}

func (e *FlagEngine) Run() error {
	e.Config.Print()
	if err := e.ETDClient.VerifyPassword(); err != nil {
		return err
	}
	e.ETDClient.GetTemplate(e.Config.DefaultTemplateId)
	return nil
}

func (e *FlagEngine) Save() {

}

//readEnvironmentFile will try to read environment variables from dotenv file
func (e *FlagEngine) readEnvironmentFile() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	e.Config.DefaultPassword = os.Getenv(constants.PasswordKey)
	e.Config.DefaultNodeId = os.Getenv(constants.NodeIDKey)
	if len(e.Config.DefaultPassword) == 0 {
		return errors.NewInvalidPasswordError(e.Config.DefaultPassword)
	}

	if len(e.Config.DefaultNodeId) == 0 {
		return errors.NewInvalidNodeIdError(e.Config.DefaultNodeId)
	}

	e.ETDClient.Password = e.Config.DefaultPassword
	e.ETDClient.NodeId = e.Config.DefaultNodeId
	return nil
}

//setupEndpoint will update default endpoint base on the given environment
func (e *FlagEngine) setupEndpoint(environment string) {
	if environment == constants.Production {
		e.Config.DefaultEndpoint = constants.ProductionEndpoint
		e.ETDClient.Url = constants.ProductionEndpoint
	}

	if environment == constants.Beta {
		e.Config.DefaultEndpoint = constants.BetaEndpoint
		e.ETDClient.Url = constants.BetaEndpoint
	}

	if environment == constants.Local {
		e.Config.DefaultEndpoint = constants.LocalEndpoint
		e.ETDClient.Url = constants.LocalEndpoint
	}
}

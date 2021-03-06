package engine

import (
	goerror "errors"
	"flag"
	"io/fs"
	"os"

	"github.com/joho/godotenv"

	"cli/internal/clients"
	"cli/internal/constants"
	"cli/internal/errors"
)

type FlagEngine struct {
	Engine
}

func NewFlagEngine(etdClient clients.ETDInterface, zipClient clients.ZipInterface) FlagEngine {
	return FlagEngine{
		Engine{
			ETDClient: etdClient,
			ZipClient: zipClient,
		},
	}
}

func (e *FlagEngine) Init() error {
	password := flag.String("password", "", "Remote admin password")
	environment := flag.String("environment", "production", "Which environment you are running?")
	template := flag.String("template", "", "Which template id you want to use?")

	flag.Parse()

	if *environment != constants.Production && *environment != constants.Beta && *environment != constants.Local {
		return errors.NewInvalidEnvironmentError(*environment)
	}

	if len(*template) == 0 {
		return errors.NewInvalidTemplateIdError(*template)
	}

	if len(*password) == 0 {
		return errors.NewInvalidPasswordError(*password)
	}

	if err := e.readEnvironmentFile(); err != nil {
		return err
	}
	e.Config.DefaultTemplateId = *template
	e.Config.DefaultPassword = *password
	e.ETDClient.SetPassword(*password)
	e.setupEndpoint(*environment)
	return nil
}

func (e *FlagEngine) Run() error {
	e.Config.Print()
	if err := e.ETDClient.VerifyPassword(); err != nil {
		return err
	}

	if err := e.ETDClient.GetTemplate(e.Config.DefaultTemplateId); err != nil {
		return err
	}

	if err := e.ZipClient.UnZip(constants.SavedTemplateFileName); err != nil {
		return err
	}

	if err := e.ZipClient.Remove(constants.SavedTemplateFileName); err != nil {
		return err
	}

	return nil
}

func (e *FlagEngine) Save() {

}

//readEnvironmentFile will try to read environment variables from dotenv file
func (e *FlagEngine) readEnvironmentFile() error {
	err := godotenv.Load()
	if err != nil && !goerror.Is(err, fs.ErrNotExist) {
		return err
	}
	e.Config.DefaultNodeId = os.Getenv(constants.NodeIDKey)

	if len(e.Config.DefaultNodeId) == 0 {
		return errors.NewInvalidNodeIdError(e.Config.DefaultNodeId)
	}

	e.ETDClient.SetNodeId(e.Config.DefaultNodeId)
	return nil
}

//setupEndpoint will update default endpoint base on the given environment
func (e *FlagEngine) setupEndpoint(environment string) {
	if environment == constants.Production {
		e.Config.DefaultEndpoint = constants.ProductionEndpoint
		e.ETDClient.SetURL(constants.ProductionEndpoint)
	}

	if environment == constants.Beta {
		e.Config.DefaultEndpoint = constants.BetaEndpoint
		e.ETDClient.SetURL(constants.BetaEndpoint)
	}

	if environment == constants.Local {
		e.Config.DefaultEndpoint = constants.LocalEndpoint
		e.ETDClient.SetURL(constants.LocalEndpoint)
	}
}

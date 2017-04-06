package client

import (
	"errors"
	"github.com/3dsim/simulation-goclient/auth0"
	"github.com/3dsim/simulation-goclient/genclient"
	"github.com/3dsim/simulation-goclient/genclient/operations"
	"github.com/3dsim/simulation-goclient/models"
	"github.com/3dsim/simulation-goclient/sugar"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/inconshreveable/log15"
	"net/url"
	"time"
)

// Log is a github.com/inconshreveable/log15.Logger.  Log is exposed so that users of this library can set
// their own log handler.  By default this Log uses the DiscardHandler, which discards log statements.
// See: https://godoc.org/github.com/inconshreveable/log15#hdr-Library_Use
//
// To set a different log handler do something like this:
//
// 		Log.SetHandler(log.LvlFilterHandler(log.LvlInfo, log.CallerFileHandler(log.StdoutHandler)))
var Log = log.New()

func init() {
	Log.SetHandler(log.DiscardHandler())
}

const (
	// SimulationAPIBasePath is the base path or "slug" for the simulation api
	SimulationAPIBasePath = "simulation-api"
)

// Client is a wrapper around the generated client found in the "genclient" package.  It provides convenience methods
// for common operations.  If the operation needed is not found in Client, use the "genclient" package using this client
// as an example of how to utilize the genclient.  PRs are welcome if more functionality is wanted in this client package.
type Client interface {
	Simulations(organizationID int32, status []string, sort []string, limit, offset int32) ([]*models.Simulation, error)
	StartSimulation(simulationID int32) error
	ThermalSimulation(simulationID int32) (*models.ThermalSimulation, error)
	SingleBeadSimulation(simulationID int32) (*models.SingleBeadSimulation, error)
	PostAssumedStrainSimulation(*models.AssumedStrainSimulation) (*models.AssumedStrainSimulation, error)
	PostPorositySimulation(*models.PorositySimulation) (*models.PorositySimulation, error)
	PostScanPatternSimulation(*models.ScanPatternSimulation) (*models.ScanPatternSimulation, error)
	PostSingleBeadSimulation(*models.SingleBeadSimulation) (*models.SingleBeadSimulation, error)
	PostThermalSimulation(*models.ThermalSimulation) (*models.ThermalSimulation, error)
	Machine(machineID int32) (*models.Machine, error)
	Material(materialID int32) (*models.Material, error)
	Simulation(simulationID int32) (*models.Simulation, error)
	ScanPatternSimulation(simulationID int32) (*models.ScanPatternSimulation, error)
	AssumedStrainSimulation(simulationID int32) (*models.AssumedStrainSimulation, error)
	PorositySimulation(simulationID int32) (*models.PorositySimulation, error)
	AddLogWithTime(simulationID int, messageDate time.Time, message string) error
	AddLog(simulationID int, message string) error
}

type client struct {
	tokenFetcher auth0.TokenFetcher
	client       *genclient.Simulation
	audience     string
}

// New creates a new client for interacting with the 3DSIM simulation api.  See the auth0 package for how to construct
// the token fetcher.  The apiGatewayURL's are as follows:
//
// 		QA 				= https://3dsim-qa.cloud.tyk.io
//		Prod and Gov 	= https://3dsim.cloud.tyk.io
//
// The audience's are:
//
// 		QA 		= https://simulation-qa.3dsim.com/v2
//		Prod 	= https://simulation.3dsim.com/v2
// 		Gov 	= https://simulation-gov.3dsim.com
func New(tokenFetcher auth0.TokenFetcher, apiGatewayURL, audience string) Client {
	parsedURL, err := url.Parse(apiGatewayURL)
	if err != nil {
		message := "API Gateway URL was invalid!"
		Log.Error(message, "apiGatewayURL", apiGatewayURL)
		panic(message + " " + err.Error())
	}
	simulationTransport := openapiclient.New(parsedURL.Host, SimulationAPIBasePath, []string{parsedURL.Scheme})
	simulationTransport.Debug = true
	simulationClient := genclient.New(simulationTransport, strfmt.Default)
	return &client{
		tokenFetcher: tokenFetcher,
		client:       simulationClient,
		audience:     audience,
	}
}

func (c *client) Machine(machineID int32) (*models.Machine, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetMachine(operations.NewGetMachineParams().WithID(machineID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) Material(materialID int32) (*models.Material, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetMaterial(operations.NewGetMaterialParams().WithID(materialID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) Simulation(simulationID int32) (*models.Simulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetSimulation(operations.NewGetSimulationParams().WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) ThermalSimulation(simulationID int32) (*models.ThermalSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetThermalSimulation(operations.NewGetThermalSimulationParams().WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) ScanPatternSimulation(simulationID int32) (*models.ScanPatternSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetScanPatternSimulation(operations.NewGetScanPatternSimulationParams().WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) AssumedStrainSimulation(simulationID int32) (*models.AssumedStrainSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetAssumedStrainSimulation(operations.NewGetAssumedStrainSimulationParams().WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PorositySimulation(simulationID int32) (*models.PorositySimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetPorositySimulation(operations.NewGetPorositySimulationParams().WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) SingleBeadSimulation(simulationID int32) (*models.SingleBeadSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Operations.GetSingleBeadSimulation(operations.NewGetSingleBeadSimulationParams().WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) AddLogWithTime(simulationID int, messageDate time.Time, message string) error {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return err
	}
	loggedAt := strfmt.DateTime(messageDate)
	simulationLog := models.SimulationLog{LoggedAt: &loggedAt, Message: &message, SimulationID: int32(simulationID)}
	_, err = c.client.Operations.PostSimulationLog(operations.NewPostSimulationLogParams().WithID(int32(simulationID)).
		WithSimulationLog(&simulationLog), openapiclient.BearerToken(token))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) AddLog(simulationID int, message string) error {
	return c.AddLogWithTime(simulationID, time.Now().UTC(), message)
}

func (c *client) Simulations(organizationID int32, status []string, sort []string, limit, offset int32) ([]*models.Simulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetSimulationsParams().
		WithOrganizationID(organizationID).
		WithStatus(status).
		WithSort(sort).
		WithLimit(sugar.Int32(limit)).
		WithOffset(sugar.Int32(offset))

	response, err := c.client.Operations.GetSimulations(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) StartSimulation(simulationID int32) error {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return err
	}
	bearerToken := openapiclient.BearerToken(token)

	simulation, err := c.Simulation(simulationID)
	if err != nil {
		return err
	}

	switch simulation.Type {
	case models.SimulationTypeAssumedStrainSimulation:
		params := operations.NewStartAssumedStrainSimulationParams().WithID(simulationID)
		_, err = c.client.Operations.StartAssumedStrainSimulation(params, bearerToken)
		if err != nil {
			return err
		}
	case models.SimulationTypePorositySimulation:
		params := operations.NewStartPorositySimulationParams().WithID(simulationID)
		_, err = c.client.Operations.StartPorositySimulation(params, bearerToken)
		if err != nil {
			return err
		}
	case models.SimulationTypeScanPatternSimulation:
		params := operations.NewStartScanPatternSimulationParams().WithID(simulationID)
		_, err = c.client.Operations.StartScanPatternSimulation(params, bearerToken)
		if err != nil {
			return err
		}
	case models.SimulationTypeSingleBeadSimulation:
		params := operations.NewStartSingleBeadSimulationParams().WithID(simulationID)
		_, err = c.client.Operations.StartSingleBeadSimulation(params, bearerToken)
		if err != nil {
			return err
		}
	case models.SimulationTypeThermalSimulation:
		params := operations.NewStartThermalSimulationParams().WithID(simulationID)
		_, err = c.client.Operations.StartThermalSimulation(params, bearerToken)
		if err != nil {
			return err
		}
	default:
		return errors.New("Unrecognized simulation type")
	}

	return nil
}

func (c *client) PostAssumedStrainSimulation(simulation *models.AssumedStrainSimulation) (*models.AssumedStrainSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	// clear simulationParts part id and simulation id so server will create
	// new simulation part entries
	for _, part := range simulation.AssumedStrainSimulationParameters.SimulationParts {
		part.ID = 0
		part.SimulationID = sugar.Int32(0)
	}
	params := operations.NewPostAssumedStrainSimulationParams().WithAssumedStrainSimulation(simulation)
	response, err := c.client.Operations.PostAssumedStrainSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostPorositySimulation(simulation *models.PorositySimulation) (*models.PorositySimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewPostPorositySimulationParams().WithPorositySimulation(simulation)
	response, err := c.client.Operations.PostPorositySimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostScanPatternSimulation(simulation *models.ScanPatternSimulation) (*models.ScanPatternSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	// clear simulationParts part id and simulation id so server will create
	// new simulation part entries
	for _, part := range simulation.ScanPatternSimulationParameters.SimulationParts {
		part.ID = 0
		part.SimulationID = sugar.Int32(0)
	}
	params := operations.NewPostScanPatternSimulationParams().WithScanPatternSimulation(simulation)
	response, err := c.client.Operations.PostScanPatternSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostSingleBeadSimulation(simulation *models.SingleBeadSimulation) (*models.SingleBeadSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewPostSingleBeadSimulationParams().WithSingleBeadSimulation(simulation)
	response, err := c.client.Operations.PostSingleBeadSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostThermalSimulation(simulation *models.ThermalSimulation) (*models.ThermalSimulation, error) {
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	// clear simulationParts part id and simulation id so server will create
	// new simulation part entries
	for _, part := range simulation.ThermalSimulationParameters.SimulationParts {
		part.ID = 0
		part.SimulationID = sugar.Int32(0)
	}
	params := operations.NewPostThermalSimulationParams().WithThermalSimulation(simulation)
	response, err := c.client.Operations.PostThermalSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

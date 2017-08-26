//go:generate counterfeiter ./ Client

package simulation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/3dsim/auth0"
	"github.com/3dsim/simulation-goclient/genclient"
	"github.com/3dsim/simulation-goclient/genclient/operations"
	"github.com/3dsim/simulation-goclient/models"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	log "github.com/inconshreveable/log15"
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
	Simulations(organizationID int32, status []string, sort []string, offset, limit int32) ([]*models.Simulation, error)
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
	PostLog(level string, message string, simulationID int32, activityID string) (err error)
	PostLogWithObject(simulationLog models.SimulationLog) (err error)
	PostLogInfo(message string, simulationID int32, activityID string) (err error)
	PostLogWarn(message string, simulationID int32, activityID string) (err error)
	PostLogError(message string, simulationID int32, activityID string) (err error)
	PatchSimulation(simulationID int32, patch *models.PatchDocument) error
	MultiPatchSimulation(simulationID int32, patches []*models.PatchDocument) error
	PostSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) (*models.SimulationActivity, error)
	SimulationActivityByActivityID(simulationID int32, activityID string) (*models.SimulationActivity, error)
	PutSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) error
	AddSimulationOutput(simulationID int32, outputType, outputFileLocation string) (*models.SimulationOutput, error)
	UpdateSimulationStatus(simulationID int32, status string) error
	// RawSimulation gets a simulation as a map instead of a struct
	RawSimulation(simulationID int32) (map[string]interface{}, error)
	BuildFiles(organizationID int32, availability []string, sort []string, offset, limit int32) ([]*models.BuildFile, error)
	BuildFile(buildFileID int32) (*models.BuildFile, error)
	PostBuildFile(*models.BuildFilePost) (*models.BuildFile, error)
	PatchBuildFile(buildFileID int32, patches []*models.PatchDocument) (*models.BuildFile, error)
	UpdateBuildFileAvailability(buildFileID int32, availability string) (*models.BuildFile, error)
	RawBuildFile(buildFileID int32) (map[string]interface{}, error)
	Part(partID int32) (*models.Part, error)
	UpdatePartAvailability(partID int32, availability string) (*models.Part, error)
	PatchPart(partID int32, patches []*models.PatchDocument) (*models.Part, error)
}

type client struct {
	tokenFetcher auth0.TokenFetcher
	client       *genclient.Simulation
	audience     string
}

// NewClient creates a client for interacting with the 3DSIM simulation api.  See the auth0 package for how to construct
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
func NewClient(tokenFetcher auth0.TokenFetcher, apiGatewayURL, audience string) Client {
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

func (c *client) Machine(machineID int32) (machine *models.Machine, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) Material(materialID int32) (material *models.Material, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) Simulation(simulationID int32) (simulation *models.Simulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) ThermalSimulation(simulationID int32) (thermalSimulation *models.ThermalSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) ScanPatternSimulation(simulationID int32) (sps *models.ScanPatternSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) AssumedStrainSimulation(simulationID int32) (a *models.AssumedStrainSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) PorositySimulation(simulationID int32) (p *models.PorositySimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) SingleBeadSimulation(simulationID int32) (s *models.SingleBeadSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) PostLogWithObject(simulationLog models.SimulationLog) (err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return err
	}
	_, err = c.client.Operations.PostSimulationLog(operations.NewPostSimulationLogParams().WithID(simulationLog.SimulationID).
		WithSimulationLog(&simulationLog), openapiclient.BearerToken(token))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) PostLogInfo(message string, simulationID int32, activityID string) (err error) {
	return c.PostLog(models.SimulationLogLevelInfo, message, simulationID, activityID)
}

func (c *client) PostLogWarn(message string, simulationID int32, activityID string) (err error) {
	return c.PostLog(models.SimulationLogLevelWarn, message, simulationID, activityID)
}

func (c *client) PostLogError(message string, simulationID int32, activityID string) (err error) {
	return c.PostLog(models.SimulationLogLevelError, message, simulationID, activityID)
}

func (c *client) PostLog(level string, message string, simulationID int32, activityID string) (err error) {
	loggedAt := strfmt.DateTime(time.Now().UTC())
	simulationLog := models.SimulationLog{
		Level:        level,
		Message:      &message,
		SimulationID: simulationID,
		ActivityID:   activityID,
		LoggedAt:     &loggedAt,
	}
	return c.PostLogWithObject(simulationLog)
}

func (c *client) Simulations(organizationID int32, status []string, sort []string, offset, limit int32) (simulations []*models.Simulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetSimulationsParams().
		WithOrganizationID(organizationID).
		WithStatus(status).
		WithSort(sort).
		WithLimit(swag.Int32(limit)).
		WithOffset(swag.Int32(offset))

	response, err := c.client.Operations.GetSimulations(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) StartSimulation(simulationID int32) (err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) PostAssumedStrainSimulation(simulation *models.AssumedStrainSimulation) (a *models.AssumedStrainSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	// clear simulationParts part id and simulation id so server will create
	// new simulation part entries
	for _, part := range simulation.AssumedStrainSimulationParameters.SimulationParts {
		part.ID = 0
		part.SimulationID = swag.Int32(0)
	}
	params := operations.NewPostAssumedStrainSimulationParams().WithAssumedStrainSimulation(simulation)
	response, err := c.client.Operations.PostAssumedStrainSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostPorositySimulation(simulation *models.PorositySimulation) (p *models.PorositySimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

// AddSimulationOutput adds the given output type and output file location to the
// simulations_outputs table, iff the output type and output file location don't
// exist already for this simulation.
func (c *client) AddSimulationOutput(simulationID int32, outputType, outputFileLocation string) (so *models.SimulationOutput, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	simulationOutputsResponse, err := c.client.Operations.GetSimulationOutputs(operations.NewGetSimulationOutputsParams().WithID(simulationID),
		openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	for _, simulationOutput := range simulationOutputsResponse.Payload {
		if *simulationOutput.FileLocation == outputFileLocation {
			// output already added, so no-op
			return simulationOutput, nil
		}
	}
	createdAt := strfmt.DateTime(time.Now().UTC())
	simulationOutput := &models.SimulationOutput{
		Label:        swag.String(outputType),
		FileLocation: swag.String(outputFileLocation),
		SimulationID: simulationID,
		CreatedAt:    &createdAt,
	}
	response, err := c.client.Operations.PostSimulationOutput(operations.NewPostSimulationOutputParams().WithID(simulationID).
		WithSimulationOutput(simulationOutput), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostScanPatternSimulation(simulation *models.ScanPatternSimulation) (s *models.ScanPatternSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	// clear simulationParts part id and simulation id so server will create
	// new simulation part entries
	for _, part := range simulation.ScanPatternSimulationParameters.SimulationParts {
		part.ID = 0
		part.SimulationID = swag.Int32(0)
	}
	params := operations.NewPostScanPatternSimulationParams().WithScanPatternSimulation(simulation)
	response, err := c.client.Operations.PostScanPatternSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostSingleBeadSimulation(simulation *models.SingleBeadSimulation) (s *models.SingleBeadSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
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

func (c *client) PostSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) (s *models.SimulationActivity, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewPostSimulationActivityParams().WithSimulationActivity(simulationActivity).WithID(simulationID)
	response, err := c.client.Operations.PostSimulationActivity(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) SimulationActivityByActivityID(simulationID int32, activityID string) (sa *models.SimulationActivity, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	activitiesResponse, err := c.client.Operations.GetSimulationActivities(operations.NewGetSimulationActivitiesParams().
		WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	for _, activity := range activitiesResponse.Payload {
		if *activity.ActivityID == activityID {
			return activity, nil
		}
	}
	return nil, fmt.Errorf("a SimulationActivity with activity id %v not found for simulation %v", activityID, simulationID)
}

func (c *client) PutSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) (err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return err
	}
	params := operations.NewPutSimulationActivityParams().WithSimulationActivity(simulationActivity).WithID(simulationID)
	_, err = c.client.Operations.PutSimulationActivity(params, openapiclient.BearerToken(token))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) PostThermalSimulation(simulation *models.ThermalSimulation) (t *models.ThermalSimulation, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	// clear simulationParts part id and simulation id so server will create
	// new simulation part entries
	for _, part := range simulation.ThermalSimulationParameters.SimulationParts {
		part.ID = 0
		part.SimulationID = swag.Int32(0)
	}
	params := operations.NewPostThermalSimulationParams().WithThermalSimulation(simulation)
	response, err := c.client.Operations.PostThermalSimulation(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

// PatchSimulation is only available to internal clients that have full access to the API
func (c *client) PatchSimulation(simulationID int32, patch *models.PatchDocument) (err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	return c.MultiPatchSimulation(simulationID, []*models.PatchDocument{patch})
}

// MultiPatchSimulation is only available to internal clients that have full access to the API
func (c *client) MultiPatchSimulation(simulationID int32, patches []*models.PatchDocument) (err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return err
	}
	_, err = c.client.Operations.PatchSimulation(operations.NewPatchSimulationParams().
		WithSimulationPatch(patches).WithID(simulationID), openapiclient.BearerToken(token))
	if err != nil {
		return err
	}
	return nil
}

func (c *client) UpdateSimulationStatus(simulationID int32, status string) (err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	patch := &models.PatchDocument{Op: swag.String(models.PatchDocumentOpReplace), Path: swag.String("/status"), Value: status}
	return c.PatchSimulation(simulationID, patch)
}

func (c *client) RawSimulation(simulationID int32) (s map[string]interface{}, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	bearerToken := openapiclient.BearerToken(token)

	simulation, err := c.Simulation(simulationID)
	if err != nil {
		return nil, err
	}

	var rawSimulation interface{}
	switch simulation.Type {
	case models.SimulationTypeAssumedStrainSimulation:
		params := operations.NewGetAssumedStrainSimulationParams().WithID(simulationID)
		payload, err := c.client.Operations.GetAssumedStrainSimulation(params, bearerToken)
		if err != nil {
			return nil, err
		}
		rawSimulation = *payload.Payload
	case models.SimulationTypePorositySimulation:
		params := operations.NewGetPorositySimulationParams().WithID(simulationID)
		payload, err := c.client.Operations.GetPorositySimulation(params, bearerToken)
		if err != nil {
			return nil, err
		}
		rawSimulation = *payload.Payload
	case models.SimulationTypeScanPatternSimulation:
		params := operations.NewGetScanPatternSimulationParams().WithID(simulationID)
		payload, err := c.client.Operations.GetScanPatternSimulation(params, bearerToken)
		if err != nil {
			return nil, err
		}
		rawSimulation = *payload.Payload
	case models.SimulationTypeSingleBeadSimulation:
		params := operations.NewGetSingleBeadSimulationParams().WithID(simulationID)
		payload, err := c.client.Operations.GetSingleBeadSimulation(params, bearerToken)
		if err != nil {
			return nil, err
		}
		rawSimulation = *payload.Payload
	case models.SimulationTypeThermalSimulation:
		params := operations.NewGetThermalSimulationParams().WithID(simulationID)
		payload, err := c.client.Operations.GetThermalSimulation(params, bearerToken)
		if err != nil {
			return nil, err
		}
		rawSimulation = *payload.Payload
	default:
		return nil, errors.New("Unrecognized simulation type")
	}
	rawBytes, err := json.Marshal(rawSimulation)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawBytes, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) BuildFiles(organizationID int32, availability []string, sort []string, offset, limit int32) (buildFiles []*models.BuildFile, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetBuildFilesParams().
		WithOrganizationID(organizationID).
		WithAvailability(availability).
		WithSort(sort).
		WithLimit(swag.Int32(limit)).
		WithOffset(swag.Int32(offset))

	response, err := c.client.Operations.GetBuildFiles(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) BuildFile(buildFileID int32) (buildFile *models.BuildFile, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetBuildFileParams().WithID(buildFileID)

	response, err := c.client.Operations.GetBuildFile(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PostBuildFile(buildFilePost *models.BuildFilePost) (buildFile *models.BuildFile, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}

	params := operations.NewPostBuildFileParams().WithBuildFilePost(buildFilePost)

	response, err := c.client.Operations.PostBuildFile(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PatchBuildFile(buildFileID int32, patches []*models.PatchDocument) (buildFile *models.BuildFile, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}

	params := operations.NewPatchBuildFileParams().WithBuildFilePatch(patches).WithID(buildFileID)

	response, err := c.client.Operations.PatchBuildFile(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) UpdateBuildFileAvailability(buildFileID int32, availability string) (buildFile *models.BuildFile, err error) {

	patch := models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	return c.PatchBuildFile(buildFileID, []*models.PatchDocument{&patch})
}

func (c *client) RawBuildFile(buildFileID int32) (s map[string]interface{}, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	bearerToken := openapiclient.BearerToken(token)

	params := operations.NewGetBuildFileParams().WithID(buildFileID)
	payload, err := c.client.Operations.GetBuildFile(params, bearerToken)
	if err != nil {
		return nil, err
	}
	rawBytes, err := json.Marshal(*payload.Payload)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawBytes, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) Part(partID int32) (part *models.Part, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}
	params := operations.NewGetPartParams().WithID(partID)

	response, err := c.client.Operations.GetPart(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) PatchPart(partID int32, patches []*models.PatchDocument) (part *models.Part, err error) {
	defer func() {
		// Until this issue is resolved: https://github.com/go-swagger/go-swagger/issues/1021, we need to recover from
		// panics.
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()

	token, err := c.tokenFetcher.Token(c.audience)
	if err != nil {
		return nil, err
	}

	params := operations.NewPatchPartParams().WithPartPatch(patches).WithID(partID)

	response, err := c.client.Operations.PatchPart(params, openapiclient.BearerToken(token))
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) UpdatePartAvailability(partID int32, availability string) (part *models.Part, err error) {

	patch := models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	return c.PatchPart(partID, []*models.PatchDocument{&patch})
}

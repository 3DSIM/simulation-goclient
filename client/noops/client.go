package noops

import "github.com/3dsim/simulation-goclient/models"
import "time"

// Client is a noop implementation of the simulation-goclient/client/Client interface
type Client struct {
}

// Simulations provides a noop function with given fields: organizationID, status, sort, limit, offset
func (_m *Client) Simulations(organizationID int32, status []string, sort []string, limit int32, offset int32) ([]*models.Simulation, error) {
	return nil, nil
}

// StartSimulation provides a noop function with given fields: simulationID
func (_m *Client) StartSimulation(simulationID int32) error {
	return nil
}

// ThermalSimulation provides a noop function with given fields: simulationID
func (_m *Client) ThermalSimulation(simulationID int32) (*models.ThermalSimulation, error) {
	return nil, nil
}

// SingleBeadSimulation provides a noop function with given fields: simulationID
func (_m *Client) SingleBeadSimulation(simulationID int32) (*models.SingleBeadSimulation, error) {
	return nil, nil
}

// PostAssumedStrainSimulation provides a noop function with given fields: _a0
func (_m *Client) PostAssumedStrainSimulation(_a0 *models.AssumedStrainSimulation) (*models.AssumedStrainSimulation, error) {
	return nil, nil
}

// PostPorositySimulation provides a noop function with given fields: _a0
func (_m *Client) PostPorositySimulation(_a0 *models.PorositySimulation) (*models.PorositySimulation, error) {
	return nil, nil
}

// PostScanPatternSimulation provides a noop function with given fields: _a0
func (_m *Client) PostScanPatternSimulation(_a0 *models.ScanPatternSimulation) (*models.ScanPatternSimulation, error) {
	return nil, nil
}

// PostSingleBeadSimulation provides a noop function with given fields: _a0
func (_m *Client) PostSingleBeadSimulation(_a0 *models.SingleBeadSimulation) (*models.SingleBeadSimulation, error) {
	return nil, nil
}

// PostThermalSimulation provides a noop function with given fields: _a0
func (_m *Client) PostThermalSimulation(_a0 *models.ThermalSimulation) (*models.ThermalSimulation, error) {
	return nil, nil
}

// Machine provides a noop function with given fields: machineID
func (_m *Client) Machine(machineID int32) (*models.Machine, error) {
	return nil, nil
}

// Material provides a noop function with given fields: materialID
func (_m *Client) Material(materialID int32) (*models.Material, error) {
	return nil, nil
}

// Simulation provides a noop function with given fields: simulationID
func (_m *Client) Simulation(simulationID int32) (*models.Simulation, error) {
	return nil, nil
}

// ScanPatternSimulation provides a noop function with given fields: simulationID
func (_m *Client) ScanPatternSimulation(simulationID int32) (*models.ScanPatternSimulation, error) {
	return nil, nil
}

// AssumedStrainSimulation provides a noop function with given fields: simulationID
func (_m *Client) AssumedStrainSimulation(simulationID int32) (*models.AssumedStrainSimulation, error) {
	return nil, nil
}

// PorositySimulation provides a noop function with given fields: simulationID
func (_m *Client) PorositySimulation(simulationID int32) (*models.PorositySimulation, error) {
	return nil, nil
}

// PostLogWithTime provides a noop function with given fields: simulationID, messageDate, message
func (_m *Client) PostLogWithTime(simulationID int32, messageDate time.Time, message string) error {
	return nil
}

// PostLog provides a noop function with given fields: simulationID, message
func (_m *Client) PostLog(simulationID int32, message string) error {
	return nil
}

// PatchSimulation provides a noop function with given fields: simulationID, patch
func (_m *Client) PatchSimulation(simulationID int32, patch *models.PatchDocument) error {
	return nil
}

// MultiPatchSimulation provides a noop function with given fields: simulationID, patches
func (_m *Client) MultiPatchSimulation(simulationID int32, patches []*models.PatchDocument) error {
	return nil
}

// PostSimulationActivity provides a noop function with given fields: simulationID, simulationActivity
func (_m *Client) PostSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) (*models.SimulationActivity, error) {
	return nil, nil
}

// SimulationActivityByActivityID provides a noop function with given fields: simulationID, activityID
func (_m *Client) SimulationActivityByActivityID(simulationID int32, activityID string) (*models.SimulationActivity, error) {
	return nil, nil
}

// PutSimulationActivity provides a noop function with given fields: simulationID, simulationActivity
func (_m *Client) PutSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) error {
	return nil
}

// AddSimulationOutput provides a noop function with given fields: simulationID, outputType, outputFileLocation
func (_m *Client) AddSimulationOutput(simulationID int32, outputType string, outputFileLocation string) (*models.SimulationOutput, error) {
	return nil, nil
}

// UpdateSimulationStatus provides a noop function with given fields: simulationID, status
func (_m *Client) UpdateSimulationStatus(simulationID int32, status string) error {
	return nil
}

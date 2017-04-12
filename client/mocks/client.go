package mocks

import "github.com/stretchr/testify/mock"

import "github.com/3dsim/simulation-goclient/models"

import "time"

type Client struct {
	mock.Mock
}

// Simulations provides a mock function with given fields: organizationID, status, sort, offset, limit
func (_m *Client) Simulations(organizationID int32, status []string, sort []string, offset int32, limit int32) ([]*models.Simulation, error) {
	ret := _m.Called(organizationID, status, sort, offset, limit)

	var r0 []*models.Simulation
	if rf, ok := ret.Get(0).(func(int32, []string, []string, int32, int32) []*models.Simulation); ok {
		r0 = rf(organizationID, status, sort, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Simulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, []string, []string, int32, int32) error); ok {
		r1 = rf(organizationID, status, sort, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartSimulation provides a mock function with given fields: simulationID
func (_m *Client) StartSimulation(simulationID int32) error {
	ret := _m.Called(simulationID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(simulationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ThermalSimulation provides a mock function with given fields: simulationID
func (_m *Client) ThermalSimulation(simulationID int32) (*models.ThermalSimulation, error) {
	ret := _m.Called(simulationID)

	var r0 *models.ThermalSimulation
	if rf, ok := ret.Get(0).(func(int32) *models.ThermalSimulation); ok {
		r0 = rf(simulationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ThermalSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(simulationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SingleBeadSimulation provides a mock function with given fields: simulationID
func (_m *Client) SingleBeadSimulation(simulationID int32) (*models.SingleBeadSimulation, error) {
	ret := _m.Called(simulationID)

	var r0 *models.SingleBeadSimulation
	if rf, ok := ret.Get(0).(func(int32) *models.SingleBeadSimulation); ok {
		r0 = rf(simulationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SingleBeadSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(simulationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostAssumedStrainSimulation provides a mock function with given fields: _a0
func (_m *Client) PostAssumedStrainSimulation(_a0 *models.AssumedStrainSimulation) (*models.AssumedStrainSimulation, error) {
	ret := _m.Called(_a0)

	var r0 *models.AssumedStrainSimulation
	if rf, ok := ret.Get(0).(func(*models.AssumedStrainSimulation) *models.AssumedStrainSimulation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AssumedStrainSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.AssumedStrainSimulation) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostPorositySimulation provides a mock function with given fields: _a0
func (_m *Client) PostPorositySimulation(_a0 *models.PorositySimulation) (*models.PorositySimulation, error) {
	ret := _m.Called(_a0)

	var r0 *models.PorositySimulation
	if rf, ok := ret.Get(0).(func(*models.PorositySimulation) *models.PorositySimulation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PorositySimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.PorositySimulation) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostScanPatternSimulation provides a mock function with given fields: _a0
func (_m *Client) PostScanPatternSimulation(_a0 *models.ScanPatternSimulation) (*models.ScanPatternSimulation, error) {
	ret := _m.Called(_a0)

	var r0 *models.ScanPatternSimulation
	if rf, ok := ret.Get(0).(func(*models.ScanPatternSimulation) *models.ScanPatternSimulation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ScanPatternSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ScanPatternSimulation) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostSingleBeadSimulation provides a mock function with given fields: _a0
func (_m *Client) PostSingleBeadSimulation(_a0 *models.SingleBeadSimulation) (*models.SingleBeadSimulation, error) {
	ret := _m.Called(_a0)

	var r0 *models.SingleBeadSimulation
	if rf, ok := ret.Get(0).(func(*models.SingleBeadSimulation) *models.SingleBeadSimulation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SingleBeadSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.SingleBeadSimulation) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostThermalSimulation provides a mock function with given fields: _a0
func (_m *Client) PostThermalSimulation(_a0 *models.ThermalSimulation) (*models.ThermalSimulation, error) {
	ret := _m.Called(_a0)

	var r0 *models.ThermalSimulation
	if rf, ok := ret.Get(0).(func(*models.ThermalSimulation) *models.ThermalSimulation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ThermalSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ThermalSimulation) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Machine provides a mock function with given fields: machineID
func (_m *Client) Machine(machineID int32) (*models.Machine, error) {
	ret := _m.Called(machineID)

	var r0 *models.Machine
	if rf, ok := ret.Get(0).(func(int32) *models.Machine); ok {
		r0 = rf(machineID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Machine)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(machineID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Material provides a mock function with given fields: materialID
func (_m *Client) Material(materialID int32) (*models.Material, error) {
	ret := _m.Called(materialID)

	var r0 *models.Material
	if rf, ok := ret.Get(0).(func(int32) *models.Material); ok {
		r0 = rf(materialID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Material)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(materialID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Simulation provides a mock function with given fields: simulationID
func (_m *Client) Simulation(simulationID int32) (*models.Simulation, error) {
	ret := _m.Called(simulationID)

	var r0 *models.Simulation
	if rf, ok := ret.Get(0).(func(int32) *models.Simulation); ok {
		r0 = rf(simulationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Simulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(simulationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanPatternSimulation provides a mock function with given fields: simulationID
func (_m *Client) ScanPatternSimulation(simulationID int32) (*models.ScanPatternSimulation, error) {
	ret := _m.Called(simulationID)

	var r0 *models.ScanPatternSimulation
	if rf, ok := ret.Get(0).(func(int32) *models.ScanPatternSimulation); ok {
		r0 = rf(simulationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ScanPatternSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(simulationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssumedStrainSimulation provides a mock function with given fields: simulationID
func (_m *Client) AssumedStrainSimulation(simulationID int32) (*models.AssumedStrainSimulation, error) {
	ret := _m.Called(simulationID)

	var r0 *models.AssumedStrainSimulation
	if rf, ok := ret.Get(0).(func(int32) *models.AssumedStrainSimulation); ok {
		r0 = rf(simulationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AssumedStrainSimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(simulationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PorositySimulation provides a mock function with given fields: simulationID
func (_m *Client) PorositySimulation(simulationID int32) (*models.PorositySimulation, error) {
	ret := _m.Called(simulationID)

	var r0 *models.PorositySimulation
	if rf, ok := ret.Get(0).(func(int32) *models.PorositySimulation); ok {
		r0 = rf(simulationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PorositySimulation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(simulationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostLogWithTime provides a mock function with given fields: simulationID, messageDate, message
func (_m *Client) PostLogWithTime(simulationID int32, messageDate time.Time, message string) error {
	ret := _m.Called(simulationID, messageDate, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, time.Time, string) error); ok {
		r0 = rf(simulationID, messageDate, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostLog provides a mock function with given fields: simulationID, message
func (_m *Client) PostLog(simulationID int32, message string) error {
	ret := _m.Called(simulationID, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, string) error); ok {
		r0 = rf(simulationID, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PatchSimulation provides a mock function with given fields: simulationID, patch
func (_m *Client) PatchSimulation(simulationID int32, patch *models.PatchDocument) error {
	ret := _m.Called(simulationID, patch)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, *models.PatchDocument) error); ok {
		r0 = rf(simulationID, patch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MultiPatchSimulation provides a mock function with given fields: simulationID, patches
func (_m *Client) MultiPatchSimulation(simulationID int32, patches []*models.PatchDocument) error {
	ret := _m.Called(simulationID, patches)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, []*models.PatchDocument) error); ok {
		r0 = rf(simulationID, patches)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostSimulationActivity provides a mock function with given fields: simulationID, simulationActivity
func (_m *Client) PostSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) (*models.SimulationActivity, error) {
	ret := _m.Called(simulationID, simulationActivity)

	var r0 *models.SimulationActivity
	if rf, ok := ret.Get(0).(func(int32, *models.SimulationActivity) *models.SimulationActivity); ok {
		r0 = rf(simulationID, simulationActivity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SimulationActivity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, *models.SimulationActivity) error); ok {
		r1 = rf(simulationID, simulationActivity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SimulationActivityByActivityID provides a mock function with given fields: simulationID, activityID
func (_m *Client) SimulationActivityByActivityID(simulationID int32, activityID string) (*models.SimulationActivity, error) {
	ret := _m.Called(simulationID, activityID)

	var r0 *models.SimulationActivity
	if rf, ok := ret.Get(0).(func(int32, string) *models.SimulationActivity); ok {
		r0 = rf(simulationID, activityID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SimulationActivity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, string) error); ok {
		r1 = rf(simulationID, activityID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutSimulationActivity provides a mock function with given fields: simulationID, simulationActivity
func (_m *Client) PutSimulationActivity(simulationID int32, simulationActivity *models.SimulationActivity) error {
	ret := _m.Called(simulationID, simulationActivity)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, *models.SimulationActivity) error); ok {
		r0 = rf(simulationID, simulationActivity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddSimulationOutput provides a mock function with given fields: simulationID, outputType, outputFileLocation
func (_m *Client) AddSimulationOutput(simulationID int32, outputType string, outputFileLocation string) (*models.SimulationOutput, error) {
	ret := _m.Called(simulationID, outputType, outputFileLocation)

	var r0 *models.SimulationOutput
	if rf, ok := ret.Get(0).(func(int32, string, string) *models.SimulationOutput); ok {
		r0 = rf(simulationID, outputType, outputFileLocation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SimulationOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, string, string) error); ok {
		r1 = rf(simulationID, outputType, outputFileLocation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSimulationStatus provides a mock function with given fields: simulationID, status
func (_m *Client) UpdateSimulationStatus(simulationID int32, status string) error {
	ret := _m.Called(simulationID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, string) error); ok {
		r0 = rf(simulationID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

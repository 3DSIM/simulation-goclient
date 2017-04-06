package mocks

import "github.com/stretchr/testify/mock"

import "github.com/3dsim/simulation-goclient/models"

import "time"

type Client struct {
	mock.Mock
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

// AddLogWithTime provides a mock function with given fields: simulationID, messageDate, message
func (_m *Client) AddLogWithTime(simulationID int, messageDate time.Time, message string) error {
	ret := _m.Called(simulationID, messageDate, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, time.Time, string) error); ok {
		r0 = rf(simulationID, messageDate, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddLog provides a mock function with given fields: simulationID, message
func (_m *Client) AddLog(simulationID int, message string) error {
	ret := _m.Called(simulationID, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(simulationID, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

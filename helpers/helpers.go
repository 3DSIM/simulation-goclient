package helpers

import (
	"github.com/3dsim/simulation-goclient/models"
)

// IsSimulationUsingMachine returns true if the simulation is a type that uses a machine (and therefore would
// have a populated machineID field) and false otherwise.
func IsSimulationUsingMachine(simulation *models.Simulation) bool {
	switch simulation.Type {
	case models.SimulationTypeThermalSimulation,
		models.SimulationTypeScanPatternSimulation,
		models.SimulationTypeSingleBeadSimulation,
		models.SimulationTypePorositySimulation:
		return true
	case models.SimulationTypeAssumedStrainSimulation:
		return false
	default:
		panic("Unknown simulation type: " + simulation.Type)
	}
}

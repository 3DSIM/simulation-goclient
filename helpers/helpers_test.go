package helpers

import (
	"fmt"
	"github.com/3dsim/simulation-goclient/models"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIsSimulationUsingMachine(t *testing.T) {
	// arrange
	testCases := []struct {
		name     string
		value    string
		expected bool
	}{
		{"ThermalSimulation", models.SimulationTypeThermalSimulation, true},
		{"SingleBeadSimulation", models.SimulationTypeSingleBeadSimulation, true},
		{"PorositySimulation", models.SimulationTypePorositySimulation, true},
		{"AssumedStrainSimulation", models.SimulationTypeAssumedStrainSimulation, false},
		{"ScanPatternSimulation", models.SimulationTypeScanPatternSimulation, true},
	}

	for _, tc := range testCases {
		booleanInTitleCase := strings.Title(fmt.Sprintf("%t", tc.expected))
		t.Run(fmt.Sprintf("When%sExpects%s", tc.name, booleanInTitleCase), func(t *testing.T) {
			simulation := &models.Simulation{
				Type: tc.value,
			}

			// act
			isSimulationUsingMachine := IsSimulationUsingMachine(simulation)

			// assert
			assert.Equal(t, tc.expected, isSimulationUsingMachine)
		})
	}
}

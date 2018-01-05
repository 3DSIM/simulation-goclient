package examples

import (
	"fmt"
	simulationclient "github.com/3dsim/simulation-goclient/genclient"
	"github.com/3dsim/simulation-goclient/genclient/operations"
	"github.com/3dsim/simulation-goclient/models"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func _TestExampleUseOfAPIWithAuthentication(t *testing.T) {
	token := "token received from auth0 (without Bearer in front)"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := simulationclient.New(openapiclient.New("3dsim-qa.cloud.tyk.io", "simulation-api", []string{"https"}), strfmt.Default)

	simulation := models.Simulation{
		OrganizationID: swag.Int32(1),
		Title:          swag.String("Title"),
	}
	partBasedSimulationParams := models.PartBasedSimulationParameters {
		ElasticModulus:                             swag.Float64(208e9),
		MaximumWallDistance:                        swag.Float64(3e-3),
		MaximumWallThickness:                       swag.Float64(1e-3),
		MinimumSupportHeight:                       swag.Float64(5e-3),
		MinimumWallDistance:                        swag.Float64(0),
		MinimumWallThickness:                       swag.Float64(5e-5),
		PoissonRatio:                               swag.Float64(0.33),
		StrainScalingFactor:                        swag.Float64(1.0),
		StressMode:                                 swag.String("LinearElastic"),
		SupportAngle:                               swag.Float64(45),
		SupportFactorOfSafety:                      swag.Float64(1.0),
		SupportYieldStrength:                       swag.Float64(480e6),
		SupportYieldStrengthRatio:                  swag.Float64(0.4375),
		VoxelSize:                                  swag.Float64(5e-4),
	}
	assumedStrainParams := models.AssumedStrainSimulationParameters{
		LayerThickness: 50e-6,
	}
	simulationToCreate := &models.AssumedStrainSimulation{
		Simulation: simulation,
		PartBasedSimulationParameters: partBasedSimulationParams,
		AssumedStrainSimulationParameters: assumedStrainParams,
	}
	createdSimulation, err := client.Operations.PostAssumedStrainSimulation(operations.NewPostAssumedStrainSimulationParams().WithAssumedStrainSimulation(simulationToCreate), bearerTokenAuth)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Result: %v\n", createdSimulation)
}

func _TestExampleGetSimulations(t *testing.T) {
	token := "token received from auth0 (without Bearer in front)"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := simulationclient.New(openapiclient.New("3dsim-qa.cloud.tyk.io", "simulation-api", []string{"https"}), strfmt.Default)

	simulations, err := client.Operations.GetSimulations(operations.NewGetSimulationsParams().WithStatus([]string{"InProgress", "Error", "Success"}).WithOrganizationID(1), bearerTokenAuth)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, simulations.Payload, "Expected some simulations returned")
	fmt.Printf("Result: %v\n", simulations.Payload)
}

func _TestPatch(t *testing.T) {
	token := "sample token"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := simulationclient.New(openapiclient.New("localhost:5000", "", []string{"http"}), strfmt.Default)

	patch := &models.PatchDocument{Op: swag.String(models.PatchDocumentOpReplace), Path: swag.String("/title"), Value: "Test Patch"}
	patch2 := &models.PatchDocument{Op: swag.String(models.PatchDocumentOpReplace), Path: swag.String("/status"), Value: models.SimulationStatusInProgress}
	patchList := []*models.PatchDocument{patch, patch2}
	patchedSimulation, err := client.Operations.PatchSimulation(operations.NewPatchSimulationParams().WithSimulationPatch(patchList).WithID(65), bearerTokenAuth)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Result: %v\n", patchedSimulation)
}

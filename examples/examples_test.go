package examples

import (
	"fmt"
	"testing"

	simulationclient "github.com/3dsim/simulation-goclient/client"
	"github.com/3dsim/simulation-goclient/client/operations"
	"github.com/3dsim/simulation-goclient/models"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
)

func _TestExampleUseOfAPIWithAuthentication(t *testing.T) {
	token := "token received from auth0 (without Bearer in front)"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := simulationclient.New(openapiclient.New("3dsim-qa.cloud.tyk.io", "simulation-api", []string{"https"}), strfmt.Default)

	simulation := models.Simulation{
		OrganizationID: swag.Int32(1),
		Title:          swag.String("Title"),
	}
	assumedStrainParams := models.AssumedStrainSimulationParameters{
		AnisotropicStrainCoefficientsZ:             swag.Float64(1.0),
		AnisotropicStrainCoefficientsParallel:      swag.Float64(1.5),
		AnisotropicStrainCoefficientsPerpendicular: swag.Float64(0.5),
		AssumedStrain:                              swag.Float64(1.0),
		ElasticModulus:                             swag.Float64(208e9),
		HatchSpacing:                               swag.Float64(100e-6),
		LaserWattage:                               swag.Float64(195),
		LayerThickness:                             swag.Float64(50e-6),
		LayerRotationAngle:                         swag.Float64(67),
		MaximumWallDistance:                        swag.Float64(3e-3),
		MaximumWallThickness:                       swag.Float64(1e-3),
		MinimumSupportHeight:                       swag.Float64(5e-3),
		MinimumWallDistance:                        swag.Float64(0),
		MinimumWallThickness:                       swag.Float64(5e-5),
		OutputShrinkage:                            swag.Bool(true),
		PoissonRatio:                               swag.Float64(0.33),
		StrainScalingFactor:                        swag.Float64(1.0),
		ScanSpeed:                                  swag.Float64(1.0),
		SlicingStripeWidth:                         swag.Float64(2e-3),
		StartingLayerAngle:                         swag.Float64(57),
		StressMode:                                 swag.String("LinearElastic"),
		SupportAngle:                               swag.Float64(45),
		SupportFactorOfSafety:                      swag.Float64(1.0),
		SupportOptimization:                        swag.Bool(true),
		SupportYieldStrength:                       swag.Float64(480e6),
		SupportYieldStrengthRatio:                  swag.Float64(0.4375),
		UsePeriodicAnalysis:                        swag.Bool(false),
		VoxelSize:                                  swag.Float64(5e-4),
	}
	simulationToCreate := &models.AssumedStrainSimulation{
		Simulation:                        simulation,
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

package examples

import (
	"fmt"
	simulationclient "github.com/3dsim/simulation-goclient/client"
	"github.com/3dsim/simulation-goclient/client/operations"
	"github.com/3dsim/simulation-goclient/models"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"testing"
)

func _TestExampleUseOfAPIWithAuthentication(t *testing.T) {
	token := "token received from auth0 (without Bearer in front)"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := simulationclient.New(openapiclient.New("3dsim-qa.cloud.tyk.io", "simulation-api", []string{"https"}), strfmt.Default)

	simulation := models.Simulation{
		OrganizationID: int64ToPointer(1),
		Title:          stringToPointer("Title"),
	}
	assumedStrainParams := models.AssumedStrainSimulationParameters{
		AnisotropicStrainCoefficientsZ:             float64ToPointer(1.0),
		AnisotropicStrainCoefficientsParallel:      float64ToPointer(1.5),
		AnisotropicStrainCoefficientsPerpendicular: float64ToPointer(0.5),
		AssumedStrain:                              float64ToPointer(1.0),
		ElasticModulus:                             float64ToPointer(208e9),
		HatchSpacing:                               float64ToPointer(100e-6),
		LaserWattage:                               float64ToPointer(195),
		LayerThickness:                             float64ToPointer(50e-6),
		LayerRotationAngle:                         float64ToPointer(67),
		MaximumWallDistance:                        float64ToPointer(3e-3),
		MaximumWallThickness:                       float64ToPointer(1e-3),
		MinimumSupportHeight:                       float64ToPointer(5e-3),
		MinimumWallDistance:                        float64ToPointer(0),
		MinimumWallThickness:                       float64ToPointer(5e-5),
		OutputShrinkage:                            boolToPointer(true),
		OutputStateMap:                             boolToPointer(true),
		PoissonRatio:                               float64ToPointer(0.33),
		RelaxationFactor:                           float64ToPointer(1.0),
		ScanSpeed:                                  float64ToPointer(1.0),
		SlicingStripeWidth:                         float64ToPointer(2e-3),
		StartingLayerAngle:                         float64ToPointer(57),
		StressMode:                                 stringToPointer("LinearElastic"),
		SupportAngle:                               float64ToPointer(45),
		SupportFactorOfSafety:                      float64ToPointer(1.0),
		SupportOptimization:                        boolToPointer(true),
		SupportYieldStrength:                       float64ToPointer(480e6),
		SupportYieldStrengthRatio:                  float64ToPointer(0.4375),
		UsePeriodicAnalysis:                        boolToPointer(false),
		VoxelSize:                                  float64ToPointer(5e-4),
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

func TestPatch(t *testing.T) {
	token := "sample token"
	bearerTokenAuth := openapiclient.BearerToken(token)

	client := simulationclient.New(openapiclient.New("localhost:5000", "", []string{"http"}), strfmt.Default)

	patch := &models.PatchDocument{Op: stringToPointer(models.PatchDocumentOpReplace), Path: stringToPointer("/title"), Value: "Test Patch"}
	patch2 := &models.PatchDocument{Op: stringToPointer(models.PatchDocumentOpReplace), Path: stringToPointer("/status"), Value: models.SimulationStatusInProgress}
	patchList := []*models.PatchDocument{patch, patch2}
	patchedSimulation, err := client.Operations.PatchSimulation(operations.NewPatchSimulationParams().WithSimulationPatch(patchList).WithID(65), bearerTokenAuth)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Result: %v\n", patchedSimulation)
}

func stringToPointer(text string) *string {
	return &text
}

func boolToPointer(b bool) *bool {
	return &b
}

func int64ToPointer(i int64) *int64 {
	return &i
}

func float64ToPointer(f float64) *float64 {
	return &f
}

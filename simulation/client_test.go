package simulation

import (
	"errors"
	"io/ioutil"
	"testing"
	"time"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/3dsim/auth0/auth0fakes"
	"github.com/3dsim/simulation-goclient/models"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	audience = "test audience"
)

func TestMachineWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	machineID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", "", audience)

	// act
	_, err := client.Machine(machineID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestMachineWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	machineID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Machine
	machineHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	machineEndpoint := "/" + SimulationAPIBasePath + "/machines/{machineID}"
	r.HandleFunc(machineEndpoint, machineHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.Machine(machineID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestMachineWhenSuccessfulExpectsMachineReturned(t *testing.T) {
	// arrange
	machineID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Machine
	machineToReturn := &models.Machine{
		ID:   swag.Int32(machineID),
		Name: swag.String("Machine name"),
	}
	machineHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedMachineID, err := strconv.Atoi(mux.Vars(r)["machineID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(machineID), receivedMachineID, "Expected machine id received to match what was passed in")
		bytes, err := json.Marshal(machineToReturn)
		if err != nil {
			t.Error("Failed to marshal machine")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	machineEndpoint := "/" + SimulationAPIBasePath + "/machines/{machineID}"
	r.HandleFunc(machineEndpoint, machineHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	machine, err := client.Machine(machineID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *machineToReturn.Name, *machine.Name, "Expected names to match")
	assert.Equal(t, *machineToReturn.ID, *machine.ID, "Expected IDs to match")
}

// Materials
func TestMaterialWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	materialID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "", "", audience)

	// act
	_, err := client.Material(materialID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestMaterialWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	materialID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Material
	materialHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	materialEndpoint := "/" + SimulationAPIBasePath + "/materials/{materialID}"
	r.HandleFunc(materialEndpoint, materialHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.Material(materialID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestMaterialWhenSuccessfulExpectsMaterialReturned(t *testing.T) {
	// arrange
	materialID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Material
	materialToReturn := &models.Material{
		ID:   swag.Int32(materialID),
		Name: swag.String("Material name"),
	}
	materialHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedMaterialID, err := strconv.Atoi(mux.Vars(r)["materialID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(materialID), receivedMaterialID, "Expected material id received to match what was passed in")
		bytes, err := json.Marshal(materialToReturn)
		if err != nil {
			t.Error("Failed to marshal material")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	materialEndpoint := "/" + SimulationAPIBasePath + "/materials/{materialID}"
	r.HandleFunc(materialEndpoint, materialHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	material, err := client.Material(materialID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *materialToReturn.Name, *material.Name, "Expected names to match")
	assert.Equal(t, *materialToReturn.ID, *material.ID, "Expected IDs to match")
}

// Simulations
func TestSimulationWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	simulationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	client := NewClient(fakeTokenFetcher, "", "", audience)

	// act
	_, err := client.Simulation(simulationID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestSimulationWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	simulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Simulation
	simulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	simulationEndpoint := "/" + SimulationAPIBasePath + "/simulations/{simulationID}"
	r.HandleFunc(simulationEndpoint, simulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.Simulation(simulationID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestSimulationWhenSuccessfulExpectsSimulationReturned(t *testing.T) {
	// arrange
	simulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Simulation
	simulationToReturn := &models.Simulation{
		ID:    simulationID,
		Title: swag.String("Simulation name"),
	}
	simulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedSimulationID, err := strconv.Atoi(mux.Vars(r)["simulationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationID), receivedSimulationID, "Expected simulation id received to match what was passed in")
		bytes, err := json.Marshal(simulationToReturn)
		if err != nil {
			t.Error("Failed to marshal simulation")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	simulationEndpoint := "/" + SimulationAPIBasePath + "/simulations/{simulationID}"
	r.HandleFunc(simulationEndpoint, simulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	simulation, err := client.Simulation(simulationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *simulationToReturn.Title, *simulation.Title, "Expected names to match")
	assert.Equal(t, simulationToReturn.ID, simulation.ID, "Expected IDs to match")
}

// ThermalSimulations
func TestThermalSimulationWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	thermalSimulationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	thermalSimulationService := NewClient(fakeTokenFetcher, "", "", audience)

	// act
	_, err := thermalSimulationService.ThermalSimulation(thermalSimulationID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestThermalSimulationWhenThermalSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	thermalSimulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// ThermalSimulation
	thermalSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	thermalSimulationEndpoint := "/" + SimulationAPIBasePath + "/thermalsimulations/{thermalSimulationID}"
	r.HandleFunc(thermalSimulationEndpoint, thermalSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	thermalSimulationService := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := thermalSimulationService.ThermalSimulation(thermalSimulationID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because thermalSimulation api send a 500 error")
}

func TestThermalSimulationWhenSuccessfulExpectsThermalSimulationReturned(t *testing.T) {
	// arrange
	thermalSimulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// ThermalSimulation
	thermalSimulationToReturn := &models.ThermalSimulation{
		Simulation: models.Simulation{
			ID:    thermalSimulationID,
			Title: swag.String("ThermalSimulation name"),
		},
		PartBasedSimulationParameters: models.PartBasedSimulationParameters{
			ElasticModulus: swag.Float64(10),
		},
		ThermalSimulationParameters: models.ThermalSimulationParameters{
			AnisotropicStrainCoefficientsParallel:      swag.Float64(21),
			AnisotropicStrainCoefficientsPerpendicular: swag.Float64(22),
			AnisotropicStrainCoefficientsZ:             swag.Float64(23),
		},
	}
	thermalSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedThermalSimulationID, err := strconv.Atoi(mux.Vars(r)["thermalSimulationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(thermalSimulationID), receivedThermalSimulationID, "Expected thermalSimulation id received to match what was passed in")
		bytes, err := json.Marshal(thermalSimulationToReturn)
		if err != nil {
			t.Error("Failed to marshal thermalSimulation")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	thermalSimulationEndpoint := "/" + SimulationAPIBasePath + "/thermalsimulations/{thermalSimulationID}"
	r.HandleFunc(thermalSimulationEndpoint, thermalSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	thermalSimulationService := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	thermalSimulation, err := thermalSimulationService.ThermalSimulation(thermalSimulationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *thermalSimulationToReturn.Title, *thermalSimulation.Title, "Expected names to match")
	assert.Equal(t, thermalSimulationToReturn.ID, thermalSimulation.ID, "Expected IDs to match")
	assert.EqualValues(t, *thermalSimulationToReturn.ElasticModulus, *thermalSimulation.ElasticModulus, "Expected ElasticModulus values to match")
	assert.EqualValues(t,
		*thermalSimulationToReturn.AnisotropicStrainCoefficientsParallel,
		*thermalSimulation.AnisotropicStrainCoefficientsParallel,
		"Expected AnisotropicStrainCoefficientsParallel values to match")
}

// ScanPatternSimulations
func TestScanPatternSimulationWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	scanPatternSimulationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	scanPatternSimulationService := NewClient(fakeTokenFetcher, "", "", audience)

	// act
	_, err := scanPatternSimulationService.ScanPatternSimulation(scanPatternSimulationID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestScanPatternSimulationWhenScanPatternSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	scanPatternSimulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// ScanPatternSimulation
	scanPatternSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	scanPatternSimulationEndpoint := "/" + SimulationAPIBasePath + "/scanpatternsimulations/{scanPatternSimulationID}"
	r.HandleFunc(scanPatternSimulationEndpoint, scanPatternSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	scanPatternSimulationService := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := scanPatternSimulationService.ScanPatternSimulation(scanPatternSimulationID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because scanPatternSimulation api send a 500 error")
}

func TestScanPatternSimulationWhenSuccessfulExpectsScanPatternSimulationReturned(t *testing.T) {
	// arrange
	scanPatternSimulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// ScanPatternSimulation
	scanPatternSimulationToReturn := &models.ScanPatternSimulation{
		Simulation: models.Simulation{
			ID:    scanPatternSimulationID,
			Title: swag.String("ScanPatternSimulation name"),
		},
		PartBasedSimulationParameters: models.PartBasedSimulationParameters{
			ElasticModulus: swag.Float64(10),
		},
		ScanPatternSimulationParameters: models.ScanPatternSimulationParameters{
			LayerThickness: swag.Float64(21),
		},
	}
	scanPatternSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedScanPatternSimulationID, err := strconv.Atoi(mux.Vars(r)["scanPatternSimulationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(scanPatternSimulationID), receivedScanPatternSimulationID, "Expected scanPatternSimulation id received to match what was passed in")
		bytes, err := json.Marshal(scanPatternSimulationToReturn)
		if err != nil {
			t.Error("Failed to marshal scanPatternSimulation")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	scanPatternSimulationEndpoint := "/" + SimulationAPIBasePath + "/scanpatternsimulations/{scanPatternSimulationID}"
	r.HandleFunc(scanPatternSimulationEndpoint, scanPatternSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	scanPatternSimulationService := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	scanPatternSimulation, err := scanPatternSimulationService.ScanPatternSimulation(scanPatternSimulationID)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *scanPatternSimulationToReturn.Title, *scanPatternSimulation.Title, "Expected names to match")
	assert.Equal(t, scanPatternSimulationToReturn.ID, scanPatternSimulation.ID, "Expected IDs to match")
	assert.EqualValues(t, *scanPatternSimulationToReturn.ElasticModulus, *scanPatternSimulation.ElasticModulus, "Expected ElasticModulus values to match")
	assert.EqualValues(t, *scanPatternSimulationToReturn.LayerThickness, *scanPatternSimulation.LayerThickness, "Expected LayerThickness values to match")

}

// AssumedStrainSimulations
func TestAssumedStrainSimulationWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	assumedStrainSimulationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	assumedStrainSimulationService := NewClient(fakeTokenFetcher, "", "", audience)

	// act
	_, err := assumedStrainSimulationService.AssumedStrainSimulation(assumedStrainSimulationID)

	// assert

	assert.Equal(t, expectedError, err, "Expected an error returned")
}

func TestAssumedStrainSimulationWhenAssumedStrainSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	assumedStrainSimulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// AssumedStrainSimulation
	assumedStrainSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	assumedStrainSimulationEndpoint := "/" + SimulationAPIBasePath + "/assumedstrainsimulations/{assumedStrainSimulationID}"
	r.HandleFunc(assumedStrainSimulationEndpoint, assumedStrainSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	assumedStrainSimulationService := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := assumedStrainSimulationService.AssumedStrainSimulation(assumedStrainSimulationID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because assumedStrainSimulation api send a 500 error")
}

func TestAssumedStrainSimulationWhenSuccessfulExpectsAssumedStrainSimulationReturned(t *testing.T) {
	// arrange
	assumedStrainSimulationID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// AssumedStrainSimulation
	assumedStrainSimulationToReturn := &models.AssumedStrainSimulation{
		Simulation: models.Simulation{
			ID:    assumedStrainSimulationID,
			Title: swag.String("AssumedStrainSimulation name"),
		},
		PartBasedSimulationParameters: models.PartBasedSimulationParameters{
			ElasticModulus: swag.Float64(10),
		},
		AssumedStrainSimulationParameters: models.AssumedStrainSimulationParameters{
			LayerThickness: 50e-6,
		},
	}
	assumedStrainSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedAssumedStrainSimulationID, err := strconv.Atoi(mux.Vars(r)["assumedStrainSimulationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(assumedStrainSimulationID), receivedAssumedStrainSimulationID, "Expected assumedStrainSimulation id received to match what was passed in")
		bytes, err := json.Marshal(assumedStrainSimulationToReturn)
		if err != nil {
			t.Error("Failed to marshal assumedStrainSimulation")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	assumedStrainSimulationEndpoint := "/" + SimulationAPIBasePath + "/assumedstrainsimulations/{assumedStrainSimulationID}"
	r.HandleFunc(assumedStrainSimulationEndpoint, assumedStrainSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	assumedStrainSimulationService := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	assumedStrainSimulation, err := assumedStrainSimulationService.AssumedStrainSimulation(assumedStrainSimulationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *assumedStrainSimulationToReturn.Title, *assumedStrainSimulation.Title, "Expected names to match")
	assert.Equal(t, assumedStrainSimulationToReturn.ID, assumedStrainSimulation.ID, "Expected IDs to match")
	assert.EqualValues(t, *assumedStrainSimulationToReturn.ElasticModulus, *assumedStrainSimulation.ElasticModulus, "Expected ElasticModulus values to match")
	assert.EqualValues(t, assumedStrainSimulationToReturn.LayerThickness, assumedStrainSimulation.LayerThickness, "Expected LayerThickness values to match")
}

func TestSimulationsWhenNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	var organizationID int32 = 10
	offset := 2
	limit := 10

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Simulation
	simulationsToReturn := []models.Simulation{
		models.Simulation{
			ID: 1,
		},
		models.Simulation{
			ID: 2,
		},
	}

	simulationsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		assert.Equal(t, strconv.Itoa(offset), r.URL.Query().Get("offset"), "Expected offset to be passed as query param")
		assert.Equal(t, strconv.Itoa(limit), r.URL.Query().Get("limit"), "Expected limit to be passed as query param")
		bytes, err := json.Marshal(simulationsToReturn)
		if err != nil {
			t.Error("Failed to marshal simulations")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	simulationEndpoint := "/" + SimulationAPIBasePath + "/simulations"
	r.HandleFunc(simulationEndpoint, simulationsHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	simulations, err := client.Simulations(organizationID, nil, nil, int32(offset), int32(limit), nil, nil, nil)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Len(t, simulations, 2, "Expected 2 simulations returned.")
}

func TestSimulationsWhenNonNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	var organizationID int32 = 10
	offset := 2
	limit := 10
	sort := "field1:asc"
	status := models.SimulationStatusInProgress
	archived := true

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Simulation
	simulationsToReturn := []models.Simulation{
		models.Simulation{
			ID: 1,
		},
		models.Simulation{
			ID: 2,
		},
	}

	simulationsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		assert.Equal(t, strconv.Itoa(offset), r.URL.Query().Get("offset"), "Expected offset to be passed as query param")
		assert.Equal(t, strconv.Itoa(limit), r.URL.Query().Get("limit"), "Expected limit to be passed as query param")
		assert.Equal(t, status, r.URL.Query().Get("status"), "Expected status to be passed as query param")
		assert.Equal(t, sort, r.URL.Query().Get("sort"), "Expected sort to be passed as query param")
		bytes, err := json.Marshal(simulationsToReturn)
		if err != nil {
			t.Error("Failed to marshal simulations")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	simulationEndpoint := "/" + SimulationAPIBasePath + "/simulations"
	r.HandleFunc(simulationEndpoint, simulationsHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	simulations, err := client.Simulations(organizationID, []string{status}, []string{sort},
		int32(offset), int32(limit), &archived, nil, nil)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Len(t, simulations, 2, "Expected 2 simulations returned.")
}

func TestRawSimulationWhenSuccessfulExpectsSimulationMapReturned(t *testing.T) {
	// arrange
	simulationID := int32(2)
	layerThickness := float64(50e-6)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Simulation
	simulationToReturn := &models.Simulation{
		ID:    simulationID,
		Title: swag.String("Simulation name"),
		Type:  models.SimulationTypeAssumedStrainSimulation,
	}
	simulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedSimulationID, err := strconv.Atoi(mux.Vars(r)["simulationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationID), receivedSimulationID, "Expected simulation id received to match what was passed in")
		bytes, err := json.Marshal(simulationToReturn)
		if err != nil {
			t.Error("Failed to marshal simulation")
		}
		w.Write(bytes)
	})

	assumedStrainSimulationToReturn := &models.AssumedStrainSimulation{
		Simulation: models.Simulation{
			ID:    simulationID,
			Title: swag.String("Simulation name"),
			Type:  models.SimulationTypeAssumedStrainSimulation,
		},
		PartBasedSimulationParameters: models.PartBasedSimulationParameters{
			ElasticModulus: swag.Float64(10),
		},
		AssumedStrainSimulationParameters: models.AssumedStrainSimulationParameters{
			LayerThickness: layerThickness,
		},
	}
	assumedStrainSimulationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedSimulationID, err := strconv.Atoi(mux.Vars(r)["simulationID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationID), receivedSimulationID, "Expected simulation id received to match what was passed in")
		bytes, err := json.Marshal(assumedStrainSimulationToReturn)
		if err != nil {
			t.Error("Failed to marshal simulation")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	simulationEndpoint := "/" + SimulationAPIBasePath + "/simulations/{simulationID}"
	assumedStrainSimulationEndpoint := "/" + SimulationAPIBasePath + "/assumedstrainsimulations/{simulationID}"
	r.HandleFunc(simulationEndpoint, simulationHandler)
	r.HandleFunc(assumedStrainSimulationEndpoint, assumedStrainSimulationHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	simulation, err := client.RawSimulation(simulationID)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *simulationToReturn.Title, simulation["title"], "Expected names to match")
	assert.EqualValues(t, simulationToReturn.ID, simulation["id"], "Expected IDs to match")
	assert.EqualValues(t, *assumedStrainSimulationToReturn.ElasticModulus, simulation["elasticModulus"].(float64), "Expected elastic modulus to match")
	assert.EqualValues(t, layerThickness, simulation["layerThickness"], "Expected layer thickness to match")
}

func TestBuildFilesWithNonNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	var organizationID int32 = 10
	offset := 2
	limit := 10
	sort := "field1:asc"
	availability := models.BuildFileAvailabilityAvailable

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Build Files
	buildFilesToReturn := []models.BuildFile{
		models.BuildFile{
			ID: swag.Int32(1),
		},
		models.BuildFile{
			ID: swag.Int32(2),
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		assert.Equal(t, strconv.Itoa(offset), r.URL.Query().Get("offset"), "Expected offset to be passed as query param")
		assert.Equal(t, strconv.Itoa(limit), r.URL.Query().Get("limit"), "Expected limit to be passed as query param")
		assert.Equal(t, availability, r.URL.Query().Get("availability"), "Expected availability to be passed as query param")
		assert.Equal(t, sort, r.URL.Query().Get("sort"), "Expected sort to be passed as query param")
		bytes, err := json.Marshal(buildFilesToReturn)
		if err != nil {
			t.Error("Failed to marshal simulations")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	buildFiles, err := client.BuildFiles(organizationID, []string{availability}, []string{sort},
		int32(offset), int32(limit))

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Len(t, buildFiles, 2, "Expected 2 build files returned.")
}

func TestBuildFilesWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	organizationID := int32(2)
	offset := int32(0)
	limit := int32(10)
	var availability []string
	var sort []string

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.BuildFiles(organizationID, availability, sort, offset, limit)

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestBuildFileWithNonNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	buildFileID := int32(10)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	buildFileToReturn := models.BuildFile{
		ID:           swag.Int32(buildFileID),
		Name:         swag.String("MyBuildFile"),
		Availability: swag.String("Available"),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedBuildFileID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(buildFileID), receivedBuildFileID, "Expected build file id received to match what was passed in")
		bytes, err := json.Marshal(buildFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	buildFile, err := client.BuildFile(int32(buildFileID))

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, buildFile)
	assert.Equal(t, *buildFileToReturn.ID, *buildFile.ID, "Expected ID values to match")
	assert.Equal(t, *buildFileToReturn.Name, *buildFile.Name, "Expected Name values to match")
	assert.Equal(t, *buildFileToReturn.Availability, *buildFile.Availability, "Expected Availability values to match")
}

func TestBuildFileWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	buidFileID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Machine
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.BuildFile(buidFileID)

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestPostBuildFileWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Build Files
	buildFileToPost := models.BuildFilePost{
		OrganizationID:     swag.Int32(10),
		FileUploadLocation: swag.String("s3bucket/mybuildfile.zip"),
		MachineType:        swag.String(models.BuildFileMachineTypeAdditiveIndustries),
		Name:               swag.String("myBuildFile"),
		Description:        "My lucky build file",
		Tags:               []string{"tag1", "tag2"},
	}

	buildFileToReturn := models.BuildFile{
		ID:             swag.Int32(99),
		OrganizationID: *buildFileToPost.OrganizationID,
		Name:           buildFileToPost.Name,
		Tags:           buildFileToPost.Tags,
		Description:    buildFileToPost.Description,
		MachineType:    buildFileToPost.MachineType,
		Availability:   swag.String("Available"),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		bytes, err := json.Marshal(buildFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	buildFile, err := client.PostBuildFile(&buildFileToPost)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, buildFile)
	assert.Equal(t, *buildFileToReturn.ID, *buildFile.ID, "Expected ID values to match")
	assert.Equal(t, *buildFileToReturn.Name, *buildFile.Name, "Expected Name values to match")
	assert.Equal(t, *buildFileToReturn.Availability, *buildFile.Availability, "Expected Availability values to match")
	assert.Equal(t, buildFileToReturn.Description, buildFile.Description, "Expected Description values to match")
	assert.Equal(t, *buildFileToReturn.MachineType, *buildFile.MachineType, "Expected MachineType values to match")
}

func TestPostBuildFileWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	buildFileToPost := models.BuildFilePost{
		OrganizationID:     swag.Int32(10),
		FileUploadLocation: swag.String("s3bucket/mybuildfile.zip"),
		MachineType:        swag.String(models.BuildFileMachineTypeAdditiveIndustries),
		Name:               swag.String("myBuildFile"),
		Description:        "My lucky build file",
		Tags:               []string{"tag1", "tag2"},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.PostBuildFile(&buildFileToPost)

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestPatchBuildFileWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	availability := "Available"
	buildFileID := int32(101)
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	buildFileToReturn := models.BuildFile{
		ID: swag.Int32(buildFileID),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedBuildFileID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(buildFileID), receivedBuildFileID, "Expected build file id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal build file")
		}
		actualPatch := actualPatchDocuments[0]
		assert.EqualValues(t, *patch, actualPatch, "Expected patch document to be passed in body of PATCH request")
		bytes, err := json.Marshal(buildFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	buildFile, err := client.PatchBuildFile(int32(buildFileID), []*models.PatchDocument{patch})

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, buildFile)
	assert.Equal(t, buildFileID, *buildFile.ID, "Expected ID values to match")
}

func TestPatchBuildFileWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	buildFileID := int32(22)
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Build Files
	patch := models.PatchDocument{
		From:  "from",
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: "Available",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.PatchBuildFile(buildFileID, []*models.PatchDocument{&patch})

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}
func TestUpdateBuildFileAvailabilityWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	buildFileID := int32(97)
	availability := "Processing"

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	expectedPatch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	// Build Files
	buildFileToReturn := models.BuildFile{
		ID: swag.Int32(buildFileID),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedBuildFileID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(buildFileID), receivedBuildFileID, "Expected build file id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal build file")
		}
		actualPatch := actualPatchDocuments[0]
		assert.EqualValues(t, *expectedPatch, actualPatch, "Expected patch document to be passed in body of PATCH request")
		bytes, err := json.Marshal(buildFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	buildFile, err := client.UpdateBuildFileAvailability(int32(buildFileID), availability)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, buildFile, "Expected buildFile to not be nil")
	assert.Equal(t, buildFileID, *buildFile.ID, "Expected build file IDs to match")
}

func TestUpdateBuildFileAvailabilityWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.UpdateBuildFileAvailability(27, "Available")

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestRawBUildFileWhenSuccessfulExpectsBuildFileMapReturned(t *testing.T) {
	// arrange
	buildFileID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Simulation
	buildFileToReturn := &models.BuildFile{
		ID:          swag.Int32(buildFileID),
		Name:        swag.String("Build File Name"),
		MachineType: swag.String(models.BuildFileMachineTypeAdditiveIndustries),
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedSimulationID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(buildFileID), receivedSimulationID, "Expected build file id received to match what was passed in")
		bytes, err := json.Marshal(buildFileToReturn)
		if err != nil {
			t.Error("Failed to marshal simulation")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/buildfiles/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	buildFileMap, err := client.RawBuildFile(int32(buildFileID))

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *buildFileToReturn.Name, buildFileMap["name"], "Expected names to match")
	assert.EqualValues(t, *buildFileToReturn.ID, buildFileMap["id"], "Expected IDs to match")
}

func TestPatchPartWithNonNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	availability := "Available"
	partID := int32(101)
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	partToReturn := models.Part{
		ID: &partID,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedPartID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(partID), receivedPartID, "Expected part id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal part")
		}
		actualPatch := actualPatchDocuments[0]
		assert.EqualValues(t, *patch, actualPatch, "Expected patch document to be passed in body of PATCH request")
		bytes, err := json.Marshal(partToReturn)
		if err != nil {
			t.Error("Failed to marshal part")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/parts/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	part, err := client.PatchPart(int32(partID), []*models.PatchDocument{patch})

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, part)
	assert.Equal(t, partID, *part.ID, "Expected ID values to match")
}

func TestPatchPartWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	partID := int32(22)
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Parts
	patch := models.PatchDocument{
		From:  "from",
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: "Available",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/parts/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.PatchPart(partID, []*models.PatchDocument{&patch})

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestPartWithNonNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	partID := int32(10)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	partToReturn := models.Part{
		ID:           &partID,
		Name:         swag.String("MyPart"),
		Availability: "Available",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedPartID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(partID), receivedPartID, "Expected part id received to match what was passed in")
		bytes, err := json.Marshal(partToReturn)
		if err != nil {
			t.Error("Failed to marshal part")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/parts/{id}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	part, err := client.Part(partID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, part)
	assert.Equal(t, partToReturn.ID, part.ID, "Expected ID values to match")
	assert.Equal(t, *partToReturn.Name, *part.Name, "Expected Name values to match")
	assert.Equal(t, partToReturn.Availability, part.Availability, "Expected Availability values to match")
}

func TestPartWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	buidFileID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Machine
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/parts"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.Part(buidFileID)

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestNewClientWithRetryWhen500ExpectsRetry(t *testing.T) {
	// arrange
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	callCounter := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCounter++
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/"+SimulationAPIBasePath+"/parts/1", handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClientWithRetry(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience, 3*time.Second)

	// act
	_, err := client.Part(1)

	// assert
	assert.True(t, callCounter > 1, "Expected to retry the failed call at least once")
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 500 error")
}

func TestPatchSimulationActivityWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	simulationID := int32(10)
	activityID := "activity id"
	simulationActivityID := int32(11)
	percentComplete := int32(98)
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/percentComplete"),
		Value: percentComplete,
	}

	activityToReturn := &models.SimulationActivity{
		ID:              simulationActivityID,
		ActivityID:      swag.String(activityID),
		SimulationID:    swag.Int32(simulationID),
		PercentComplete: percentComplete,
	}

	activitiesHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		activitiesBytes, err := json.Marshal([]models.SimulationActivity{*activityToReturn})
		if err != nil {
			t.Fatal(err)
		}
		w.Write(activitiesBytes)
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationID), receivedID, "Expected simulation id received to match what was passed in")
		receivedActivityID, err := strconv.Atoi(mux.Vars(r)["activityId"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationActivityID), receivedActivityID, "Expected activity id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal patch documents")
		}
		actualPatch := actualPatchDocuments[0]
		assert.Equal(t, *patch.Path, *actualPatch.Path, "Expected patch document path to be passed in body of PATCH request")
		assert.Equal(t, *patch.Op, *actualPatch.Op, "Expected patch document op to be passed in body of PATCH request")
		assert.EqualValues(t, patch.Value, actualPatch.Value, "Expected patch document value to be passed in body of PATCH request")
		bytes, err := json.Marshal(activityToReturn)
		if err != nil {
			t.Error("Failed to marshal patch")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/simulations/{id}/activities/{activityId}"
	r.HandleFunc(endpoint, handler)
	r.HandleFunc("/"+SimulationAPIBasePath+"/simulations/{id}/activities", activitiesHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	activity, err := client.PatchSimulationActivity(simulationID, activityID, []*models.PatchDocument{patch})

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, activity)
	assert.Equal(t, simulationActivityID, activity.ID, "Expected ID values to match")
	assert.Equal(t, simulationID, *activity.SimulationID, "Expected simulation ID values to match")
	assert.Equal(t, percentComplete, activity.PercentComplete, "Expected percent complete values to match")
}

func TestPutSimulationActivityWithNonNilValuesExpectsSuccess(t *testing.T) {
	// arrange
	simulationID := int32(10)
	activityID := "activity id"
	simulationActivityID := int32(11)
	percentComplete := int32(98)
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	simulationActivity := &models.SimulationActivity{
		ID:              simulationActivityID,
		ActivityID:      swag.String(activityID),
		SimulationID:    swag.Int32(simulationID),
		PercentComplete: percentComplete,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationID), receivedID, "Expected simulation id received to match what was passed in")
		receivedActivityID, err := strconv.Atoi(mux.Vars(r)["activityId"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(simulationActivityID), receivedActivityID, "Expected activity id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualSimulationActivity models.SimulationActivity
		err = json.Unmarshal(bodyBytes, &actualSimulationActivity)
		if err != nil {
			t.Error("Failed to unmarshal activity documents")
		}
		assert.Equal(t, simulationActivity.ID, actualSimulationActivity.ID, "Expected simulation activity id to be passed in body")
		w.WriteHeader(http.StatusOK)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/simulations/{id}/activities/{activityId}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	err := client.PutSimulationActivity(simulationID, simulationActivity)

	// assert

	assert.Nil(t, err, "Expected no error returned")
}

func TestPatchSimulationActivityWhenSimulationAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	simulationID := int32(10)
	activityID := "activity id"
	percentComplete := int32(98)
	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/percentComplete"),
		Value: percentComplete,
	}

	// Setup routes
	r := mux.NewRouter()
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	activity, err := client.PatchSimulationActivity(simulationID, activityID, []*models.PatchDocument{patch})

	// assert
	assert.NotNil(t, err, "Expected an error returned because simulation api sent a 404 error")
	assert.Nil(t, activity, "Expected activity to be nil due to error from api")
}

func TestPatchPartSupportWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	availability := "Available"
	partSupportID := int32(101)
	partID := int32(102)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	partSupportFileToReturn := models.PartSupport{
		ID:     swag.Int32(partSupportID),
		PartID: swag.Int32(partID),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedPartSupportID, err := strconv.Atoi(mux.Vars(r)["supportID"])
		if err != nil {
			t.Fatal(err)
		}
		receivedPartID, err := strconv.Atoi(mux.Vars(r)["partID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(partSupportID), receivedPartSupportID, "Expected part support id received to match what was passed in")
		assert.EqualValues(t, int(partID), receivedPartID, "Expected part id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal build file")
		}
		actualPatch := actualPatchDocuments[0]
		assert.EqualValues(t, *patch, actualPatch, "Expected patch document to be passed in body of PATCH request")
		bytes, err := json.Marshal(partSupportFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/parts/{partID}/supports/{supportID}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	partSupport, err := client.PatchPartSupport(int32(partID), int32(partSupportID), []*models.PatchDocument{patch})

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, partSupport)
	assert.Equal(t, partSupportID, *partSupport.ID, "Expected ID values to match")
}

func TestPatchPartSupportWhenPatchErrorsExpectsErrorReturned(t *testing.T) {

	// arrange
	availability := "Available"
	partSupportID := int32(101)
	partID := int32(102)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/parts/{partID}/supports/{supportID}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.PatchPartSupport(int32(partID), int32(partSupportID), []*models.PatchDocument{patch})

	// assert

	assert.Error(t, err, "Expected an error returned because patch returned a 500")
}

func TestPatchPartSupportWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {

	// arrange
	partSupportID := int32(101)
	partID := int32(102)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	expectedError := errors.New("some token error")
	fakeTokenFetcher.TokenReturns("Token", expectedError)

	client := NewClient(fakeTokenFetcher, "some url", SimulationAPIBasePath, audience)

	// act
	_, err := client.PatchPartSupport(int32(partID), int32(partSupportID), []*models.PatchDocument{})

	// assert
	assert.Error(t, err, "Expected an error returned because token fetcher errored")
}

func TestPartSupportByIDWhenSuccessfulExpectsPartSupportReturned(t *testing.T) {
	// arrange
	partSupportID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	// Material
	partSupportToReturn := &models.PartSupport{
		ID:   swag.Int32(partSupportID),
		Name: swag.String("Part support name"),
	}
	partSupportHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedPartSupportID, err := strconv.Atoi(mux.Vars(r)["partSupportID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(partSupportID), receivedPartSupportID, "Expected partSupport id received to match what was passed in")
		bytes, err := json.Marshal(partSupportToReturn)
		if err != nil {
			t.Error("Failed to marshal returned object")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	materialEndpoint := "/" + SimulationAPIBasePath + "/supports/{partSupportID}"
	r.HandleFunc(materialEndpoint, partSupportHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	partSupport, err := client.PartSupportByID(partSupportID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *partSupportToReturn.Name, *partSupport.Name, "Expected names to match")
	assert.Equal(t, *partSupportToReturn.ID, *partSupport.ID, "Expected IDs to match")
}

func TestPartSupportByIDWhenAPIErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	partSupportID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	partSupportHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	materialEndpoint := "/" + SimulationAPIBasePath + "/supports/{partSupportID}"
	r.HandleFunc(materialEndpoint, partSupportHandler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.PartSupportByID(partSupportID)

	// assert
	assert.Error(t, err, "Expected an error returned because API returned a 500")
}

func TestPartSupportByIDWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	partSupportID := int32(2)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	expectedError := errors.New("some error")
	fakeTokenFetcher.TokenReturns("Token", expectedError)

	client := NewClient(fakeTokenFetcher, "some url", SimulationAPIBasePath, audience)

	// act
	_, err := client.PartSupportByID(partSupportID)

	// assert

	assert.Error(t, err, "Expected error returned because token fetcher errored")
}

func TestPatchPartSupportByIDWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	availability := "Available"
	partSupportID := int32(101)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	partSupportFileToReturn := models.PartSupport{
		ID:     swag.Int32(partSupportID),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedPartSupportID, err := strconv.Atoi(mux.Vars(r)["supportID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(partSupportID), receivedPartSupportID, "Expected part support id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal build file")
		}
		actualPatch := actualPatchDocuments[0]
		assert.EqualValues(t, *patch, actualPatch, "Expected patch document to be passed in body of PATCH request")
		bytes, err := json.Marshal(partSupportFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/supports/{supportID}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	partSupport, err := client.PatchPartSupportByID(int32(partSupportID), []*models.PatchDocument{patch})

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, partSupport)
	assert.Equal(t, partSupportID, *partSupport.ID, "Expected ID values to match")
}

func TestPatchPartSupportByIDWhenPatchErrorsExpectsErrorReturned(t *testing.T) {

	// arrange
	availability := "Available"
	partSupportID := int32(101)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/supports/{supportID}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	_, err := client.PatchPartSupportByID(int32(partSupportID), []*models.PatchDocument{patch})

	// assert
	assert.Error(t, err, "Expected an error returned because patch returned a 500")
}

func TestPatchPartSupportByIDWhenTokenFetcherErrorsExpectsErrorReturned(t *testing.T) {

	// arrange
	partSupportID := int32(101)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	expectedError := errors.New("some token error")
	fakeTokenFetcher.TokenReturns("Token", expectedError)

	client := NewClient(fakeTokenFetcher, "some url", SimulationAPIBasePath, audience)

	// act
	_, err := client.PatchPartSupportByID(int32(partSupportID), []*models.PatchDocument{})

	// assert
	assert.Error(t, err, "Expected an error returned because token fetcher errored")
}

func TestUpdatePartSupportAvailabilityWithNonNilValuesExpectsSuccess(t *testing.T) {

	// arrange
	availability := "Available"
	partSupportID := int32(101)

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("Token", nil)

	patch := &models.PatchDocument{
		Op:    swag.String(models.PatchDocumentOpReplace),
		Path:  swag.String("/availability"),
		Value: availability,
	}

	partSupportFileToReturn := models.PartSupport{
		ID:     swag.Int32(partSupportID),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assert.NotEmpty(t, r.Header.Get("Authorization"), "Authorization header should not be empty")
		receivedPartSupportID, err := strconv.Atoi(mux.Vars(r)["supportID"])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, int(partSupportID), receivedPartSupportID, "Expected part support id received to match what was passed in")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var actualPatchDocuments []models.PatchDocument
		err = json.Unmarshal(bodyBytes, &actualPatchDocuments)
		if err != nil {
			t.Error("Failed to unmarshal build file")
		}
		actualPatch := actualPatchDocuments[0]
		assert.EqualValues(t, *patch, actualPatch, "Expected patch document to be passed in body of PATCH request")
		bytes, err := json.Marshal(partSupportFileToReturn)
		if err != nil {
			t.Error("Failed to marshal build file")
		}
		w.Write(bytes)
	})

	// Setup routes
	r := mux.NewRouter()
	endpoint := "/" + SimulationAPIBasePath + "/supports/{supportID}"
	r.HandleFunc(endpoint, handler)
	testServer := httptest.NewServer(r)
	defer testServer.Close()
	client := NewClient(fakeTokenFetcher, testServer.URL, SimulationAPIBasePath, audience)

	// act
	partSupport, err := client.UpdatePartSupportAvailability(int32(partSupportID), availability)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.NotNil(t, partSupport)
	assert.Equal(t, partSupportID, *partSupport.ID, "Expected ID values to match")
}

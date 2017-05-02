package simulation

import (
	"errors"
	"testing"

	"encoding/json"
	"github.com/3dsim/auth0/auth0fakes"
	"github.com/3dsim/simulation-goclient/models"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
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

	client := NewClient(fakeTokenFetcher, "apiGatewayURL", audience)

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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	_, err := client.Machine(machineID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because simulation api send a 500 error")
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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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

	client := NewClient(fakeTokenFetcher, "", audience)

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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	_, err := client.Material(materialID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because simulation api send a 500 error")
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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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

	client := NewClient(fakeTokenFetcher, "", audience)

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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	_, err := client.Simulation(simulationID)

	// assert

	assert.NotNil(t, err, "Expected an error returned because simulation api send a 500 error")
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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

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

	thermalSimulationService := NewClient(fakeTokenFetcher, "", audience)

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
	thermalSimulationService := NewClient(fakeTokenFetcher, testServer.URL, audience)

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
		ThermalSimulationParameters: models.ThermalSimulationParameters{
			ElasticModulus: swag.Float64(21),
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
	thermalSimulationService := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	thermalSimulation, err := thermalSimulationService.ThermalSimulation(thermalSimulationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *thermalSimulationToReturn.Title, *thermalSimulation.Title, "Expected names to match")
	assert.Equal(t, thermalSimulationToReturn.ID, thermalSimulation.ID, "Expected IDs to match")
	assert.EqualValues(t, thermalSimulationToReturn.ElasticModulus, thermalSimulation.ElasticModulus, "Expected ElasticModulus values to match")
}

// ScanPatternSimulations
func TestScanPatternSimulationWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	scanPatternSimulationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	scanPatternSimulationService := NewClient(fakeTokenFetcher, "", audience)

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
	scanPatternSimulationService := NewClient(fakeTokenFetcher, testServer.URL, audience)

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
		ScanPatternSimulationParameters: models.ScanPatternSimulationParameters{
			ElasticModulus: swag.Float64(21),
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
	scanPatternSimulationService := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	scanPatternSimulation, err := scanPatternSimulationService.ScanPatternSimulation(scanPatternSimulationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *scanPatternSimulationToReturn.Title, *scanPatternSimulation.Title, "Expected names to match")
	assert.Equal(t, scanPatternSimulationToReturn.ID, scanPatternSimulation.ID, "Expected IDs to match")
	assert.EqualValues(t, scanPatternSimulationToReturn.ElasticModulus, scanPatternSimulation.ElasticModulus, "Expected ElasticModulus values to match")
}

// AssumedStrainSimulations
func TestAssumedStrainSimulationWhenFetcherErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	assumedStrainSimulationID := int32(2)
	expectedError := errors.New("Some auth0 error")

	// Token
	fakeTokenFetcher := &auth0fakes.FakeTokenFetcher{}
	fakeTokenFetcher.TokenReturns("", expectedError)

	assumedStrainSimulationService := NewClient(fakeTokenFetcher, "", audience)

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
	assumedStrainSimulationService := NewClient(fakeTokenFetcher, testServer.URL, audience)

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
		AssumedStrainSimulationParameters: models.AssumedStrainSimulationParameters{
			ElasticModulus: swag.Float64(21),
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
	assumedStrainSimulationService := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	assumedStrainSimulation, err := assumedStrainSimulationService.AssumedStrainSimulation(assumedStrainSimulationID)

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *assumedStrainSimulationToReturn.Title, *assumedStrainSimulation.Title, "Expected names to match")
	assert.Equal(t, assumedStrainSimulationToReturn.ID, assumedStrainSimulation.ID, "Expected IDs to match")
	assert.EqualValues(t, assumedStrainSimulationToReturn.ElasticModulus, assumedStrainSimulation.ElasticModulus, "Expected ElasticModulus values to match")
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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	simulations, err := client.Simulations(organizationID, nil, nil, int32(offset), int32(limit))

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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	simulations, err := client.Simulations(organizationID, []string{status}, []string{sort},
		int32(offset), int32(limit))

	// assert

	assert.Nil(t, err, "Expected no error returned")
	assert.Len(t, simulations, 2, "Expected 2 simulations returned.")
}

func TestRawSimulationWhenSuccessfulExpectsSimulationMapReturned(t *testing.T) {
	// arrange
	simulationID := int32(2)

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
		AssumedStrainSimulationParameters: models.AssumedStrainSimulationParameters{
			OutputDisplacementAfterCutoff: swag.Bool(true),
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
	client := NewClient(fakeTokenFetcher, testServer.URL, audience)

	// act
	simulation, err := client.RawSimulation(simulationID)

	// assert
	assert.Nil(t, err, "Expected no error returned")
	assert.Equal(t, *simulationToReturn.Title, simulation["title"], "Expected names to match")
	assert.EqualValues(t, simulationToReturn.ID, simulation["id"], "Expected IDs to match")
	_, ok := simulation["outputDisplacementAfterCutoff"].(bool)
	assert.True(t, ok, "Expected outputDisplacementAfterCutoff to be of type bool")
}

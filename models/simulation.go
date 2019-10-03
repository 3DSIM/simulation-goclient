// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Simulation simulation
// swagger:model Simulation
type Simulation struct {

	// True if simulation has been archived, false otherwise
	Archived bool `json:"archived,omitempty"`

	// time stamp, set server-side, read only
	Completed strfmt.DateTime `json:"completed,omitempty"`

	// created time stamp, set server-side, read only
	Created strfmt.DateTime `json:"created,omitempty"`

	// assigned user, set server-side, read only
	CreatedBy string `json:"createdBy,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// last modified time stamp, set server-side, read only
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// assigned user, set server-side, read only
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// length (in meters) of the longest fill scan vector in the scan pattern
	LongestScanVectorLength float64 `json:"longestScanVectorLength,omitempty"`

	// ID of the machine configuration that was used in the simulation, set server-side, read only
	MachineConfigurationID int32 `json:"machineConfigurationId,omitempty"`

	// ID of the machine to use in the simulation
	MachineID int32 `json:"machineId,omitempty"`

	// ID of the material configuration that was used in the simulation, set server-side, read only
	MaterialConfigurationID int32 `json:"materialConfigurationId,omitempty"`

	// ID of the material to use in the simulation
	MaterialID int32 `json:"materialId,omitempty"`

	// indicate the number of cores with which to run the simulation
	NumberOfCores int32 `json:"numberOfCores,omitempty"`

	// organization Id
	// Required: true
	OrganizationID *int32 `json:"organizationId"`

	// parent Id
	ParentID *int32 `json:"parentId,omitempty"`

	// set server-side, read only
	PercentComplete float64 `json:"percentComplete,omitempty"`

	// s3 folder used to store input and output files for the simulation
	SimulationFolder string `json:"simulationFolder,omitempty"`

	// number of slice layers in the part
	SliceLayers int32 `json:"sliceLayers,omitempty"`

	// time stamp, set server-side, read only
	Started strfmt.DateTime `json:"started,omitempty"`

	// set server-side, read only
	StartedBy string `json:"startedBy,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	Type SimulationType `json:"type,omitempty"`

	// number of voxel layers in the voxelized part + supports
	VoxelLayers int32 `json:"voxelLayers,omitempty"`

	// workflow identifier - will be provided by the system
	WorkflowID string `json:"workflowId,omitempty"`
}

// Validate validates this simulation
func (m *Simulation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Simulation) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

var simulationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Draft","Requested","InProgress","Cancelled","Error","Success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simulationTypeStatusPropEnum = append(simulationTypeStatusPropEnum, v)
	}
}

const (
	// SimulationStatusDraft captures enum value "Draft"
	SimulationStatusDraft string = "Draft"
	// SimulationStatusRequested captures enum value "Requested"
	SimulationStatusRequested string = "Requested"
	// SimulationStatusInProgress captures enum value "InProgress"
	SimulationStatusInProgress string = "InProgress"
	// SimulationStatusCancelled captures enum value "Cancelled"
	SimulationStatusCancelled string = "Cancelled"
	// SimulationStatusError captures enum value "Error"
	SimulationStatusError string = "Error"
	// SimulationStatusSuccess captures enum value "Success"
	SimulationStatusSuccess string = "Success"
)

// prop value enum
func (m *Simulation) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, simulationTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Simulation) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Simulation) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

func (m *Simulation) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Simulation) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Simulation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Simulation) UnmarshalBinary(b []byte) error {
	var res Simulation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

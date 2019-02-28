// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPartsParams creates a new GetPartsParams object
// with the default values initialized.
func NewGetPartsParams() *GetPartsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetPartsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPartsParamsWithTimeout creates a new GetPartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPartsParamsWithTimeout(timeout time.Duration) *GetPartsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetPartsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetPartsParamsWithContext creates a new GetPartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPartsParamsWithContext(ctx context.Context) *GetPartsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetPartsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetPartsParamsWithHTTPClient creates a new GetPartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPartsParamsWithHTTPClient(client *http.Client) *GetPartsParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetPartsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetPartsParams contains all the parameters to send to the API endpoint
for the get parts operation typically these are written to a http.Request
*/
type GetPartsParams struct {

	/*Archived
	  If true, will only return archived parts.  If false, will only return unarchived parts.  If not provided, will return both archived and unarchived parts.

	*/
	Archived *bool
	/*Availability
	  If specified, will filter to only include parts with the given availability.  Uploaded - the part has been uploaded.  Processing - the part is being processed.  Available - the part was processed successfully and can be used in simulations.  Error - an error occurred, contact Ansys Support.

	*/
	Availability *string
	/*Limit
	  number of items to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*OrganizationID
	  the organization id to get items for.  Must be provided as API callers only have access to items belonging to their organization.

	*/
	OrganizationID int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders.  e.g. sort=key1:desc,key2:asc

	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get parts params
func (o *GetPartsParams) WithTimeout(timeout time.Duration) *GetPartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parts params
func (o *GetPartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parts params
func (o *GetPartsParams) WithContext(ctx context.Context) *GetPartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parts params
func (o *GetPartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parts params
func (o *GetPartsParams) WithHTTPClient(client *http.Client) *GetPartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parts params
func (o *GetPartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchived adds the archived to the get parts params
func (o *GetPartsParams) WithArchived(archived *bool) *GetPartsParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the get parts params
func (o *GetPartsParams) SetArchived(archived *bool) {
	o.Archived = archived
}

// WithAvailability adds the availability to the get parts params
func (o *GetPartsParams) WithAvailability(availability *string) *GetPartsParams {
	o.SetAvailability(availability)
	return o
}

// SetAvailability adds the availability to the get parts params
func (o *GetPartsParams) SetAvailability(availability *string) {
	o.Availability = availability
}

// WithLimit adds the limit to the get parts params
func (o *GetPartsParams) WithLimit(limit *int32) *GetPartsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get parts params
func (o *GetPartsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get parts params
func (o *GetPartsParams) WithOffset(offset *int32) *GetPartsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get parts params
func (o *GetPartsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get parts params
func (o *GetPartsParams) WithOrganizationID(organizationID int32) *GetPartsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get parts params
func (o *GetPartsParams) SetOrganizationID(organizationID int32) {
	o.OrganizationID = organizationID
}

// WithSort adds the sort to the get parts params
func (o *GetPartsParams) WithSort(sort []string) *GetPartsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get parts params
func (o *GetPartsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetPartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Archived != nil {

		// query param archived
		var qrArchived bool
		if o.Archived != nil {
			qrArchived = *o.Archived
		}
		qArchived := swag.FormatBool(qrArchived)
		if qArchived != "" {
			if err := r.SetQueryParam("archived", qArchived); err != nil {
				return err
			}
		}

	}

	if o.Availability != nil {

		// query param availability
		var qrAvailability string
		if o.Availability != nil {
			qrAvailability = *o.Availability
		}
		qAvailability := qrAvailability
		if qAvailability != "" {
			if err := r.SetQueryParam("availability", qAvailability); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param organizationId
	qrOrganizationID := o.OrganizationID
	qOrganizationID := swag.FormatInt32(qrOrganizationID)
	if qOrganizationID != "" {
		if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
			return err
		}
	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

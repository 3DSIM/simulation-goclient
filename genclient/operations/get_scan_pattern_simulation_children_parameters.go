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

// NewGetScanPatternSimulationChildrenParams creates a new GetScanPatternSimulationChildrenParams object
// with the default values initialized.
func NewGetScanPatternSimulationChildrenParams() *GetScanPatternSimulationChildrenParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetScanPatternSimulationChildrenParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanPatternSimulationChildrenParamsWithTimeout creates a new GetScanPatternSimulationChildrenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScanPatternSimulationChildrenParamsWithTimeout(timeout time.Duration) *GetScanPatternSimulationChildrenParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetScanPatternSimulationChildrenParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetScanPatternSimulationChildrenParamsWithContext creates a new GetScanPatternSimulationChildrenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScanPatternSimulationChildrenParamsWithContext(ctx context.Context) *GetScanPatternSimulationChildrenParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetScanPatternSimulationChildrenParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetScanPatternSimulationChildrenParamsWithHTTPClient creates a new GetScanPatternSimulationChildrenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScanPatternSimulationChildrenParamsWithHTTPClient(client *http.Client) *GetScanPatternSimulationChildrenParams {
	var (
		limitDefault  = int32(10)
		offsetDefault = int32(0)
	)
	return &GetScanPatternSimulationChildrenParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetScanPatternSimulationChildrenParams contains all the parameters to send to the API endpoint
for the get scan pattern simulation children operation typically these are written to a http.Request
*/
type GetScanPatternSimulationChildrenParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32
	/*Limit
	  number of materials to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. offset of 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders.  e.g. sort=key1:desc,key2:asc

	*/
	Sort []string
	/*Status
	  simulation status for items retrieved.  If an array of items is sent, they are treated as "OR" operations. e.g. status=InProgress,Requested would yield a list of simulations that are in either state.

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithTimeout(timeout time.Duration) *GetScanPatternSimulationChildrenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithContext(ctx context.Context) *GetScanPatternSimulationChildrenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithHTTPClient(client *http.Client) *GetScanPatternSimulationChildrenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithID(id int32) *GetScanPatternSimulationChildrenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetID(id int32) {
	o.ID = id
}

// WithLimit adds the limit to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithLimit(limit *int32) *GetScanPatternSimulationChildrenParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithOffset(offset *int32) *GetScanPatternSimulationChildrenParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithSort(sort []string) *GetScanPatternSimulationChildrenParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithStatus adds the status to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) WithStatus(status []string) *GetScanPatternSimulationChildrenParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get scan pattern simulation children params
func (o *GetScanPatternSimulationChildrenParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanPatternSimulationChildrenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

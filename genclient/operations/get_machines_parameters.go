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

// NewGetMachinesParams creates a new GetMachinesParams object
// with the default values initialized.
func NewGetMachinesParams() *GetMachinesParams {
	var ()
	return &GetMachinesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachinesParamsWithTimeout creates a new GetMachinesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachinesParamsWithTimeout(timeout time.Duration) *GetMachinesParams {
	var ()
	return &GetMachinesParams{

		timeout: timeout,
	}
}

// NewGetMachinesParamsWithContext creates a new GetMachinesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachinesParamsWithContext(ctx context.Context) *GetMachinesParams {
	var ()
	return &GetMachinesParams{

		Context: ctx,
	}
}

// NewGetMachinesParamsWithHTTPClient creates a new GetMachinesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMachinesParamsWithHTTPClient(client *http.Client) *GetMachinesParams {
	var ()
	return &GetMachinesParams{
		HTTPClient: client,
	}
}

/*GetMachinesParams contains all the parameters to send to the API endpoint
for the get machines operation typically these are written to a http.Request
*/
type GetMachinesParams struct {

	/*Archived
	  If true, will only return archived parts.  If false, will only return unarchived parts.  If not provided, will return both archived and unarchived parts.

	*/
	Archived *bool
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

// WithTimeout adds the timeout to the get machines params
func (o *GetMachinesParams) WithTimeout(timeout time.Duration) *GetMachinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machines params
func (o *GetMachinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machines params
func (o *GetMachinesParams) WithContext(ctx context.Context) *GetMachinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machines params
func (o *GetMachinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machines params
func (o *GetMachinesParams) WithHTTPClient(client *http.Client) *GetMachinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machines params
func (o *GetMachinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchived adds the archived to the get machines params
func (o *GetMachinesParams) WithArchived(archived *bool) *GetMachinesParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the get machines params
func (o *GetMachinesParams) SetArchived(archived *bool) {
	o.Archived = archived
}

// WithLimit adds the limit to the get machines params
func (o *GetMachinesParams) WithLimit(limit *int32) *GetMachinesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get machines params
func (o *GetMachinesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get machines params
func (o *GetMachinesParams) WithOffset(offset *int32) *GetMachinesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get machines params
func (o *GetMachinesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get machines params
func (o *GetMachinesParams) WithOrganizationID(organizationID int32) *GetMachinesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get machines params
func (o *GetMachinesParams) SetOrganizationID(organizationID int32) {
	o.OrganizationID = organizationID
}

// WithSort adds the sort to the get machines params
func (o *GetMachinesParams) WithSort(sort []string) *GetMachinesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get machines params
func (o *GetMachinesParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

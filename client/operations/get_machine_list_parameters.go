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

// NewGetMachineListParams creates a new GetMachineListParams object
// with the default values initialized.
func NewGetMachineListParams() *GetMachineListParams {
	var ()
	return &GetMachineListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineListParamsWithTimeout creates a new GetMachineListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachineListParamsWithTimeout(timeout time.Duration) *GetMachineListParams {
	var ()
	return &GetMachineListParams{

		timeout: timeout,
	}
}

// NewGetMachineListParamsWithContext creates a new GetMachineListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachineListParamsWithContext(ctx context.Context) *GetMachineListParams {
	var ()
	return &GetMachineListParams{

		Context: ctx,
	}
}

/*GetMachineListParams contains all the parameters to send to the API endpoint
for the get machine list operation typically these are written to a http.Request
*/
type GetMachineListParams struct {

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
	  the organization id to get parts for.  Must be provided as API callers only have access to parts belonging to their organization.

	*/
	OrganizationID string
	/*Sort
	  key:direction pairs for one or multiple field sort orders

	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get machine list params
func (o *GetMachineListParams) WithTimeout(timeout time.Duration) *GetMachineListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine list params
func (o *GetMachineListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine list params
func (o *GetMachineListParams) WithContext(ctx context.Context) *GetMachineListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine list params
func (o *GetMachineListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithArchived adds the archived to the get machine list params
func (o *GetMachineListParams) WithArchived(archived *bool) *GetMachineListParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the get machine list params
func (o *GetMachineListParams) SetArchived(archived *bool) {
	o.Archived = archived
}

// WithLimit adds the limit to the get machine list params
func (o *GetMachineListParams) WithLimit(limit *int32) *GetMachineListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get machine list params
func (o *GetMachineListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get machine list params
func (o *GetMachineListParams) WithOffset(offset *int32) *GetMachineListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get machine list params
func (o *GetMachineListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get machine list params
func (o *GetMachineListParams) WithOrganizationID(organizationID string) *GetMachineListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get machine list params
func (o *GetMachineListParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithSort adds the sort to the get machine list params
func (o *GetMachineListParams) WithSort(sort []string) *GetMachineListParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get machine list params
func (o *GetMachineListParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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
	qOrganizationID := qrOrganizationID
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

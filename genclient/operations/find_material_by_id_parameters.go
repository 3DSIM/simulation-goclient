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

// NewFindMaterialByIDParams creates a new FindMaterialByIDParams object
// with the default values initialized.
func NewFindMaterialByIDParams() *FindMaterialByIDParams {
	var ()
	return &FindMaterialByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindMaterialByIDParamsWithTimeout creates a new FindMaterialByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindMaterialByIDParamsWithTimeout(timeout time.Duration) *FindMaterialByIDParams {
	var ()
	return &FindMaterialByIDParams{

		timeout: timeout,
	}
}

// NewFindMaterialByIDParamsWithContext creates a new FindMaterialByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindMaterialByIDParamsWithContext(ctx context.Context) *FindMaterialByIDParams {
	var ()
	return &FindMaterialByIDParams{

		Context: ctx,
	}
}

/*FindMaterialByIDParams contains all the parameters to send to the API endpoint
for the find material by Id operation typically these are written to a http.Request
*/
type FindMaterialByIDParams struct {

	/*ID
	  ID of material to fetch

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find material by Id params
func (o *FindMaterialByIDParams) WithTimeout(timeout time.Duration) *FindMaterialByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find material by Id params
func (o *FindMaterialByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find material by Id params
func (o *FindMaterialByIDParams) WithContext(ctx context.Context) *FindMaterialByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find material by Id params
func (o *FindMaterialByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the find material by Id params
func (o *FindMaterialByIDParams) WithID(id int32) *FindMaterialByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find material by Id params
func (o *FindMaterialByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindMaterialByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

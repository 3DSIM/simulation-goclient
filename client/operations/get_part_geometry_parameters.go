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

// NewGetPartGeometryParams creates a new GetPartGeometryParams object
// with the default values initialized.
func NewGetPartGeometryParams() *GetPartGeometryParams {
	var ()
	return &GetPartGeometryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPartGeometryParamsWithTimeout creates a new GetPartGeometryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPartGeometryParamsWithTimeout(timeout time.Duration) *GetPartGeometryParams {
	var ()
	return &GetPartGeometryParams{

		timeout: timeout,
	}
}

// NewGetPartGeometryParamsWithContext creates a new GetPartGeometryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPartGeometryParamsWithContext(ctx context.Context) *GetPartGeometryParams {
	var ()
	return &GetPartGeometryParams{

		Context: ctx,
	}
}

/*GetPartGeometryParams contains all the parameters to send to the API endpoint
for the get part geometry operation typically these are written to a http.Request
*/
type GetPartGeometryParams struct {

	/*ID
	  ID of part

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get part geometry params
func (o *GetPartGeometryParams) WithTimeout(timeout time.Duration) *GetPartGeometryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get part geometry params
func (o *GetPartGeometryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get part geometry params
func (o *GetPartGeometryParams) WithContext(ctx context.Context) *GetPartGeometryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get part geometry params
func (o *GetPartGeometryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the get part geometry params
func (o *GetPartGeometryParams) WithID(id int32) *GetPartGeometryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get part geometry params
func (o *GetPartGeometryParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPartGeometryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

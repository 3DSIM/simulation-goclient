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

// NewCancelThermalSimulationParams creates a new CancelThermalSimulationParams object
// with the default values initialized.
func NewCancelThermalSimulationParams() *CancelThermalSimulationParams {
	var ()
	return &CancelThermalSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelThermalSimulationParamsWithTimeout creates a new CancelThermalSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelThermalSimulationParamsWithTimeout(timeout time.Duration) *CancelThermalSimulationParams {
	var ()
	return &CancelThermalSimulationParams{

		timeout: timeout,
	}
}

// NewCancelThermalSimulationParamsWithContext creates a new CancelThermalSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelThermalSimulationParamsWithContext(ctx context.Context) *CancelThermalSimulationParams {
	var ()
	return &CancelThermalSimulationParams{

		Context: ctx,
	}
}

/*CancelThermalSimulationParams contains all the parameters to send to the API endpoint
for the cancel thermal simulation operation typically these are written to a http.Request
*/
type CancelThermalSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel thermal simulation params
func (o *CancelThermalSimulationParams) WithTimeout(timeout time.Duration) *CancelThermalSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel thermal simulation params
func (o *CancelThermalSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel thermal simulation params
func (o *CancelThermalSimulationParams) WithContext(ctx context.Context) *CancelThermalSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel thermal simulation params
func (o *CancelThermalSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the cancel thermal simulation params
func (o *CancelThermalSimulationParams) WithID(id int64) *CancelThermalSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel thermal simulation params
func (o *CancelThermalSimulationParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelThermalSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

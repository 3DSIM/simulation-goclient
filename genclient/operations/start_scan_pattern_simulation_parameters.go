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

// NewStartScanPatternSimulationParams creates a new StartScanPatternSimulationParams object
// with the default values initialized.
func NewStartScanPatternSimulationParams() *StartScanPatternSimulationParams {
	var ()
	return &StartScanPatternSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartScanPatternSimulationParamsWithTimeout creates a new StartScanPatternSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartScanPatternSimulationParamsWithTimeout(timeout time.Duration) *StartScanPatternSimulationParams {
	var ()
	return &StartScanPatternSimulationParams{

		timeout: timeout,
	}
}

// NewStartScanPatternSimulationParamsWithContext creates a new StartScanPatternSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartScanPatternSimulationParamsWithContext(ctx context.Context) *StartScanPatternSimulationParams {
	var ()
	return &StartScanPatternSimulationParams{

		Context: ctx,
	}
}

// NewStartScanPatternSimulationParamsWithHTTPClient creates a new StartScanPatternSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartScanPatternSimulationParamsWithHTTPClient(client *http.Client) *StartScanPatternSimulationParams {
	var ()
	return &StartScanPatternSimulationParams{
		HTTPClient: client,
	}
}

/*StartScanPatternSimulationParams contains all the parameters to send to the API endpoint
for the start scan pattern simulation operation typically these are written to a http.Request
*/
type StartScanPatternSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) WithTimeout(timeout time.Duration) *StartScanPatternSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) WithContext(ctx context.Context) *StartScanPatternSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) WithHTTPClient(client *http.Client) *StartScanPatternSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) WithID(id int32) *StartScanPatternSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the start scan pattern simulation params
func (o *StartScanPatternSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StartScanPatternSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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

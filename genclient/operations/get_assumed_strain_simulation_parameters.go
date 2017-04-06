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

// NewGetAssumedStrainSimulationParams creates a new GetAssumedStrainSimulationParams object
// with the default values initialized.
func NewGetAssumedStrainSimulationParams() *GetAssumedStrainSimulationParams {
	var ()
	return &GetAssumedStrainSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssumedStrainSimulationParamsWithTimeout creates a new GetAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssumedStrainSimulationParamsWithTimeout(timeout time.Duration) *GetAssumedStrainSimulationParams {
	var ()
	return &GetAssumedStrainSimulationParams{

		timeout: timeout,
	}
}

// NewGetAssumedStrainSimulationParamsWithContext creates a new GetAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssumedStrainSimulationParamsWithContext(ctx context.Context) *GetAssumedStrainSimulationParams {
	var ()
	return &GetAssumedStrainSimulationParams{

		Context: ctx,
	}
}

// NewGetAssumedStrainSimulationParamsWithHTTPClient creates a new GetAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssumedStrainSimulationParamsWithHTTPClient(client *http.Client) *GetAssumedStrainSimulationParams {
	var ()
	return &GetAssumedStrainSimulationParams{
		HTTPClient: client,
	}
}

/*GetAssumedStrainSimulationParams contains all the parameters to send to the API endpoint
for the get assumed strain simulation operation typically these are written to a http.Request
*/
type GetAssumedStrainSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) WithTimeout(timeout time.Duration) *GetAssumedStrainSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) WithContext(ctx context.Context) *GetAssumedStrainSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) WithHTTPClient(client *http.Client) *GetAssumedStrainSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) WithID(id int32) *GetAssumedStrainSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get assumed strain simulation params
func (o *GetAssumedStrainSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssumedStrainSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

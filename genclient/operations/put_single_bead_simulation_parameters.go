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

	"github.com/3dsim/simulation-goclient/models"
)

// NewPutSingleBeadSimulationParams creates a new PutSingleBeadSimulationParams object
// with the default values initialized.
func NewPutSingleBeadSimulationParams() *PutSingleBeadSimulationParams {
	var ()
	return &PutSingleBeadSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSingleBeadSimulationParamsWithTimeout creates a new PutSingleBeadSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSingleBeadSimulationParamsWithTimeout(timeout time.Duration) *PutSingleBeadSimulationParams {
	var ()
	return &PutSingleBeadSimulationParams{

		timeout: timeout,
	}
}

// NewPutSingleBeadSimulationParamsWithContext creates a new PutSingleBeadSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSingleBeadSimulationParamsWithContext(ctx context.Context) *PutSingleBeadSimulationParams {
	var ()
	return &PutSingleBeadSimulationParams{

		Context: ctx,
	}
}

// NewPutSingleBeadSimulationParamsWithHTTPClient creates a new PutSingleBeadSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSingleBeadSimulationParamsWithHTTPClient(client *http.Client) *PutSingleBeadSimulationParams {
	var ()
	return &PutSingleBeadSimulationParams{
		HTTPClient: client,
	}
}

/*PutSingleBeadSimulationParams contains all the parameters to send to the API endpoint
for the put single bead simulation operation typically these are written to a http.Request
*/
type PutSingleBeadSimulationParams struct {

	/*SingleBeadSimulation
	  SingleBeadSimulation fields required to update a simulation

	*/
	SingleBeadSimulation *models.SingleBeadSimulation
	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) WithTimeout(timeout time.Duration) *PutSingleBeadSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) WithContext(ctx context.Context) *PutSingleBeadSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) WithHTTPClient(client *http.Client) *PutSingleBeadSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSingleBeadSimulation adds the singleBeadSimulation to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) WithSingleBeadSimulation(singleBeadSimulation *models.SingleBeadSimulation) *PutSingleBeadSimulationParams {
	o.SetSingleBeadSimulation(singleBeadSimulation)
	return o
}

// SetSingleBeadSimulation adds the singleBeadSimulation to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) SetSingleBeadSimulation(singleBeadSimulation *models.SingleBeadSimulation) {
	o.SingleBeadSimulation = singleBeadSimulation
}

// WithID adds the id to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) WithID(id int32) *PutSingleBeadSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put single bead simulation params
func (o *PutSingleBeadSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutSingleBeadSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SingleBeadSimulation == nil {
		o.SingleBeadSimulation = new(models.SingleBeadSimulation)
	}

	if err := r.SetBodyParam(o.SingleBeadSimulation); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

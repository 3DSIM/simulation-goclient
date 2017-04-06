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

// NewPutSimulationActivityParams creates a new PutSimulationActivityParams object
// with the default values initialized.
func NewPutSimulationActivityParams() *PutSimulationActivityParams {
	var ()
	return &PutSimulationActivityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSimulationActivityParamsWithTimeout creates a new PutSimulationActivityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSimulationActivityParamsWithTimeout(timeout time.Duration) *PutSimulationActivityParams {
	var ()
	return &PutSimulationActivityParams{

		timeout: timeout,
	}
}

// NewPutSimulationActivityParamsWithContext creates a new PutSimulationActivityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSimulationActivityParamsWithContext(ctx context.Context) *PutSimulationActivityParams {
	var ()
	return &PutSimulationActivityParams{

		Context: ctx,
	}
}

// NewPutSimulationActivityParamsWithHTTPClient creates a new PutSimulationActivityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSimulationActivityParamsWithHTTPClient(client *http.Client) *PutSimulationActivityParams {
	var ()
	return &PutSimulationActivityParams{
		HTTPClient: client,
	}
}

/*PutSimulationActivityParams contains all the parameters to send to the API endpoint
for the put simulation activity operation typically these are written to a http.Request
*/
type PutSimulationActivityParams struct {

	/*SimulationActivity
	  An activity for a simulation.  An activity represents the execution of a worker in the simulation pipeline.

	*/
	SimulationActivity *models.SimulationActivity
	/*ActivityID
	  activity identifier

	*/
	ActivityID int32
	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put simulation activity params
func (o *PutSimulationActivityParams) WithTimeout(timeout time.Duration) *PutSimulationActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put simulation activity params
func (o *PutSimulationActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put simulation activity params
func (o *PutSimulationActivityParams) WithContext(ctx context.Context) *PutSimulationActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put simulation activity params
func (o *PutSimulationActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put simulation activity params
func (o *PutSimulationActivityParams) WithHTTPClient(client *http.Client) *PutSimulationActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put simulation activity params
func (o *PutSimulationActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSimulationActivity adds the simulationActivity to the put simulation activity params
func (o *PutSimulationActivityParams) WithSimulationActivity(simulationActivity *models.SimulationActivity) *PutSimulationActivityParams {
	o.SetSimulationActivity(simulationActivity)
	return o
}

// SetSimulationActivity adds the simulationActivity to the put simulation activity params
func (o *PutSimulationActivityParams) SetSimulationActivity(simulationActivity *models.SimulationActivity) {
	o.SimulationActivity = simulationActivity
}

// WithActivityID adds the activityID to the put simulation activity params
func (o *PutSimulationActivityParams) WithActivityID(activityID int32) *PutSimulationActivityParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the put simulation activity params
func (o *PutSimulationActivityParams) SetActivityID(activityID int32) {
	o.ActivityID = activityID
}

// WithID adds the id to the put simulation activity params
func (o *PutSimulationActivityParams) WithID(id int32) *PutSimulationActivityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put simulation activity params
func (o *PutSimulationActivityParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutSimulationActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SimulationActivity == nil {
		o.SimulationActivity = new(models.SimulationActivity)
	}

	if err := r.SetBodyParam(o.SimulationActivity); err != nil {
		return err
	}

	// path param activityId
	if err := r.SetPathParam("activityId", swag.FormatInt32(o.ActivityID)); err != nil {
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

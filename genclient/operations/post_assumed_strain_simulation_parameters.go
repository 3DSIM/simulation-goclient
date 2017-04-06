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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// NewPostAssumedStrainSimulationParams creates a new PostAssumedStrainSimulationParams object
// with the default values initialized.
func NewPostAssumedStrainSimulationParams() *PostAssumedStrainSimulationParams {
	var ()
	return &PostAssumedStrainSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAssumedStrainSimulationParamsWithTimeout creates a new PostAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAssumedStrainSimulationParamsWithTimeout(timeout time.Duration) *PostAssumedStrainSimulationParams {
	var ()
	return &PostAssumedStrainSimulationParams{

		timeout: timeout,
	}
}

// NewPostAssumedStrainSimulationParamsWithContext creates a new PostAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAssumedStrainSimulationParamsWithContext(ctx context.Context) *PostAssumedStrainSimulationParams {
	var ()
	return &PostAssumedStrainSimulationParams{

		Context: ctx,
	}
}

/*PostAssumedStrainSimulationParams contains all the parameters to send to the API endpoint
for the post assumed strain simulation operation typically these are written to a http.Request
*/
type PostAssumedStrainSimulationParams struct {

	/*AssumedStrainSimulation
	  AssumedStrainSimulation fields required to add a simulation

	*/
	AssumedStrainSimulation *models.AssumedStrainSimulation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post assumed strain simulation params
func (o *PostAssumedStrainSimulationParams) WithTimeout(timeout time.Duration) *PostAssumedStrainSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post assumed strain simulation params
func (o *PostAssumedStrainSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post assumed strain simulation params
func (o *PostAssumedStrainSimulationParams) WithContext(ctx context.Context) *PostAssumedStrainSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post assumed strain simulation params
func (o *PostAssumedStrainSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAssumedStrainSimulation adds the assumedStrainSimulation to the post assumed strain simulation params
func (o *PostAssumedStrainSimulationParams) WithAssumedStrainSimulation(assumedStrainSimulation *models.AssumedStrainSimulation) *PostAssumedStrainSimulationParams {
	o.SetAssumedStrainSimulation(assumedStrainSimulation)
	return o
}

// SetAssumedStrainSimulation adds the assumedStrainSimulation to the post assumed strain simulation params
func (o *PostAssumedStrainSimulationParams) SetAssumedStrainSimulation(assumedStrainSimulation *models.AssumedStrainSimulation) {
	o.AssumedStrainSimulation = assumedStrainSimulation
}

// WriteToRequest writes these params to a swagger request
func (o *PostAssumedStrainSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.AssumedStrainSimulation == nil {
		o.AssumedStrainSimulation = new(models.AssumedStrainSimulation)
	}

	if err := r.SetBodyParam(o.AssumedStrainSimulation); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
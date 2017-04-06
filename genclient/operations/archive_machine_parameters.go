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

// NewArchiveMachineParams creates a new ArchiveMachineParams object
// with the default values initialized.
func NewArchiveMachineParams() *ArchiveMachineParams {
	var ()
	return &ArchiveMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArchiveMachineParamsWithTimeout creates a new ArchiveMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArchiveMachineParamsWithTimeout(timeout time.Duration) *ArchiveMachineParams {
	var ()
	return &ArchiveMachineParams{

		timeout: timeout,
	}
}

// NewArchiveMachineParamsWithContext creates a new ArchiveMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewArchiveMachineParamsWithContext(ctx context.Context) *ArchiveMachineParams {
	var ()
	return &ArchiveMachineParams{

		Context: ctx,
	}
}

// NewArchiveMachineParamsWithHTTPClient creates a new ArchiveMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArchiveMachineParamsWithHTTPClient(client *http.Client) *ArchiveMachineParams {
	var ()
	return &ArchiveMachineParams{
		HTTPClient: client,
	}
}

/*ArchiveMachineParams contains all the parameters to send to the API endpoint
for the archive machine operation typically these are written to a http.Request
*/
type ArchiveMachineParams struct {

	/*ID
	  machine identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the archive machine params
func (o *ArchiveMachineParams) WithTimeout(timeout time.Duration) *ArchiveMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the archive machine params
func (o *ArchiveMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the archive machine params
func (o *ArchiveMachineParams) WithContext(ctx context.Context) *ArchiveMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the archive machine params
func (o *ArchiveMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the archive machine params
func (o *ArchiveMachineParams) WithHTTPClient(client *http.Client) *ArchiveMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the archive machine params
func (o *ArchiveMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the archive machine params
func (o *ArchiveMachineParams) WithID(id int32) *ArchiveMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the archive machine params
func (o *ArchiveMachineParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ArchiveMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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
)

// NewStatusParams creates a new StatusParams object
// with the default values initialized.
func NewStatusParams() *StatusParams {

	return &StatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatusParamsWithTimeout creates a new StatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatusParamsWithTimeout(timeout time.Duration) *StatusParams {

	return &StatusParams{

		timeout: timeout,
	}
}

// NewStatusParamsWithContext creates a new StatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatusParamsWithContext(ctx context.Context) *StatusParams {

	return &StatusParams{

		Context: ctx,
	}
}

/*StatusParams contains all the parameters to send to the API endpoint
for the status operation typically these are written to a http.Request
*/
type StatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status params
func (o *StatusParams) WithTimeout(timeout time.Duration) *StatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status params
func (o *StatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status params
func (o *StatusParams) WithContext(ctx context.Context) *StatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status params
func (o *StatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *StatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

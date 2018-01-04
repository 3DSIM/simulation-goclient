// Code generated by go-swagger; DO NOT EDIT.

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

// NewUpdatePartParams creates a new UpdatePartParams object
// with the default values initialized.
func NewUpdatePartParams() *UpdatePartParams {
	var ()
	return &UpdatePartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePartParamsWithTimeout creates a new UpdatePartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePartParamsWithTimeout(timeout time.Duration) *UpdatePartParams {
	var ()
	return &UpdatePartParams{

		timeout: timeout,
	}
}

// NewUpdatePartParamsWithContext creates a new UpdatePartParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePartParamsWithContext(ctx context.Context) *UpdatePartParams {
	var ()
	return &UpdatePartParams{

		Context: ctx,
	}
}

// NewUpdatePartParamsWithHTTPClient creates a new UpdatePartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePartParamsWithHTTPClient(client *http.Client) *UpdatePartParams {
	var ()
	return &UpdatePartParams{
		HTTPClient: client,
	}
}

/*UpdatePartParams contains all the parameters to send to the API endpoint
for the update part operation typically these are written to a http.Request
*/
type UpdatePartParams struct {

	/*Part
	  Part to update.

	*/
	Part *models.Part

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update part params
func (o *UpdatePartParams) WithTimeout(timeout time.Duration) *UpdatePartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update part params
func (o *UpdatePartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update part params
func (o *UpdatePartParams) WithContext(ctx context.Context) *UpdatePartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update part params
func (o *UpdatePartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update part params
func (o *UpdatePartParams) WithHTTPClient(client *http.Client) *UpdatePartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update part params
func (o *UpdatePartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPart adds the part to the update part params
func (o *UpdatePartParams) WithPart(part *models.Part) *UpdatePartParams {
	o.SetPart(part)
	return o
}

// SetPart adds the part to the update part params
func (o *UpdatePartParams) SetPart(part *models.Part) {
	o.Part = part
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Part != nil {
		if err := r.SetBodyParam(o.Part); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

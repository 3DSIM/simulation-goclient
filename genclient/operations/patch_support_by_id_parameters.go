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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// NewPatchSupportByIDParams creates a new PatchSupportByIDParams object
// with the default values initialized.
func NewPatchSupportByIDParams() *PatchSupportByIDParams {
	var ()
	return &PatchSupportByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSupportByIDParamsWithTimeout creates a new PatchSupportByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSupportByIDParamsWithTimeout(timeout time.Duration) *PatchSupportByIDParams {
	var ()
	return &PatchSupportByIDParams{

		timeout: timeout,
	}
}

// NewPatchSupportByIDParamsWithContext creates a new PatchSupportByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSupportByIDParamsWithContext(ctx context.Context) *PatchSupportByIDParams {
	var ()
	return &PatchSupportByIDParams{

		Context: ctx,
	}
}

// NewPatchSupportByIDParamsWithHTTPClient creates a new PatchSupportByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSupportByIDParamsWithHTTPClient(client *http.Client) *PatchSupportByIDParams {
	var ()
	return &PatchSupportByIDParams{
		HTTPClient: client,
	}
}

/*PatchSupportByIDParams contains all the parameters to send to the API endpoint
for the patch support by Id operation typically these are written to a http.Request
*/
type PatchSupportByIDParams struct {

	/*ID
	  ID of support to update

	*/
	ID int32
	/*SupportPatch
	  This endpoint uses JSON Patch, RFC 6092.

	*/
	SupportPatch []*models.PatchDocument

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch support by Id params
func (o *PatchSupportByIDParams) WithTimeout(timeout time.Duration) *PatchSupportByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch support by Id params
func (o *PatchSupportByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch support by Id params
func (o *PatchSupportByIDParams) WithContext(ctx context.Context) *PatchSupportByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch support by Id params
func (o *PatchSupportByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch support by Id params
func (o *PatchSupportByIDParams) WithHTTPClient(client *http.Client) *PatchSupportByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch support by Id params
func (o *PatchSupportByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch support by Id params
func (o *PatchSupportByIDParams) WithID(id int32) *PatchSupportByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch support by Id params
func (o *PatchSupportByIDParams) SetID(id int32) {
	o.ID = id
}

// WithSupportPatch adds the supportPatch to the patch support by Id params
func (o *PatchSupportByIDParams) WithSupportPatch(supportPatch []*models.PatchDocument) *PatchSupportByIDParams {
	o.SetSupportPatch(supportPatch)
	return o
}

// SetSupportPatch adds the supportPatch to the patch support by Id params
func (o *PatchSupportByIDParams) SetSupportPatch(supportPatch []*models.PatchDocument) {
	o.SupportPatch = supportPatch
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSupportByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.SupportPatch); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
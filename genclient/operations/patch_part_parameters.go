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

// NewPatchPartParams creates a new PatchPartParams object
// with the default values initialized.
func NewPatchPartParams() *PatchPartParams {
	var ()
	return &PatchPartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPartParamsWithTimeout creates a new PatchPartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPartParamsWithTimeout(timeout time.Duration) *PatchPartParams {
	var ()
	return &PatchPartParams{

		timeout: timeout,
	}
}

// NewPatchPartParamsWithContext creates a new PatchPartParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPartParamsWithContext(ctx context.Context) *PatchPartParams {
	var ()
	return &PatchPartParams{

		Context: ctx,
	}
}

// NewPatchPartParamsWithHTTPClient creates a new PatchPartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPartParamsWithHTTPClient(client *http.Client) *PatchPartParams {
	var ()
	return &PatchPartParams{
		HTTPClient: client,
	}
}

/*PatchPartParams contains all the parameters to send to the API endpoint
for the patch part operation typically these are written to a http.Request
*/
type PatchPartParams struct {

	/*ID
	  ID of part to update

	*/
	ID int32
	/*PartPatch
	  This endpoint uses JSON Patch, RFC 6092.

	*/
	PartPatch models.PatchPartParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch part params
func (o *PatchPartParams) WithTimeout(timeout time.Duration) *PatchPartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch part params
func (o *PatchPartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch part params
func (o *PatchPartParams) WithContext(ctx context.Context) *PatchPartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch part params
func (o *PatchPartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch part params
func (o *PatchPartParams) WithHTTPClient(client *http.Client) *PatchPartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch part params
func (o *PatchPartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch part params
func (o *PatchPartParams) WithID(id int32) *PatchPartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch part params
func (o *PatchPartParams) SetID(id int32) {
	o.ID = id
}

// WithPartPatch adds the partPatch to the patch part params
func (o *PatchPartParams) WithPartPatch(partPatch models.PatchPartParamsBody) *PatchPartParams {
	o.SetPartPatch(partPatch)
	return o
}

// SetPartPatch adds the partPatch to the patch part params
func (o *PatchPartParams) SetPartPatch(partPatch models.PatchPartParamsBody) {
	o.PartPatch = partPatch
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

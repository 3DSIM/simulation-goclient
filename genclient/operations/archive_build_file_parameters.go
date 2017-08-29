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
)

// NewArchiveBuildFileParams creates a new ArchiveBuildFileParams object
// with the default values initialized.
func NewArchiveBuildFileParams() *ArchiveBuildFileParams {
	var ()
	return &ArchiveBuildFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArchiveBuildFileParamsWithTimeout creates a new ArchiveBuildFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArchiveBuildFileParamsWithTimeout(timeout time.Duration) *ArchiveBuildFileParams {
	var ()
	return &ArchiveBuildFileParams{

		timeout: timeout,
	}
}

// NewArchiveBuildFileParamsWithContext creates a new ArchiveBuildFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewArchiveBuildFileParamsWithContext(ctx context.Context) *ArchiveBuildFileParams {
	var ()
	return &ArchiveBuildFileParams{

		Context: ctx,
	}
}

// NewArchiveBuildFileParamsWithHTTPClient creates a new ArchiveBuildFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArchiveBuildFileParamsWithHTTPClient(client *http.Client) *ArchiveBuildFileParams {
	var ()
	return &ArchiveBuildFileParams{
		HTTPClient: client,
	}
}

/*ArchiveBuildFileParams contains all the parameters to send to the API endpoint
for the archive build file operation typically these are written to a http.Request
*/
type ArchiveBuildFileParams struct {

	/*ID
	  ID of build file to archive

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the archive build file params
func (o *ArchiveBuildFileParams) WithTimeout(timeout time.Duration) *ArchiveBuildFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the archive build file params
func (o *ArchiveBuildFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the archive build file params
func (o *ArchiveBuildFileParams) WithContext(ctx context.Context) *ArchiveBuildFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the archive build file params
func (o *ArchiveBuildFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the archive build file params
func (o *ArchiveBuildFileParams) WithHTTPClient(client *http.Client) *ArchiveBuildFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the archive build file params
func (o *ArchiveBuildFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the archive build file params
func (o *ArchiveBuildFileParams) WithID(id int32) *ArchiveBuildFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the archive build file params
func (o *ArchiveBuildFileParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ArchiveBuildFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
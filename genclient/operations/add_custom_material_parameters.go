// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddCustomMaterialParams creates a new AddCustomMaterialParams object
// with the default values initialized.
func NewAddCustomMaterialParams() *AddCustomMaterialParams {
	var ()
	return &AddCustomMaterialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCustomMaterialParamsWithTimeout creates a new AddCustomMaterialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCustomMaterialParamsWithTimeout(timeout time.Duration) *AddCustomMaterialParams {
	var ()
	return &AddCustomMaterialParams{

		timeout: timeout,
	}
}

// NewAddCustomMaterialParamsWithContext creates a new AddCustomMaterialParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCustomMaterialParamsWithContext(ctx context.Context) *AddCustomMaterialParams {
	var ()
	return &AddCustomMaterialParams{

		Context: ctx,
	}
}

// NewAddCustomMaterialParamsWithHTTPClient creates a new AddCustomMaterialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCustomMaterialParamsWithHTTPClient(client *http.Client) *AddCustomMaterialParams {
	var ()
	return &AddCustomMaterialParams{
		HTTPClient: client,
	}
}

/*AddCustomMaterialParams contains all the parameters to send to the API endpoint
for the add custom material operation typically these are written to a http.Request
*/
type AddCustomMaterialParams struct {

	/*Configuration
	  configuration defining material scientific specification, each field name must match expected material data template

	*/
	Configuration os.File
	/*CustomMaterialPost
	  json formatted data to generate the material -  schema /definitions/CustomMaterialPost

	*/
	CustomMaterialPost string
	/*Cwlookup
	  csv file defining columns Speed (mm/s), Power (W), CW (m)

	*/
	Cwlookup os.File
	/*Lookup
	  csv file defining columns Temperature (K), Thermal Conductivity (W/m/K), Specific Heat (J/kg/K), Density (kg/m3), Thermal Conductivity Ratio, Density Ratio, Specific Heat Ratio

	*/
	Lookup os.File

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add custom material params
func (o *AddCustomMaterialParams) WithTimeout(timeout time.Duration) *AddCustomMaterialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add custom material params
func (o *AddCustomMaterialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add custom material params
func (o *AddCustomMaterialParams) WithContext(ctx context.Context) *AddCustomMaterialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add custom material params
func (o *AddCustomMaterialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add custom material params
func (o *AddCustomMaterialParams) WithHTTPClient(client *http.Client) *AddCustomMaterialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add custom material params
func (o *AddCustomMaterialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the add custom material params
func (o *AddCustomMaterialParams) WithConfiguration(configuration os.File) *AddCustomMaterialParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the add custom material params
func (o *AddCustomMaterialParams) SetConfiguration(configuration os.File) {
	o.Configuration = configuration
}

// WithCustomMaterialPost adds the customMaterialPost to the add custom material params
func (o *AddCustomMaterialParams) WithCustomMaterialPost(customMaterialPost string) *AddCustomMaterialParams {
	o.SetCustomMaterialPost(customMaterialPost)
	return o
}

// SetCustomMaterialPost adds the customMaterialPost to the add custom material params
func (o *AddCustomMaterialParams) SetCustomMaterialPost(customMaterialPost string) {
	o.CustomMaterialPost = customMaterialPost
}

// WithCwlookup adds the cwlookup to the add custom material params
func (o *AddCustomMaterialParams) WithCwlookup(cwlookup os.File) *AddCustomMaterialParams {
	o.SetCwlookup(cwlookup)
	return o
}

// SetCwlookup adds the cwlookup to the add custom material params
func (o *AddCustomMaterialParams) SetCwlookup(cwlookup os.File) {
	o.Cwlookup = cwlookup
}

// WithLookup adds the lookup to the add custom material params
func (o *AddCustomMaterialParams) WithLookup(lookup os.File) *AddCustomMaterialParams {
	o.SetLookup(lookup)
	return o
}

// SetLookup adds the lookup to the add custom material params
func (o *AddCustomMaterialParams) SetLookup(lookup os.File) {
	o.Lookup = lookup
}

// WriteToRequest writes these params to a swagger request
func (o *AddCustomMaterialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param configuration
	if err := r.SetFileParam("configuration", &o.Configuration); err != nil {
		return err
	}

	// form param customMaterialPost
	frCustomMaterialPost := o.CustomMaterialPost
	fCustomMaterialPost := frCustomMaterialPost
	if fCustomMaterialPost != "" {
		if err := r.SetFormParam("customMaterialPost", fCustomMaterialPost); err != nil {
			return err
		}
	}

	// form file param cwlookup
	if err := r.SetFileParam("cwlookup", &o.Cwlookup); err != nil {
		return err
	}

	// form file param lookup
	if err := r.SetFileParam("lookup", &o.Lookup); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

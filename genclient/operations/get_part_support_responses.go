// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// GetPartSupportReader is a Reader for the GetPartSupport structure.
type GetPartSupportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPartSupportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPartSupportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetPartSupportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPartSupportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPartSupportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPartSupportOK creates a GetPartSupportOK with default headers values
func NewGetPartSupportOK() *GetPartSupportOK {
	return &GetPartSupportOK{}
}

/*GetPartSupportOK handles this case with default header values.

Successfully found a support
*/
type GetPartSupportOK struct {
	Payload *models.PartSupport
}

func (o *GetPartSupportOK) Error() string {
	return fmt.Sprintf("[GET /parts/{partId}/supports/{supportId}][%d] getPartSupportOK  %+v", 200, o.Payload)
}

func (o *GetPartSupportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartSupport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartSupportForbidden creates a GetPartSupportForbidden with default headers values
func NewGetPartSupportForbidden() *GetPartSupportForbidden {
	return &GetPartSupportForbidden{}
}

/*GetPartSupportForbidden handles this case with default header values.

Forbidden
*/
type GetPartSupportForbidden struct {
	Payload *models.Error
}

func (o *GetPartSupportForbidden) Error() string {
	return fmt.Sprintf("[GET /parts/{partId}/supports/{supportId}][%d] getPartSupportForbidden  %+v", 403, o.Payload)
}

func (o *GetPartSupportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartSupportNotFound creates a GetPartSupportNotFound with default headers values
func NewGetPartSupportNotFound() *GetPartSupportNotFound {
	return &GetPartSupportNotFound{}
}

/*GetPartSupportNotFound handles this case with default header values.

Not found
*/
type GetPartSupportNotFound struct {
}

func (o *GetPartSupportNotFound) Error() string {
	return fmt.Sprintf("[GET /parts/{partId}/supports/{supportId}][%d] getPartSupportNotFound ", 404)
}

func (o *GetPartSupportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartSupportDefault creates a GetPartSupportDefault with default headers values
func NewGetPartSupportDefault(code int) *GetPartSupportDefault {
	return &GetPartSupportDefault{
		_statusCode: code,
	}
}

/*GetPartSupportDefault handles this case with default header values.

unexpected error
*/
type GetPartSupportDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get part support default response
func (o *GetPartSupportDefault) Code() int {
	return o._statusCode
}

func (o *GetPartSupportDefault) Error() string {
	return fmt.Sprintf("[GET /parts/{partId}/supports/{supportId}][%d] getPartSupport default  %+v", o._statusCode, o.Payload)
}

func (o *GetPartSupportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
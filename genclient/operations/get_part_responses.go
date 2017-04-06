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

// GetPartReader is a Reader for the GetPart structure.
type GetPartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetPartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPartDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPartOK creates a GetPartOK with default headers values
func NewGetPartOK() *GetPartOK {
	return &GetPartOK{}
}

/*GetPartOK handles this case with default header values.

Successfully found a part
*/
type GetPartOK struct {
	Payload *models.Part
}

func (o *GetPartOK) Error() string {
	return fmt.Sprintf("[GET /parts/{id}][%d] getPartOK  %+v", 200, o.Payload)
}

func (o *GetPartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Part)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartForbidden creates a GetPartForbidden with default headers values
func NewGetPartForbidden() *GetPartForbidden {
	return &GetPartForbidden{}
}

/*GetPartForbidden handles this case with default header values.

Forbidden
*/
type GetPartForbidden struct {
	Payload *models.Error
}

func (o *GetPartForbidden) Error() string {
	return fmt.Sprintf("[GET /parts/{id}][%d] getPartForbidden  %+v", 403, o.Payload)
}

func (o *GetPartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartNotFound creates a GetPartNotFound with default headers values
func NewGetPartNotFound() *GetPartNotFound {
	return &GetPartNotFound{}
}

/*GetPartNotFound handles this case with default header values.

Not found
*/
type GetPartNotFound struct {
}

func (o *GetPartNotFound) Error() string {
	return fmt.Sprintf("[GET /parts/{id}][%d] getPartNotFound ", 404)
}

func (o *GetPartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPartDefault creates a GetPartDefault with default headers values
func NewGetPartDefault(code int) *GetPartDefault {
	return &GetPartDefault{
		_statusCode: code,
	}
}

/*GetPartDefault handles this case with default header values.

unexpected error
*/
type GetPartDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get part default response
func (o *GetPartDefault) Code() int {
	return o._statusCode
}

func (o *GetPartDefault) Error() string {
	return fmt.Sprintf("[GET /parts/{id}][%d] getPart default  %+v", o._statusCode, o.Payload)
}

func (o *GetPartDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
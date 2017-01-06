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

// AddPartsReader is a Reader for the AddParts structure.
type AddPartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddPartsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewAddPartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddPartsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPartsOK creates a AddPartsOK with default headers values
func NewAddPartsOK() *AddPartsOK {
	return &AddPartsOK{}
}

/*AddPartsOK handles this case with default header values.

Successfully added a part
*/
type AddPartsOK struct {
	Payload *models.Part
}

func (o *AddPartsOK) Error() string {
	return fmt.Sprintf("[POST /parts][%d] addPartsOK  %+v", 200, o.Payload)
}

func (o *AddPartsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Part)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartsForbidden creates a AddPartsForbidden with default headers values
func NewAddPartsForbidden() *AddPartsForbidden {
	return &AddPartsForbidden{}
}

/*AddPartsForbidden handles this case with default header values.

Forbidden
*/
type AddPartsForbidden struct {
	Payload *models.Error
}

func (o *AddPartsForbidden) Error() string {
	return fmt.Sprintf("[POST /parts][%d] addPartsForbidden  %+v", 403, o.Payload)
}

func (o *AddPartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartsDefault creates a AddPartsDefault with default headers values
func NewAddPartsDefault(code int) *AddPartsDefault {
	return &AddPartsDefault{
		_statusCode: code,
	}
}

/*AddPartsDefault handles this case with default header values.

unexpected error
*/
type AddPartsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add parts default response
func (o *AddPartsDefault) Code() int {
	return o._statusCode
}

func (o *AddPartsDefault) Error() string {
	return fmt.Sprintf("[POST /parts][%d] addParts default  %+v", o._statusCode, o.Payload)
}

func (o *AddPartsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

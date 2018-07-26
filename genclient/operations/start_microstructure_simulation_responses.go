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

// StartMicrostructureSimulationReader is a Reader for the StartMicrostructureSimulation structure.
type StartMicrostructureSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMicrostructureSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartMicrostructureSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewStartMicrostructureSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewStartMicrostructureSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStartMicrostructureSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewStartMicrostructureSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMicrostructureSimulationOK creates a StartMicrostructureSimulationOK with default headers values
func NewStartMicrostructureSimulationOK() *StartMicrostructureSimulationOK {
	return &StartMicrostructureSimulationOK{}
}

/*StartMicrostructureSimulationOK handles this case with default header values.

Simulation was successfully started.
*/
type StartMicrostructureSimulationOK struct {
}

func (o *StartMicrostructureSimulationOK) Error() string {
	return fmt.Sprintf("[PUT /microstructuresimulations/{id}/start][%d] startMicrostructureSimulationOK ", 200)
}

func (o *StartMicrostructureSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartMicrostructureSimulationUnauthorized creates a StartMicrostructureSimulationUnauthorized with default headers values
func NewStartMicrostructureSimulationUnauthorized() *StartMicrostructureSimulationUnauthorized {
	return &StartMicrostructureSimulationUnauthorized{}
}

/*StartMicrostructureSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type StartMicrostructureSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *StartMicrostructureSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /microstructuresimulations/{id}/start][%d] startMicrostructureSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *StartMicrostructureSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMicrostructureSimulationForbidden creates a StartMicrostructureSimulationForbidden with default headers values
func NewStartMicrostructureSimulationForbidden() *StartMicrostructureSimulationForbidden {
	return &StartMicrostructureSimulationForbidden{}
}

/*StartMicrostructureSimulationForbidden handles this case with default header values.

Forbidden
*/
type StartMicrostructureSimulationForbidden struct {
	Payload *models.Error
}

func (o *StartMicrostructureSimulationForbidden) Error() string {
	return fmt.Sprintf("[PUT /microstructuresimulations/{id}/start][%d] startMicrostructureSimulationForbidden  %+v", 403, o.Payload)
}

func (o *StartMicrostructureSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMicrostructureSimulationNotFound creates a StartMicrostructureSimulationNotFound with default headers values
func NewStartMicrostructureSimulationNotFound() *StartMicrostructureSimulationNotFound {
	return &StartMicrostructureSimulationNotFound{}
}

/*StartMicrostructureSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type StartMicrostructureSimulationNotFound struct {
	Payload *models.Error
}

func (o *StartMicrostructureSimulationNotFound) Error() string {
	return fmt.Sprintf("[PUT /microstructuresimulations/{id}/start][%d] startMicrostructureSimulationNotFound  %+v", 404, o.Payload)
}

func (o *StartMicrostructureSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMicrostructureSimulationDefault creates a StartMicrostructureSimulationDefault with default headers values
func NewStartMicrostructureSimulationDefault(code int) *StartMicrostructureSimulationDefault {
	return &StartMicrostructureSimulationDefault{
		_statusCode: code,
	}
}

/*StartMicrostructureSimulationDefault handles this case with default header values.

unexpected error
*/
type StartMicrostructureSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the start microstructure simulation default response
func (o *StartMicrostructureSimulationDefault) Code() int {
	return o._statusCode
}

func (o *StartMicrostructureSimulationDefault) Error() string {
	return fmt.Sprintf("[PUT /microstructuresimulations/{id}/start][%d] startMicrostructureSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *StartMicrostructureSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

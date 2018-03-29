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

// DeleteThermalSimulationReader is a Reader for the DeleteThermalSimulation structure.
type DeleteThermalSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteThermalSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteThermalSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteThermalSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteThermalSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteThermalSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteThermalSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteThermalSimulationOK creates a DeleteThermalSimulationOK with default headers values
func NewDeleteThermalSimulationOK() *DeleteThermalSimulationOK {
	return &DeleteThermalSimulationOK{}
}

/*DeleteThermalSimulationOK handles this case with default header values.

Successfully deleted a simulation
*/
type DeleteThermalSimulationOK struct {
}

func (o *DeleteThermalSimulationOK) Error() string {
	return fmt.Sprintf("[DELETE /thermalsimulations/{id}][%d] deleteThermalSimulationOK ", 200)
}

func (o *DeleteThermalSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteThermalSimulationUnauthorized creates a DeleteThermalSimulationUnauthorized with default headers values
func NewDeleteThermalSimulationUnauthorized() *DeleteThermalSimulationUnauthorized {
	return &DeleteThermalSimulationUnauthorized{}
}

/*DeleteThermalSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type DeleteThermalSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteThermalSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /thermalsimulations/{id}][%d] deleteThermalSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteThermalSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteThermalSimulationForbidden creates a DeleteThermalSimulationForbidden with default headers values
func NewDeleteThermalSimulationForbidden() *DeleteThermalSimulationForbidden {
	return &DeleteThermalSimulationForbidden{}
}

/*DeleteThermalSimulationForbidden handles this case with default header values.

Forbidden
*/
type DeleteThermalSimulationForbidden struct {
	Payload *models.Error
}

func (o *DeleteThermalSimulationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /thermalsimulations/{id}][%d] deleteThermalSimulationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteThermalSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteThermalSimulationNotFound creates a DeleteThermalSimulationNotFound with default headers values
func NewDeleteThermalSimulationNotFound() *DeleteThermalSimulationNotFound {
	return &DeleteThermalSimulationNotFound{}
}

/*DeleteThermalSimulationNotFound handles this case with default header values.

Not Found
*/
type DeleteThermalSimulationNotFound struct {
}

func (o *DeleteThermalSimulationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /thermalsimulations/{id}][%d] deleteThermalSimulationNotFound ", 404)
}

func (o *DeleteThermalSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteThermalSimulationDefault creates a DeleteThermalSimulationDefault with default headers values
func NewDeleteThermalSimulationDefault(code int) *DeleteThermalSimulationDefault {
	return &DeleteThermalSimulationDefault{
		_statusCode: code,
	}
}

/*DeleteThermalSimulationDefault handles this case with default header values.

unexpected error
*/
type DeleteThermalSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete thermal simulation default response
func (o *DeleteThermalSimulationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteThermalSimulationDefault) Error() string {
	return fmt.Sprintf("[DELETE /thermalsimulations/{id}][%d] DeleteThermalSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteThermalSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
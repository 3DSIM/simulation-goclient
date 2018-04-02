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

// GetThermalSimulationChildrenReader is a Reader for the GetThermalSimulationChildren structure.
type GetThermalSimulationChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetThermalSimulationChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetThermalSimulationChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetThermalSimulationChildrenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetThermalSimulationChildrenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetThermalSimulationChildrenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetThermalSimulationChildrenOK creates a GetThermalSimulationChildrenOK with default headers values
func NewGetThermalSimulationChildrenOK() *GetThermalSimulationChildrenOK {
	return &GetThermalSimulationChildrenOK{}
}

/*GetThermalSimulationChildrenOK handles this case with default header values.

Successfully found the list of items
*/
type GetThermalSimulationChildrenOK struct {
	/*Contains paging information in json format - totalCount, totalPages
	 */
	XPagination string

	Payload []*models.ThermalSimulation
}

func (o *GetThermalSimulationChildrenOK) Error() string {
	return fmt.Sprintf("[GET /thermalsimulations/{id}/children][%d] getThermalSimulationChildrenOK  %+v", 200, o.Payload)
}

func (o *GetThermalSimulationChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Pagination
	o.XPagination = response.GetHeader("X-Pagination")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetThermalSimulationChildrenUnauthorized creates a GetThermalSimulationChildrenUnauthorized with default headers values
func NewGetThermalSimulationChildrenUnauthorized() *GetThermalSimulationChildrenUnauthorized {
	return &GetThermalSimulationChildrenUnauthorized{}
}

/*GetThermalSimulationChildrenUnauthorized handles this case with default header values.

Not authorized
*/
type GetThermalSimulationChildrenUnauthorized struct {
	Payload *models.Error
}

func (o *GetThermalSimulationChildrenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /thermalsimulations/{id}/children][%d] getThermalSimulationChildrenUnauthorized  %+v", 401, o.Payload)
}

func (o *GetThermalSimulationChildrenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetThermalSimulationChildrenForbidden creates a GetThermalSimulationChildrenForbidden with default headers values
func NewGetThermalSimulationChildrenForbidden() *GetThermalSimulationChildrenForbidden {
	return &GetThermalSimulationChildrenForbidden{}
}

/*GetThermalSimulationChildrenForbidden handles this case with default header values.

Forbidden
*/
type GetThermalSimulationChildrenForbidden struct {
	Payload *models.Error
}

func (o *GetThermalSimulationChildrenForbidden) Error() string {
	return fmt.Sprintf("[GET /thermalsimulations/{id}/children][%d] getThermalSimulationChildrenForbidden  %+v", 403, o.Payload)
}

func (o *GetThermalSimulationChildrenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetThermalSimulationChildrenDefault creates a GetThermalSimulationChildrenDefault with default headers values
func NewGetThermalSimulationChildrenDefault(code int) *GetThermalSimulationChildrenDefault {
	return &GetThermalSimulationChildrenDefault{
		_statusCode: code,
	}
}

/*GetThermalSimulationChildrenDefault handles this case with default header values.

unexpected error
*/
type GetThermalSimulationChildrenDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get thermal simulation children default response
func (o *GetThermalSimulationChildrenDefault) Code() int {
	return o._statusCode
}

func (o *GetThermalSimulationChildrenDefault) Error() string {
	return fmt.Sprintf("[GET /thermalsimulations/{id}/children][%d] getThermalSimulationChildren default  %+v", o._statusCode, o.Payload)
}

func (o *GetThermalSimulationChildrenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

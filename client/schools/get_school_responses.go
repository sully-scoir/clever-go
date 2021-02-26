// Code generated by go-swagger; DO NOT EDIT.

package schools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/go-clever/models"
)

// GetSchoolReader is a Reader for the GetSchool structure.
type GetSchoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSchoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchoolOK creates a GetSchoolOK with default headers values
func NewGetSchoolOK() *GetSchoolOK {
	return &GetSchoolOK{}
}

/* GetSchoolOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSchoolOK struct {
	Payload *models.SchoolResponse
}

func (o *GetSchoolOK) Error() string {
	return fmt.Sprintf("[GET /schools/{id}][%d] getSchoolOK  %+v", 200, o.Payload)
}
func (o *GetSchoolOK) GetPayload() *models.SchoolResponse {
	return o.Payload
}

func (o *GetSchoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchoolResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchoolNotFound creates a GetSchoolNotFound with default headers values
func NewGetSchoolNotFound() *GetSchoolNotFound {
	return &GetSchoolNotFound{}
}

/* GetSchoolNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetSchoolNotFound struct {
	Payload *models.NotFound
}

func (o *GetSchoolNotFound) Error() string {
	return fmt.Sprintf("[GET /schools/{id}][%d] getSchoolNotFound  %+v", 404, o.Payload)
}
func (o *GetSchoolNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSchoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

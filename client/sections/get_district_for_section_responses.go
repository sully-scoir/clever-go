// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/go-clever/models"
)

// GetDistrictForSectionReader is a Reader for the GetDistrictForSection structure.
type GetDistrictForSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistrictForSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistrictForSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDistrictForSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistrictForSectionOK creates a GetDistrictForSectionOK with default headers values
func NewGetDistrictForSectionOK() *GetDistrictForSectionOK {
	return &GetDistrictForSectionOK{}
}

/* GetDistrictForSectionOK describes a response with status code 200, with default header values.

OK Response
*/
type GetDistrictForSectionOK struct {
	Payload *models.DistrictResponse
}

func (o *GetDistrictForSectionOK) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/district][%d] getDistrictForSectionOK  %+v", 200, o.Payload)
}
func (o *GetDistrictForSectionOK) GetPayload() *models.DistrictResponse {
	return o.Payload
}

func (o *GetDistrictForSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistrictResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistrictForSectionNotFound creates a GetDistrictForSectionNotFound with default headers values
func NewGetDistrictForSectionNotFound() *GetDistrictForSectionNotFound {
	return &GetDistrictForSectionNotFound{}
}

/* GetDistrictForSectionNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetDistrictForSectionNotFound struct {
	Payload *models.NotFound
}

func (o *GetDistrictForSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/district][%d] getDistrictForSectionNotFound  %+v", 404, o.Payload)
}
func (o *GetDistrictForSectionNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetDistrictForSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
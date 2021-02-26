// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/go-clever/models"
)

// GetSectionsForCourseReader is a Reader for the GetSectionsForCourse structure.
type GetSectionsForCourseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSectionsForCourseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSectionsForCourseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSectionsForCourseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSectionsForCourseOK creates a GetSectionsForCourseOK with default headers values
func NewGetSectionsForCourseOK() *GetSectionsForCourseOK {
	return &GetSectionsForCourseOK{}
}

/* GetSectionsForCourseOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSectionsForCourseOK struct {
	Payload *models.SectionsResponse
}

func (o *GetSectionsForCourseOK) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/sections][%d] getSectionsForCourseOK  %+v", 200, o.Payload)
}
func (o *GetSectionsForCourseOK) GetPayload() *models.SectionsResponse {
	return o.Payload
}

func (o *GetSectionsForCourseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSectionsForCourseNotFound creates a GetSectionsForCourseNotFound with default headers values
func NewGetSectionsForCourseNotFound() *GetSectionsForCourseNotFound {
	return &GetSectionsForCourseNotFound{}
}

/* GetSectionsForCourseNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetSectionsForCourseNotFound struct {
	Payload *models.NotFound
}

func (o *GetSectionsForCourseNotFound) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/sections][%d] getSectionsForCourseNotFound  %+v", 404, o.Payload)
}
func (o *GetSectionsForCourseNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSectionsForCourseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

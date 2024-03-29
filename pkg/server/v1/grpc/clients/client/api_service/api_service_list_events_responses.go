// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/clients/models"
)

// APIServiceListEventsReader is a Reader for the APIServiceListEvents structure.
type APIServiceListEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceListEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceListEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceListEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceListEventsOK creates a APIServiceListEventsOK with default headers values
func NewAPIServiceListEventsOK() *APIServiceListEventsOK {
	return &APIServiceListEventsOK{}
}

/*APIServiceListEventsOK handles this case with default header values.

A successful response.
*/
type APIServiceListEventsOK struct {
	Payload *models.GrpcListEventsResponse
}

func (o *APIServiceListEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] apiServiceListEventsOK  %+v", 200, o.Payload)
}

func (o *APIServiceListEventsOK) GetPayload() *models.GrpcListEventsResponse {
	return o.Payload
}

func (o *APIServiceListEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcListEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceListEventsDefault creates a APIServiceListEventsDefault with default headers values
func NewAPIServiceListEventsDefault(code int) *APIServiceListEventsDefault {
	return &APIServiceListEventsDefault{
		_statusCode: code,
	}
}

/*APIServiceListEventsDefault handles this case with default header values.

An unexpected error response.
*/
type APIServiceListEventsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service list events default response
func (o *APIServiceListEventsDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceListEventsDefault) Error() string {
	return fmt.Sprintf("[GET /events][%d] ApiService_ListEvents default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceListEventsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceListEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// APIServiceDeleteEventReader is a Reader for the APIServiceDeleteEvent structure.
type APIServiceDeleteEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceDeleteEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceDeleteEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceDeleteEventDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceDeleteEventOK creates a APIServiceDeleteEventOK with default headers values
func NewAPIServiceDeleteEventOK() *APIServiceDeleteEventOK {
	return &APIServiceDeleteEventOK{}
}

/*APIServiceDeleteEventOK handles this case with default header values.

A successful response.
*/
type APIServiceDeleteEventOK struct {
	Payload *models.GrpcDeleteEventResponse
}

func (o *APIServiceDeleteEventOK) Error() string {
	return fmt.Sprintf("[DELETE /event/{eventId}][%d] apiServiceDeleteEventOK  %+v", 200, o.Payload)
}

func (o *APIServiceDeleteEventOK) GetPayload() *models.GrpcDeleteEventResponse {
	return o.Payload
}

func (o *APIServiceDeleteEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcDeleteEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceDeleteEventDefault creates a APIServiceDeleteEventDefault with default headers values
func NewAPIServiceDeleteEventDefault(code int) *APIServiceDeleteEventDefault {
	return &APIServiceDeleteEventDefault{
		_statusCode: code,
	}
}

/*APIServiceDeleteEventDefault handles this case with default header values.

An unexpected error response.
*/
type APIServiceDeleteEventDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service delete event default response
func (o *APIServiceDeleteEventDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceDeleteEventDefault) Error() string {
	return fmt.Sprintf("[DELETE /event/{eventId}][%d] ApiService_DeleteEvent default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceDeleteEventDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceDeleteEventDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

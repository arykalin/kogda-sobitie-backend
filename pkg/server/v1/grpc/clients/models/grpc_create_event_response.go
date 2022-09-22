// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GrpcCreateEventResponse grpc create event response
//
// swagger:model grpcCreateEventResponse
type GrpcCreateEventResponse struct {

	// event
	Event *GrpcEvent `json:"event,omitempty"`
}

// Validate validates this grpc create event response
func (m *GrpcCreateEventResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GrpcCreateEventResponse) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GrpcCreateEventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrpcCreateEventResponse) UnmarshalBinary(b []byte) error {
	var res GrpcCreateEventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

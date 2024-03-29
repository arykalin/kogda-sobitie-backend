// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	APIServiceAuthenticate(params *APIServiceAuthenticateParams) (*APIServiceAuthenticateOK, error)

	APIServiceCreateEvent(params *APIServiceCreateEventParams) (*APIServiceCreateEventOK, error)

	APIServiceDeleteEvent(params *APIServiceDeleteEventParams) (*APIServiceDeleteEventOK, error)

	APIServiceGetEvent(params *APIServiceGetEventParams) (*APIServiceGetEventOK, error)

	APIServiceListEvents(params *APIServiceListEventsParams) (*APIServiceListEventsOK, error)

	APIServiceUpdateEvent(params *APIServiceUpdateEventParams) (*APIServiceUpdateEventOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  APIServiceAuthenticate Api service authenticate API
*/
func (a *Client) APIServiceAuthenticate(params *APIServiceAuthenticateParams) (*APIServiceAuthenticateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIServiceAuthenticateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiService_Authenticate",
		Method:             "POST",
		PathPattern:        "/authenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIServiceAuthenticateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIServiceAuthenticateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIServiceAuthenticateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  APIServiceCreateEvent Api service create event API
*/
func (a *Client) APIServiceCreateEvent(params *APIServiceCreateEventParams) (*APIServiceCreateEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIServiceCreateEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiService_CreateEvent",
		Method:             "POST",
		PathPattern:        "/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIServiceCreateEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIServiceCreateEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIServiceCreateEventDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  APIServiceDeleteEvent Api service delete event API
*/
func (a *Client) APIServiceDeleteEvent(params *APIServiceDeleteEventParams) (*APIServiceDeleteEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIServiceDeleteEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiService_DeleteEvent",
		Method:             "DELETE",
		PathPattern:        "/event/{eventId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIServiceDeleteEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIServiceDeleteEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIServiceDeleteEventDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  APIServiceGetEvent Api service get event API
*/
func (a *Client) APIServiceGetEvent(params *APIServiceGetEventParams) (*APIServiceGetEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIServiceGetEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiService_GetEvent",
		Method:             "GET",
		PathPattern:        "/event/{eventId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIServiceGetEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIServiceGetEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIServiceGetEventDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  APIServiceListEvents Api service list events API
*/
func (a *Client) APIServiceListEvents(params *APIServiceListEventsParams) (*APIServiceListEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIServiceListEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiService_ListEvents",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIServiceListEventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIServiceListEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIServiceListEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  APIServiceUpdateEvent Api service update event API
*/
func (a *Client) APIServiceUpdateEvent(params *APIServiceUpdateEventParams) (*APIServiceUpdateEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIServiceUpdateEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiService_UpdateEvent",
		Method:             "PUT",
		PathPattern:        "/event/{eventId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIServiceUpdateEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIServiceUpdateEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*APIServiceUpdateEventDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

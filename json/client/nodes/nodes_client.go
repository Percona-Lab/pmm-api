// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new nodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddBareMetal adds bare metal adds bare metal node
*/
func (a *Client) AddBareMetal(params *AddBareMetalParams) (*AddBareMetalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddBareMetalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddBareMetal",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Nodes/AddBareMetal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddBareMetalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddBareMetalOK), nil

}

/*
Get gets returns a single node by ID
*/
func (a *Client) Get(params *GetParams) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Nodes/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOK), nil

}

/*
List lists returns a list of all nodes
*/
func (a *Client) List(params *ListParams) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "List",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Nodes/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new agents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddMySqldExporterAgent adds my sqld exporter agent adds mysqld exporter agent
*/
func (a *Client) AddMySqldExporterAgent(params *AddMySqldExporterAgentParams) (*AddMySqldExporterAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySqldExporterAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQLdExporterAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/AddMySQLdExporterAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySqldExporterAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddMySqldExporterAgentOK), nil

}

/*
GetAgent gets agent returns a single agent by ID
*/
func (a *Client) GetAgent(params *GetAgentParams) (*GetAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/GetAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAgentOK), nil

}

/*
ListAgents lists agents returns a list of all agents
*/
func (a *Client) ListAgents(params *ListAgentsParams) (*ListAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAgents",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/ListAgents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAgentsOK), nil

}

/*
RemoveAgent removes agent removes agent
*/
func (a *Client) RemoveAgent(params *RemoveAgentParams) (*RemoveAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveAgent",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Agents/RemoveAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveAgentOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Percona-Lab/pmm-api/json/models"
)

// NewRemoveAgentParams creates a new RemoveAgentParams object
// with the default values initialized.
func NewRemoveAgentParams() *RemoveAgentParams {
	var ()
	return &RemoveAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveAgentParamsWithTimeout creates a new RemoveAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveAgentParamsWithTimeout(timeout time.Duration) *RemoveAgentParams {
	var ()
	return &RemoveAgentParams{

		timeout: timeout,
	}
}

// NewRemoveAgentParamsWithContext creates a new RemoveAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveAgentParamsWithContext(ctx context.Context) *RemoveAgentParams {
	var ()
	return &RemoveAgentParams{

		Context: ctx,
	}
}

// NewRemoveAgentParamsWithHTTPClient creates a new RemoveAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveAgentParamsWithHTTPClient(client *http.Client) *RemoveAgentParams {
	var ()
	return &RemoveAgentParams{
		HTTPClient: client,
	}
}

/*RemoveAgentParams contains all the parameters to send to the API endpoint
for the remove agent operation typically these are written to a http.Request
*/
type RemoveAgentParams struct {

	/*Body*/
	Body *models.InventoryRemoveAgentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove agent params
func (o *RemoveAgentParams) WithTimeout(timeout time.Duration) *RemoveAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove agent params
func (o *RemoveAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove agent params
func (o *RemoveAgentParams) WithContext(ctx context.Context) *RemoveAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove agent params
func (o *RemoveAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove agent params
func (o *RemoveAgentParams) WithHTTPClient(client *http.Client) *RemoveAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove agent params
func (o *RemoveAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove agent params
func (o *RemoveAgentParams) WithBody(body *models.InventoryRemoveAgentRequest) *RemoveAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove agent params
func (o *RemoveAgentParams) SetBody(body *models.InventoryRemoveAgentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

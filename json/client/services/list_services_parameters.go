// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewListServicesParams creates a new ListServicesParams object
// with the default values initialized.
func NewListServicesParams() *ListServicesParams {
	var ()
	return &ListServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServicesParamsWithTimeout creates a new ListServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServicesParamsWithTimeout(timeout time.Duration) *ListServicesParams {
	var ()
	return &ListServicesParams{

		timeout: timeout,
	}
}

// NewListServicesParamsWithContext creates a new ListServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServicesParamsWithContext(ctx context.Context) *ListServicesParams {
	var ()
	return &ListServicesParams{

		Context: ctx,
	}
}

// NewListServicesParamsWithHTTPClient creates a new ListServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServicesParamsWithHTTPClient(client *http.Client) *ListServicesParams {
	var ()
	return &ListServicesParams{
		HTTPClient: client,
	}
}

/*ListServicesParams contains all the parameters to send to the API endpoint
for the list services operation typically these are written to a http.Request
*/
type ListServicesParams struct {

	/*Body*/
	Body models.InventoryListServicesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list services params
func (o *ListServicesParams) WithTimeout(timeout time.Duration) *ListServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list services params
func (o *ListServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list services params
func (o *ListServicesParams) WithContext(ctx context.Context) *ListServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list services params
func (o *ListServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list services params
func (o *ListServicesParams) WithHTTPClient(client *http.Client) *ListServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list services params
func (o *ListServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list services params
func (o *ListServicesParams) WithBody(body models.InventoryListServicesRequest) *ListServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list services params
func (o *ListServicesParams) SetBody(body models.InventoryListServicesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

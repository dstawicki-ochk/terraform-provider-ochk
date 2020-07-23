// Code generated by go-swagger; DO NOT EDIT.

package routers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRouterListUsingGETParams creates a new RouterListUsingGETParams object
// with the default values initialized.
func NewRouterListUsingGETParams() *RouterListUsingGETParams {

	return &RouterListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRouterListUsingGETParamsWithTimeout creates a new RouterListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRouterListUsingGETParamsWithTimeout(timeout time.Duration) *RouterListUsingGETParams {

	return &RouterListUsingGETParams{

		timeout: timeout,
	}
}

// NewRouterListUsingGETParamsWithContext creates a new RouterListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewRouterListUsingGETParamsWithContext(ctx context.Context) *RouterListUsingGETParams {

	return &RouterListUsingGETParams{

		Context: ctx,
	}
}

// NewRouterListUsingGETParamsWithHTTPClient creates a new RouterListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRouterListUsingGETParamsWithHTTPClient(client *http.Client) *RouterListUsingGETParams {

	return &RouterListUsingGETParams{
		HTTPClient: client,
	}
}

/*RouterListUsingGETParams contains all the parameters to send to the API endpoint
for the router list using g e t operation typically these are written to a http.Request
*/
type RouterListUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the router list using g e t params
func (o *RouterListUsingGETParams) WithTimeout(timeout time.Duration) *RouterListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the router list using g e t params
func (o *RouterListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the router list using g e t params
func (o *RouterListUsingGETParams) WithContext(ctx context.Context) *RouterListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the router list using g e t params
func (o *RouterListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the router list using g e t params
func (o *RouterListUsingGETParams) WithHTTPClient(client *http.Client) *RouterListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the router list using g e t params
func (o *RouterListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RouterListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

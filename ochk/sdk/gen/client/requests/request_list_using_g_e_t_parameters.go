// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewRequestListUsingGETParams creates a new RequestListUsingGETParams object
// with the default values initialized.
func NewRequestListUsingGETParams() *RequestListUsingGETParams {

	return &RequestListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestListUsingGETParamsWithTimeout creates a new RequestListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestListUsingGETParamsWithTimeout(timeout time.Duration) *RequestListUsingGETParams {

	return &RequestListUsingGETParams{

		timeout: timeout,
	}
}

// NewRequestListUsingGETParamsWithContext creates a new RequestListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestListUsingGETParamsWithContext(ctx context.Context) *RequestListUsingGETParams {

	return &RequestListUsingGETParams{

		Context: ctx,
	}
}

// NewRequestListUsingGETParamsWithHTTPClient creates a new RequestListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestListUsingGETParamsWithHTTPClient(client *http.Client) *RequestListUsingGETParams {

	return &RequestListUsingGETParams{
		HTTPClient: client,
	}
}

/*RequestListUsingGETParams contains all the parameters to send to the API endpoint
for the request list using g e t operation typically these are written to a http.Request
*/
type RequestListUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request list using g e t params
func (o *RequestListUsingGETParams) WithTimeout(timeout time.Duration) *RequestListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request list using g e t params
func (o *RequestListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request list using g e t params
func (o *RequestListUsingGETParams) WithContext(ctx context.Context) *RequestListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request list using g e t params
func (o *RequestListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request list using g e t params
func (o *RequestListUsingGETParams) WithHTTPClient(client *http.Client) *RequestListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request list using g e t params
func (o *RequestListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RequestListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

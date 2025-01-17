// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGroupListUsingGETParams creates a new GroupListUsingGETParams object
// with the default values initialized.
func NewGroupListUsingGETParams() *GroupListUsingGETParams {
	var ()
	return &GroupListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupListUsingGETParamsWithTimeout creates a new GroupListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupListUsingGETParamsWithTimeout(timeout time.Duration) *GroupListUsingGETParams {
	var ()
	return &GroupListUsingGETParams{

		timeout: timeout,
	}
}

// NewGroupListUsingGETParamsWithContext creates a new GroupListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupListUsingGETParamsWithContext(ctx context.Context) *GroupListUsingGETParams {
	var ()
	return &GroupListUsingGETParams{

		Context: ctx,
	}
}

// NewGroupListUsingGETParamsWithHTTPClient creates a new GroupListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupListUsingGETParamsWithHTTPClient(client *http.Client) *GroupListUsingGETParams {
	var ()
	return &GroupListUsingGETParams{
		HTTPClient: client,
	}
}

/*GroupListUsingGETParams contains all the parameters to send to the API endpoint
for the group list using g e t operation typically these are written to a http.Request
*/
type GroupListUsingGETParams struct {

	/*Name
	  name

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group list using g e t params
func (o *GroupListUsingGETParams) WithTimeout(timeout time.Duration) *GroupListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group list using g e t params
func (o *GroupListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group list using g e t params
func (o *GroupListUsingGETParams) WithContext(ctx context.Context) *GroupListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group list using g e t params
func (o *GroupListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group list using g e t params
func (o *GroupListUsingGETParams) WithHTTPClient(client *http.Client) *GroupListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group list using g e t params
func (o *GroupListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the group list using g e t params
func (o *GroupListUsingGETParams) WithName(name *string) *GroupListUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the group list using g e t params
func (o *GroupListUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GroupListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

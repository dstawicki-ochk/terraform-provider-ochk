// Code generated by go-swagger; DO NOT EDIT.

package edge_clusters

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

// NewLogicalPortListUsingGETParams creates a new LogicalPortListUsingGETParams object
// with the default values initialized.
func NewLogicalPortListUsingGETParams() *LogicalPortListUsingGETParams {
	var ()
	return &LogicalPortListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogicalPortListUsingGETParamsWithTimeout creates a new LogicalPortListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogicalPortListUsingGETParamsWithTimeout(timeout time.Duration) *LogicalPortListUsingGETParams {
	var ()
	return &LogicalPortListUsingGETParams{

		timeout: timeout,
	}
}

// NewLogicalPortListUsingGETParamsWithContext creates a new LogicalPortListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogicalPortListUsingGETParamsWithContext(ctx context.Context) *LogicalPortListUsingGETParams {
	var ()
	return &LogicalPortListUsingGETParams{

		Context: ctx,
	}
}

// NewLogicalPortListUsingGETParamsWithHTTPClient creates a new LogicalPortListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogicalPortListUsingGETParamsWithHTTPClient(client *http.Client) *LogicalPortListUsingGETParams {
	var ()
	return &LogicalPortListUsingGETParams{
		HTTPClient: client,
	}
}

/*LogicalPortListUsingGETParams contains all the parameters to send to the API endpoint
for the logical port list using g e t operation typically these are written to a http.Request
*/
type LogicalPortListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) WithTimeout(timeout time.Duration) *LogicalPortListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) WithContext(ctx context.Context) *LogicalPortListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) WithHTTPClient(client *http.Client) *LogicalPortListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) WithDisplayName(displayName *string) *LogicalPortListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the logical port list using g e t params
func (o *LogicalPortListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *LogicalPortListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string
		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {
			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package ip_collection

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

// NewIPCollectionListUsingGETParams creates a new IPCollectionListUsingGETParams object
// with the default values initialized.
func NewIPCollectionListUsingGETParams() *IPCollectionListUsingGETParams {
	var ()
	return &IPCollectionListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPCollectionListUsingGETParamsWithTimeout creates a new IPCollectionListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPCollectionListUsingGETParamsWithTimeout(timeout time.Duration) *IPCollectionListUsingGETParams {
	var ()
	return &IPCollectionListUsingGETParams{

		timeout: timeout,
	}
}

// NewIPCollectionListUsingGETParamsWithContext creates a new IPCollectionListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPCollectionListUsingGETParamsWithContext(ctx context.Context) *IPCollectionListUsingGETParams {
	var ()
	return &IPCollectionListUsingGETParams{

		Context: ctx,
	}
}

// NewIPCollectionListUsingGETParamsWithHTTPClient creates a new IPCollectionListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPCollectionListUsingGETParamsWithHTTPClient(client *http.Client) *IPCollectionListUsingGETParams {
	var ()
	return &IPCollectionListUsingGETParams{
		HTTPClient: client,
	}
}

/*IPCollectionListUsingGETParams contains all the parameters to send to the API endpoint
for the Ip collection list using g e t operation typically these are written to a http.Request
*/
type IPCollectionListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) WithTimeout(timeout time.Duration) *IPCollectionListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) WithContext(ctx context.Context) *IPCollectionListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) WithHTTPClient(client *http.Client) *IPCollectionListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) WithDisplayName(displayName *string) *IPCollectionListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the Ip collection list using g e t params
func (o *IPCollectionListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *IPCollectionListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

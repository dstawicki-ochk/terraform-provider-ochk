// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// NewSnapshotListUsingGETParams creates a new SnapshotListUsingGETParams object
// with the default values initialized.
func NewSnapshotListUsingGETParams() *SnapshotListUsingGETParams {
	var ()
	return &SnapshotListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotListUsingGETParamsWithTimeout creates a new SnapshotListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSnapshotListUsingGETParamsWithTimeout(timeout time.Duration) *SnapshotListUsingGETParams {
	var ()
	return &SnapshotListUsingGETParams{

		timeout: timeout,
	}
}

// NewSnapshotListUsingGETParamsWithContext creates a new SnapshotListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSnapshotListUsingGETParamsWithContext(ctx context.Context) *SnapshotListUsingGETParams {
	var ()
	return &SnapshotListUsingGETParams{

		Context: ctx,
	}
}

// NewSnapshotListUsingGETParamsWithHTTPClient creates a new SnapshotListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSnapshotListUsingGETParamsWithHTTPClient(client *http.Client) *SnapshotListUsingGETParams {
	var ()
	return &SnapshotListUsingGETParams{
		HTTPClient: client,
	}
}

/*SnapshotListUsingGETParams contains all the parameters to send to the API endpoint
for the snapshot list using g e t operation typically these are written to a http.Request
*/
type SnapshotListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) WithTimeout(timeout time.Duration) *SnapshotListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) WithContext(ctx context.Context) *SnapshotListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) WithHTTPClient(client *http.Client) *SnapshotListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) WithDisplayName(displayName *string) *SnapshotListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the snapshot list using g e t params
func (o *SnapshotListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

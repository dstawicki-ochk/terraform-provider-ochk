// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_e_w

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

// NewDfwRuleListUsingGETParams creates a new DfwRuleListUsingGETParams object
// with the default values initialized.
func NewDfwRuleListUsingGETParams() *DfwRuleListUsingGETParams {
	var ()
	return &DfwRuleListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDfwRuleListUsingGETParamsWithTimeout creates a new DfwRuleListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDfwRuleListUsingGETParamsWithTimeout(timeout time.Duration) *DfwRuleListUsingGETParams {
	var ()
	return &DfwRuleListUsingGETParams{

		timeout: timeout,
	}
}

// NewDfwRuleListUsingGETParamsWithContext creates a new DfwRuleListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewDfwRuleListUsingGETParamsWithContext(ctx context.Context) *DfwRuleListUsingGETParams {
	var ()
	return &DfwRuleListUsingGETParams{

		Context: ctx,
	}
}

// NewDfwRuleListUsingGETParamsWithHTTPClient creates a new DfwRuleListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDfwRuleListUsingGETParamsWithHTTPClient(client *http.Client) *DfwRuleListUsingGETParams {
	var ()
	return &DfwRuleListUsingGETParams{
		HTTPClient: client,
	}
}

/*DfwRuleListUsingGETParams contains all the parameters to send to the API endpoint
for the dfw rule list using g e t operation typically these are written to a http.Request
*/
type DfwRuleListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string
	/*SecurityPolicyID
	  securityPolicyId

	*/
	SecurityPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) WithTimeout(timeout time.Duration) *DfwRuleListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) WithContext(ctx context.Context) *DfwRuleListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) WithHTTPClient(client *http.Client) *DfwRuleListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) WithDisplayName(displayName *string) *DfwRuleListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithSecurityPolicyID adds the securityPolicyID to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) WithSecurityPolicyID(securityPolicyID string) *DfwRuleListUsingGETParams {
	o.SetSecurityPolicyID(securityPolicyID)
	return o
}

// SetSecurityPolicyID adds the securityPolicyId to the dfw rule list using g e t params
func (o *DfwRuleListUsingGETParams) SetSecurityPolicyID(securityPolicyID string) {
	o.SecurityPolicyID = securityPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *DfwRuleListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param securityPolicyId
	if err := r.SetPathParam("securityPolicyId", o.SecurityPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
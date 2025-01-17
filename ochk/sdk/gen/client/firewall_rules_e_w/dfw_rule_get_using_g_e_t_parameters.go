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

// NewDfwRuleGetUsingGETParams creates a new DfwRuleGetUsingGETParams object
// with the default values initialized.
func NewDfwRuleGetUsingGETParams() *DfwRuleGetUsingGETParams {
	var ()
	return &DfwRuleGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDfwRuleGetUsingGETParamsWithTimeout creates a new DfwRuleGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDfwRuleGetUsingGETParamsWithTimeout(timeout time.Duration) *DfwRuleGetUsingGETParams {
	var ()
	return &DfwRuleGetUsingGETParams{

		timeout: timeout,
	}
}

// NewDfwRuleGetUsingGETParamsWithContext creates a new DfwRuleGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewDfwRuleGetUsingGETParamsWithContext(ctx context.Context) *DfwRuleGetUsingGETParams {
	var ()
	return &DfwRuleGetUsingGETParams{

		Context: ctx,
	}
}

// NewDfwRuleGetUsingGETParamsWithHTTPClient creates a new DfwRuleGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDfwRuleGetUsingGETParamsWithHTTPClient(client *http.Client) *DfwRuleGetUsingGETParams {
	var ()
	return &DfwRuleGetUsingGETParams{
		HTTPClient: client,
	}
}

/*DfwRuleGetUsingGETParams contains all the parameters to send to the API endpoint
for the dfw rule get using g e t operation typically these are written to a http.Request
*/
type DfwRuleGetUsingGETParams struct {

	/*RouterID
	  routerId

	*/
	RouterID string
	/*RuleID
	  ruleId

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) WithTimeout(timeout time.Duration) *DfwRuleGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) WithContext(ctx context.Context) *DfwRuleGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) WithHTTPClient(client *http.Client) *DfwRuleGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRouterID adds the routerID to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) WithRouterID(routerID string) *DfwRuleGetUsingGETParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) SetRouterID(routerID string) {
	o.RouterID = routerID
}

// WithRuleID adds the ruleID to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) WithRuleID(ruleID string) *DfwRuleGetUsingGETParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the dfw rule get using g e t params
func (o *DfwRuleGetUsingGETParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DfwRuleGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param routerId
	if err := r.SetPathParam("routerId", o.RouterID); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

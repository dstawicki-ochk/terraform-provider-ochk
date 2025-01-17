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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewDfwRuleUpdateUsingPUTParams creates a new DfwRuleUpdateUsingPUTParams object
// with the default values initialized.
func NewDfwRuleUpdateUsingPUTParams() *DfwRuleUpdateUsingPUTParams {
	var ()
	return &DfwRuleUpdateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDfwRuleUpdateUsingPUTParamsWithTimeout creates a new DfwRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDfwRuleUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *DfwRuleUpdateUsingPUTParams {
	var ()
	return &DfwRuleUpdateUsingPUTParams{

		timeout: timeout,
	}
}

// NewDfwRuleUpdateUsingPUTParamsWithContext creates a new DfwRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewDfwRuleUpdateUsingPUTParamsWithContext(ctx context.Context) *DfwRuleUpdateUsingPUTParams {
	var ()
	return &DfwRuleUpdateUsingPUTParams{

		Context: ctx,
	}
}

// NewDfwRuleUpdateUsingPUTParamsWithHTTPClient creates a new DfwRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDfwRuleUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *DfwRuleUpdateUsingPUTParams {
	var ()
	return &DfwRuleUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*DfwRuleUpdateUsingPUTParams contains all the parameters to send to the API endpoint
for the dfw rule update using p u t operation typically these are written to a http.Request
*/
type DfwRuleUpdateUsingPUTParams struct {

	/*DfwRule
	  dfwRule

	*/
	DfwRule *models.DFWRule
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

// WithTimeout adds the timeout to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *DfwRuleUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) WithContext(ctx context.Context) *DfwRuleUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *DfwRuleUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDfwRule adds the dfwRule to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) WithDfwRule(dfwRule *models.DFWRule) *DfwRuleUpdateUsingPUTParams {
	o.SetDfwRule(dfwRule)
	return o
}

// SetDfwRule adds the dfwRule to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) SetDfwRule(dfwRule *models.DFWRule) {
	o.DfwRule = dfwRule
}

// WithRouterID adds the routerID to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) WithRouterID(routerID string) *DfwRuleUpdateUsingPUTParams {
	o.SetRouterID(routerID)
	return o
}

// SetRouterID adds the routerId to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) SetRouterID(routerID string) {
	o.RouterID = routerID
}

// WithRuleID adds the ruleID to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) WithRuleID(ruleID string) *DfwRuleUpdateUsingPUTParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the dfw rule update using p u t params
func (o *DfwRuleUpdateUsingPUTParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DfwRuleUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DfwRule != nil {
		if err := r.SetBodyParam(o.DfwRule); err != nil {
			return err
		}
	}

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

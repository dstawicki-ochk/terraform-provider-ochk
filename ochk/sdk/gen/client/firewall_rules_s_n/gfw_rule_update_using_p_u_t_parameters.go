// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_s_n

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

// NewGfwRuleUpdateUsingPUTParams creates a new GfwRuleUpdateUsingPUTParams object
// with the default values initialized.
func NewGfwRuleUpdateUsingPUTParams() *GfwRuleUpdateUsingPUTParams {
	var ()
	return &GfwRuleUpdateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGfwRuleUpdateUsingPUTParamsWithTimeout creates a new GfwRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGfwRuleUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *GfwRuleUpdateUsingPUTParams {
	var ()
	return &GfwRuleUpdateUsingPUTParams{

		timeout: timeout,
	}
}

// NewGfwRuleUpdateUsingPUTParamsWithContext creates a new GfwRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGfwRuleUpdateUsingPUTParamsWithContext(ctx context.Context) *GfwRuleUpdateUsingPUTParams {
	var ()
	return &GfwRuleUpdateUsingPUTParams{

		Context: ctx,
	}
}

// NewGfwRuleUpdateUsingPUTParamsWithHTTPClient creates a new GfwRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGfwRuleUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *GfwRuleUpdateUsingPUTParams {
	var ()
	return &GfwRuleUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*GfwRuleUpdateUsingPUTParams contains all the parameters to send to the API endpoint
for the gfw rule update using p u t operation typically these are written to a http.Request
*/
type GfwRuleUpdateUsingPUTParams struct {

	/*GatewayPolicyID
	  gatewayPolicyId

	*/
	GatewayPolicyID string
	/*GfwRule
	  gfwRule

	*/
	GfwRule *models.GFWRule
	/*RuleID
	  ruleId

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *GfwRuleUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) WithContext(ctx context.Context) *GfwRuleUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *GfwRuleUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayPolicyID adds the gatewayPolicyID to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) WithGatewayPolicyID(gatewayPolicyID string) *GfwRuleUpdateUsingPUTParams {
	o.SetGatewayPolicyID(gatewayPolicyID)
	return o
}

// SetGatewayPolicyID adds the gatewayPolicyId to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) SetGatewayPolicyID(gatewayPolicyID string) {
	o.GatewayPolicyID = gatewayPolicyID
}

// WithGfwRule adds the gfwRule to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) WithGfwRule(gfwRule *models.GFWRule) *GfwRuleUpdateUsingPUTParams {
	o.SetGfwRule(gfwRule)
	return o
}

// SetGfwRule adds the gfwRule to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) SetGfwRule(gfwRule *models.GFWRule) {
	o.GfwRule = gfwRule
}

// WithRuleID adds the ruleID to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) WithRuleID(ruleID string) *GfwRuleUpdateUsingPUTParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the gfw rule update using p u t params
func (o *GfwRuleUpdateUsingPUTParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GfwRuleUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gatewayPolicyId
	if err := r.SetPathParam("gatewayPolicyId", o.GatewayPolicyID); err != nil {
		return err
	}

	if o.GfwRule != nil {
		if err := r.SetBodyParam(o.GfwRule); err != nil {
			return err
		}
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

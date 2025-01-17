// Code generated by go-swagger; DO NOT EDIT.

package subtenant_custom_groups

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

// NewGroupMemberCreateUsingPUTParams creates a new GroupMemberCreateUsingPUTParams object
// with the default values initialized.
func NewGroupMemberCreateUsingPUTParams() *GroupMemberCreateUsingPUTParams {
	var ()
	return &GroupMemberCreateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupMemberCreateUsingPUTParamsWithTimeout creates a new GroupMemberCreateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupMemberCreateUsingPUTParamsWithTimeout(timeout time.Duration) *GroupMemberCreateUsingPUTParams {
	var ()
	return &GroupMemberCreateUsingPUTParams{

		timeout: timeout,
	}
}

// NewGroupMemberCreateUsingPUTParamsWithContext creates a new GroupMemberCreateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupMemberCreateUsingPUTParamsWithContext(ctx context.Context) *GroupMemberCreateUsingPUTParams {
	var ()
	return &GroupMemberCreateUsingPUTParams{

		Context: ctx,
	}
}

// NewGroupMemberCreateUsingPUTParamsWithHTTPClient creates a new GroupMemberCreateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupMemberCreateUsingPUTParamsWithHTTPClient(client *http.Client) *GroupMemberCreateUsingPUTParams {
	var ()
	return &GroupMemberCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*GroupMemberCreateUsingPUTParams contains all the parameters to send to the API endpoint
for the group member create using p u t operation typically these are written to a http.Request
*/
type GroupMemberCreateUsingPUTParams struct {

	/*GroupInstance
	  groupInstance

	*/
	GroupInstance *models.GroupInstance
	/*ParentGroupID
	  parentGroupId

	*/
	ParentGroupID string
	/*SubtenantID
	  subtenantId

	*/
	SubtenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) WithTimeout(timeout time.Duration) *GroupMemberCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) WithContext(ctx context.Context) *GroupMemberCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) WithHTTPClient(client *http.Client) *GroupMemberCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupInstance adds the groupInstance to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) WithGroupInstance(groupInstance *models.GroupInstance) *GroupMemberCreateUsingPUTParams {
	o.SetGroupInstance(groupInstance)
	return o
}

// SetGroupInstance adds the groupInstance to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) SetGroupInstance(groupInstance *models.GroupInstance) {
	o.GroupInstance = groupInstance
}

// WithParentGroupID adds the parentGroupID to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) WithParentGroupID(parentGroupID string) *GroupMemberCreateUsingPUTParams {
	o.SetParentGroupID(parentGroupID)
	return o
}

// SetParentGroupID adds the parentGroupId to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) SetParentGroupID(parentGroupID string) {
	o.ParentGroupID = parentGroupID
}

// WithSubtenantID adds the subtenantID to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) WithSubtenantID(subtenantID string) *GroupMemberCreateUsingPUTParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the group member create using p u t params
func (o *GroupMemberCreateUsingPUTParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupMemberCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GroupInstance != nil {
		if err := r.SetBodyParam(o.GroupInstance); err != nil {
			return err
		}
	}

	// path param parentGroupId
	if err := r.SetPathParam("parentGroupId", o.ParentGroupID); err != nil {
		return err
	}

	// path param subtenantId
	if err := r.SetPathParam("subtenantId", o.SubtenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
)

// NewGroupMemberListUsingGETParams creates a new GroupMemberListUsingGETParams object
// with the default values initialized.
func NewGroupMemberListUsingGETParams() *GroupMemberListUsingGETParams {
	var ()
	return &GroupMemberListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupMemberListUsingGETParamsWithTimeout creates a new GroupMemberListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupMemberListUsingGETParamsWithTimeout(timeout time.Duration) *GroupMemberListUsingGETParams {
	var ()
	return &GroupMemberListUsingGETParams{

		timeout: timeout,
	}
}

// NewGroupMemberListUsingGETParamsWithContext creates a new GroupMemberListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupMemberListUsingGETParamsWithContext(ctx context.Context) *GroupMemberListUsingGETParams {
	var ()
	return &GroupMemberListUsingGETParams{

		Context: ctx,
	}
}

// NewGroupMemberListUsingGETParamsWithHTTPClient creates a new GroupMemberListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupMemberListUsingGETParamsWithHTTPClient(client *http.Client) *GroupMemberListUsingGETParams {
	var ()
	return &GroupMemberListUsingGETParams{
		HTTPClient: client,
	}
}

/*GroupMemberListUsingGETParams contains all the parameters to send to the API endpoint
for the group member list using g e t operation typically these are written to a http.Request
*/
type GroupMemberListUsingGETParams struct {

	/*Name
	  name

	*/
	Name *string
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

// WithTimeout adds the timeout to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) WithTimeout(timeout time.Duration) *GroupMemberListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) WithContext(ctx context.Context) *GroupMemberListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) WithHTTPClient(client *http.Client) *GroupMemberListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) WithName(name *string) *GroupMemberListUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WithParentGroupID adds the parentGroupID to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) WithParentGroupID(parentGroupID string) *GroupMemberListUsingGETParams {
	o.SetParentGroupID(parentGroupID)
	return o
}

// SetParentGroupID adds the parentGroupId to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) SetParentGroupID(parentGroupID string) {
	o.ParentGroupID = parentGroupID
}

// WithSubtenantID adds the subtenantID to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) WithSubtenantID(subtenantID string) *GroupMemberListUsingGETParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the group member list using g e t params
func (o *GroupMemberListUsingGETParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupMemberListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

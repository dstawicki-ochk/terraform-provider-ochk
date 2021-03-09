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

// NewMemberDeleteUsingDELETEParams creates a new MemberDeleteUsingDELETEParams object
// with the default values initialized.
func NewMemberDeleteUsingDELETEParams() *MemberDeleteUsingDELETEParams {
	var ()
	return &MemberDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMemberDeleteUsingDELETEParamsWithTimeout creates a new MemberDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMemberDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *MemberDeleteUsingDELETEParams {
	var ()
	return &MemberDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewMemberDeleteUsingDELETEParamsWithContext creates a new MemberDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewMemberDeleteUsingDELETEParamsWithContext(ctx context.Context) *MemberDeleteUsingDELETEParams {
	var ()
	return &MemberDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewMemberDeleteUsingDELETEParamsWithHTTPClient creates a new MemberDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMemberDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *MemberDeleteUsingDELETEParams {
	var ()
	return &MemberDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*MemberDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the member delete using d e l e t e operation typically these are written to a http.Request
*/
type MemberDeleteUsingDELETEParams struct {

	/*GroupID
	  groupId

	*/
	GroupID string
	/*SubtenantID
	  subtenantId

	*/
	SubtenantID string
	/*UserID
	  userId

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *MemberDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) WithContext(ctx context.Context) *MemberDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *MemberDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) WithGroupID(groupID string) *MemberDeleteUsingDELETEParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithSubtenantID adds the subtenantID to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) WithSubtenantID(subtenantID string) *MemberDeleteUsingDELETEParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WithUserID adds the userID to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) WithUserID(userID string) *MemberDeleteUsingDELETEParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the member delete using d e l e t e params
func (o *MemberDeleteUsingDELETEParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *MemberDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	// path param subtenantId
	if err := r.SetPathParam("subtenantId", o.SubtenantID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

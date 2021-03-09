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

// NewMemberGetUsingGETParams creates a new MemberGetUsingGETParams object
// with the default values initialized.
func NewMemberGetUsingGETParams() *MemberGetUsingGETParams {
	var ()
	return &MemberGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMemberGetUsingGETParamsWithTimeout creates a new MemberGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMemberGetUsingGETParamsWithTimeout(timeout time.Duration) *MemberGetUsingGETParams {
	var ()
	return &MemberGetUsingGETParams{

		timeout: timeout,
	}
}

// NewMemberGetUsingGETParamsWithContext creates a new MemberGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewMemberGetUsingGETParamsWithContext(ctx context.Context) *MemberGetUsingGETParams {
	var ()
	return &MemberGetUsingGETParams{

		Context: ctx,
	}
}

// NewMemberGetUsingGETParamsWithHTTPClient creates a new MemberGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMemberGetUsingGETParamsWithHTTPClient(client *http.Client) *MemberGetUsingGETParams {
	var ()
	return &MemberGetUsingGETParams{
		HTTPClient: client,
	}
}

/*MemberGetUsingGETParams contains all the parameters to send to the API endpoint
for the member get using g e t operation typically these are written to a http.Request
*/
type MemberGetUsingGETParams struct {

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

// WithTimeout adds the timeout to the member get using g e t params
func (o *MemberGetUsingGETParams) WithTimeout(timeout time.Duration) *MemberGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the member get using g e t params
func (o *MemberGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the member get using g e t params
func (o *MemberGetUsingGETParams) WithContext(ctx context.Context) *MemberGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the member get using g e t params
func (o *MemberGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the member get using g e t params
func (o *MemberGetUsingGETParams) WithHTTPClient(client *http.Client) *MemberGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the member get using g e t params
func (o *MemberGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the member get using g e t params
func (o *MemberGetUsingGETParams) WithGroupID(groupID string) *MemberGetUsingGETParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the member get using g e t params
func (o *MemberGetUsingGETParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithSubtenantID adds the subtenantID to the member get using g e t params
func (o *MemberGetUsingGETParams) WithSubtenantID(subtenantID string) *MemberGetUsingGETParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the member get using g e t params
func (o *MemberGetUsingGETParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WithUserID adds the userID to the member get using g e t params
func (o *MemberGetUsingGETParams) WithUserID(userID string) *MemberGetUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the member get using g e t params
func (o *MemberGetUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *MemberGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

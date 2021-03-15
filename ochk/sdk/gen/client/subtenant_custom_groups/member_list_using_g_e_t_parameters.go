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

// NewMemberListUsingGETParams creates a new MemberListUsingGETParams object
// with the default values initialized.
func NewMemberListUsingGETParams() *MemberListUsingGETParams {
	var ()
	return &MemberListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMemberListUsingGETParamsWithTimeout creates a new MemberListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMemberListUsingGETParamsWithTimeout(timeout time.Duration) *MemberListUsingGETParams {
	var ()
	return &MemberListUsingGETParams{

		timeout: timeout,
	}
}

// NewMemberListUsingGETParamsWithContext creates a new MemberListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewMemberListUsingGETParamsWithContext(ctx context.Context) *MemberListUsingGETParams {
	var ()
	return &MemberListUsingGETParams{

		Context: ctx,
	}
}

// NewMemberListUsingGETParamsWithHTTPClient creates a new MemberListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMemberListUsingGETParamsWithHTTPClient(client *http.Client) *MemberListUsingGETParams {
	var ()
	return &MemberListUsingGETParams{
		HTTPClient: client,
	}
}

/*MemberListUsingGETParams contains all the parameters to send to the API endpoint
for the member list using g e t operation typically these are written to a http.Request
*/
type MemberListUsingGETParams struct {

	/*GroupID
	  groupId

	*/
	GroupID string
	/*Name
	  name

	*/
	Name *string
	/*SubtenantID
	  subtenantId

	*/
	SubtenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the member list using g e t params
func (o *MemberListUsingGETParams) WithTimeout(timeout time.Duration) *MemberListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the member list using g e t params
func (o *MemberListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the member list using g e t params
func (o *MemberListUsingGETParams) WithContext(ctx context.Context) *MemberListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the member list using g e t params
func (o *MemberListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the member list using g e t params
func (o *MemberListUsingGETParams) WithHTTPClient(client *http.Client) *MemberListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the member list using g e t params
func (o *MemberListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the member list using g e t params
func (o *MemberListUsingGETParams) WithGroupID(groupID string) *MemberListUsingGETParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the member list using g e t params
func (o *MemberListUsingGETParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithName adds the name to the member list using g e t params
func (o *MemberListUsingGETParams) WithName(name *string) *MemberListUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the member list using g e t params
func (o *MemberListUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WithSubtenantID adds the subtenantID to the member list using g e t params
func (o *MemberListUsingGETParams) WithSubtenantID(subtenantID string) *MemberListUsingGETParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the member list using g e t params
func (o *MemberListUsingGETParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WriteToRequest writes these params to a swagger request
func (o *MemberListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

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

	// path param subtenantId
	if err := r.SetPathParam("subtenantId", o.SubtenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
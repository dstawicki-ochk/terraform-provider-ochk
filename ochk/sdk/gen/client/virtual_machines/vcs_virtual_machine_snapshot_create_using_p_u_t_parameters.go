// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

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
	"github.com/go-openapi/swag"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewVcsVirtualMachineSnapshotCreateUsingPUTParams creates a new VcsVirtualMachineSnapshotCreateUsingPUTParams object
// with the default values initialized.
func NewVcsVirtualMachineSnapshotCreateUsingPUTParams() *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineSnapshotCreateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVcsVirtualMachineSnapshotCreateUsingPUTParamsWithTimeout creates a new VcsVirtualMachineSnapshotCreateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVcsVirtualMachineSnapshotCreateUsingPUTParamsWithTimeout(timeout time.Duration) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineSnapshotCreateUsingPUTParams{

		timeout: timeout,
	}
}

// NewVcsVirtualMachineSnapshotCreateUsingPUTParamsWithContext creates a new VcsVirtualMachineSnapshotCreateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewVcsVirtualMachineSnapshotCreateUsingPUTParamsWithContext(ctx context.Context) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineSnapshotCreateUsingPUTParams{

		Context: ctx,
	}
}

// NewVcsVirtualMachineSnapshotCreateUsingPUTParamsWithHTTPClient creates a new VcsVirtualMachineSnapshotCreateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVcsVirtualMachineSnapshotCreateUsingPUTParamsWithHTTPClient(client *http.Client) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	var ()
	return &VcsVirtualMachineSnapshotCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*VcsVirtualMachineSnapshotCreateUsingPUTParams contains all the parameters to send to the API endpoint
for the vcs virtual machine snapshot create using p u t operation typically these are written to a http.Request
*/
type VcsVirtualMachineSnapshotCreateUsingPUTParams struct {

	/*RAMSnapshot
	  ramSnapshot

	*/
	RAMSnapshot *bool
	/*SnapshotInstance
	  snapshotInstance

	*/
	SnapshotInstance *models.SnapshotInstance
	/*VirtualMachineID
	  virtualMachineId

	*/
	VirtualMachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WithTimeout(timeout time.Duration) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WithContext(ctx context.Context) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WithHTTPClient(client *http.Client) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRAMSnapshot adds the rAMSnapshot to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WithRAMSnapshot(rAMSnapshot *bool) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	o.SetRAMSnapshot(rAMSnapshot)
	return o
}

// SetRAMSnapshot adds the ramSnapshot to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) SetRAMSnapshot(rAMSnapshot *bool) {
	o.RAMSnapshot = rAMSnapshot
}

// WithSnapshotInstance adds the snapshotInstance to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WithSnapshotInstance(snapshotInstance *models.SnapshotInstance) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	o.SetSnapshotInstance(snapshotInstance)
	return o
}

// SetSnapshotInstance adds the snapshotInstance to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) SetSnapshotInstance(snapshotInstance *models.SnapshotInstance) {
	o.SnapshotInstance = snapshotInstance
}

// WithVirtualMachineID adds the virtualMachineID to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WithVirtualMachineID(virtualMachineID string) *VcsVirtualMachineSnapshotCreateUsingPUTParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the vcs virtual machine snapshot create using p u t params
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) SetVirtualMachineID(virtualMachineID string) {
	o.VirtualMachineID = virtualMachineID
}

// WriteToRequest writes these params to a swagger request
func (o *VcsVirtualMachineSnapshotCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RAMSnapshot != nil {

		// query param ramSnapshot
		var qrRAMSnapshot bool
		if o.RAMSnapshot != nil {
			qrRAMSnapshot = *o.RAMSnapshot
		}
		qRAMSnapshot := swag.FormatBool(qrRAMSnapshot)
		if qRAMSnapshot != "" {
			if err := r.SetQueryParam("ramSnapshot", qRAMSnapshot); err != nil {
				return err
			}
		}

	}

	if o.SnapshotInstance != nil {
		if err := r.SetBodyParam(o.SnapshotInstance); err != nil {
			return err
		}
	}

	// path param virtualMachineId
	if err := r.SetPathParam("virtualMachineId", o.VirtualMachineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

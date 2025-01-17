// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VcsVirtualMachineSnapshotListUsingGETReader is a Reader for the VcsVirtualMachineSnapshotListUsingGET structure.
type VcsVirtualMachineSnapshotListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineSnapshotListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineSnapshotListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineSnapshotListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVcsVirtualMachineSnapshotListUsingGETOK creates a VcsVirtualMachineSnapshotListUsingGETOK with default headers values
func NewVcsVirtualMachineSnapshotListUsingGETOK() *VcsVirtualMachineSnapshotListUsingGETOK {
	return &VcsVirtualMachineSnapshotListUsingGETOK{}
}

/*VcsVirtualMachineSnapshotListUsingGETOK handles this case with default header values.

Request has been completed successfully
*/
type VcsVirtualMachineSnapshotListUsingGETOK struct {
	Payload *models.SnapshotListResponse
}

func (o *VcsVirtualMachineSnapshotListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}/snapshots][%d] vcsVirtualMachineSnapshotListUsingGETOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineSnapshotListUsingGETOK) GetPayload() *models.SnapshotListResponse {
	return o.Payload
}

func (o *VcsVirtualMachineSnapshotListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineSnapshotListUsingGETBadRequest creates a VcsVirtualMachineSnapshotListUsingGETBadRequest with default headers values
func NewVcsVirtualMachineSnapshotListUsingGETBadRequest() *VcsVirtualMachineSnapshotListUsingGETBadRequest {
	return &VcsVirtualMachineSnapshotListUsingGETBadRequest{}
}

/*VcsVirtualMachineSnapshotListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineSnapshotListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VcsVirtualMachineSnapshotListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}/snapshots][%d] vcsVirtualMachineSnapshotListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VcsVirtualMachineSnapshotListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VcsVirtualMachineSnapshotListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

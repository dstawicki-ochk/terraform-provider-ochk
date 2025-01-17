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

// VcsVirtualMachineListUsingGET1Reader is a Reader for the VcsVirtualMachineListUsingGET1 structure.
type VcsVirtualMachineListUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineListUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineListUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineListUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVcsVirtualMachineListUsingGET1OK creates a VcsVirtualMachineListUsingGET1OK with default headers values
func NewVcsVirtualMachineListUsingGET1OK() *VcsVirtualMachineListUsingGET1OK {
	return &VcsVirtualMachineListUsingGET1OK{}
}

/*VcsVirtualMachineListUsingGET1OK handles this case with default header values.

OK
*/
type VcsVirtualMachineListUsingGET1OK struct {
	Payload *models.VcsVirtualMachineListResponse
}

func (o *VcsVirtualMachineListUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines][%d] vcsVirtualMachineListUsingGET1OK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineListUsingGET1OK) GetPayload() *models.VcsVirtualMachineListResponse {
	return o.Payload
}

func (o *VcsVirtualMachineListUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsVirtualMachineListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineListUsingGET1BadRequest creates a VcsVirtualMachineListUsingGET1BadRequest with default headers values
func NewVcsVirtualMachineListUsingGET1BadRequest() *VcsVirtualMachineListUsingGET1BadRequest {
	return &VcsVirtualMachineListUsingGET1BadRequest{}
}

/*VcsVirtualMachineListUsingGET1BadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineListUsingGET1BadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VcsVirtualMachineListUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines][%d] vcsVirtualMachineListUsingGET1BadRequest  %+v", 400, o.Payload)
}

func (o *VcsVirtualMachineListUsingGET1BadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VcsVirtualMachineListUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

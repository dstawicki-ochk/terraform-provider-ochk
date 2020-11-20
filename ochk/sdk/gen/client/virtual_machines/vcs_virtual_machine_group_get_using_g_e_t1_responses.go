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

// VcsVirtualMachineGroupGetUsingGET1Reader is a Reader for the VcsVirtualMachineGroupGetUsingGET1 structure.
type VcsVirtualMachineGroupGetUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineGroupGetUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineGroupGetUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineGroupGetUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVcsVirtualMachineGroupGetUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVcsVirtualMachineGroupGetUsingGET1OK creates a VcsVirtualMachineGroupGetUsingGET1OK with default headers values
func NewVcsVirtualMachineGroupGetUsingGET1OK() *VcsVirtualMachineGroupGetUsingGET1OK {
	return &VcsVirtualMachineGroupGetUsingGET1OK{}
}

/*VcsVirtualMachineGroupGetUsingGET1OK handles this case with default header values.

OK
*/
type VcsVirtualMachineGroupGetUsingGET1OK struct {
	Payload *models.VcsVirtualMachineGetResponse
}

func (o *VcsVirtualMachineGroupGetUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGroupGetUsingGET1OK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineGroupGetUsingGET1OK) GetPayload() *models.VcsVirtualMachineGetResponse {
	return o.Payload
}

func (o *VcsVirtualMachineGroupGetUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsVirtualMachineGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineGroupGetUsingGET1BadRequest creates a VcsVirtualMachineGroupGetUsingGET1BadRequest with default headers values
func NewVcsVirtualMachineGroupGetUsingGET1BadRequest() *VcsVirtualMachineGroupGetUsingGET1BadRequest {
	return &VcsVirtualMachineGroupGetUsingGET1BadRequest{}
}

/*VcsVirtualMachineGroupGetUsingGET1BadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineGroupGetUsingGET1BadRequest struct {
}

func (o *VcsVirtualMachineGroupGetUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGroupGetUsingGET1BadRequest ", 400)
}

func (o *VcsVirtualMachineGroupGetUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVcsVirtualMachineGroupGetUsingGET1NotFound creates a VcsVirtualMachineGroupGetUsingGET1NotFound with default headers values
func NewVcsVirtualMachineGroupGetUsingGET1NotFound() *VcsVirtualMachineGroupGetUsingGET1NotFound {
	return &VcsVirtualMachineGroupGetUsingGET1NotFound{}
}

/*VcsVirtualMachineGroupGetUsingGET1NotFound handles this case with default header values.

Entity not found.
*/
type VcsVirtualMachineGroupGetUsingGET1NotFound struct {
}

func (o *VcsVirtualMachineGroupGetUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineGroupGetUsingGET1NotFound ", 404)
}

func (o *VcsVirtualMachineGroupGetUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

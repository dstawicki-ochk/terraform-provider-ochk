// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines_n_s_x

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VirtualMachineListUsingGETReader is a Reader for the VirtualMachineListUsingGET structure.
type VirtualMachineListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualMachineListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualMachineListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineListUsingGETOK creates a VirtualMachineListUsingGETOK with default headers values
func NewVirtualMachineListUsingGETOK() *VirtualMachineListUsingGETOK {
	return &VirtualMachineListUsingGETOK{}
}

/*VirtualMachineListUsingGETOK handles this case with default header values.

OK
*/
type VirtualMachineListUsingGETOK struct {
	Payload *models.VirtualMachineListResponse
}

func (o *VirtualMachineListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/virtual-machines][%d] virtualMachineListUsingGETOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineListUsingGETOK) GetPayload() *models.VirtualMachineListResponse {
	return o.Payload
}

func (o *VirtualMachineListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachineListUsingGETBadRequest creates a VirtualMachineListUsingGETBadRequest with default headers values
func NewVirtualMachineListUsingGETBadRequest() *VirtualMachineListUsingGETBadRequest {
	return &VirtualMachineListUsingGETBadRequest{}
}

/*VirtualMachineListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VirtualMachineListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VirtualMachineListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/virtual-machines][%d] virtualMachineListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualMachineListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VirtualMachineListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

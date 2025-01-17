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

// VirtualMachineGetUsingGETReader is a Reader for the VirtualMachineGetUsingGET structure.
type VirtualMachineGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualMachineGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualMachineGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualMachineGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVirtualMachineGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualMachineGetUsingGETOK creates a VirtualMachineGetUsingGETOK with default headers values
func NewVirtualMachineGetUsingGETOK() *VirtualMachineGetUsingGETOK {
	return &VirtualMachineGetUsingGETOK{}
}

/*VirtualMachineGetUsingGETOK handles this case with default header values.

OK
*/
type VirtualMachineGetUsingGETOK struct {
	Payload *models.VirtualMachineGetResponse
}

func (o *VirtualMachineGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/virtual-machines/{virtualMachineId}][%d] virtualMachineGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *VirtualMachineGetUsingGETOK) GetPayload() *models.VirtualMachineGetResponse {
	return o.Payload
}

func (o *VirtualMachineGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachineGetUsingGETBadRequest creates a VirtualMachineGetUsingGETBadRequest with default headers values
func NewVirtualMachineGetUsingGETBadRequest() *VirtualMachineGetUsingGETBadRequest {
	return &VirtualMachineGetUsingGETBadRequest{}
}

/*VirtualMachineGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VirtualMachineGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *VirtualMachineGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/virtual-machines/{virtualMachineId}][%d] virtualMachineGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualMachineGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *VirtualMachineGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualMachineGetUsingGETNotFound creates a VirtualMachineGetUsingGETNotFound with default headers values
func NewVirtualMachineGetUsingGETNotFound() *VirtualMachineGetUsingGETNotFound {
	return &VirtualMachineGetUsingGETNotFound{}
}

/*VirtualMachineGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type VirtualMachineGetUsingGETNotFound struct {
}

func (o *VirtualMachineGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/virtual-machines/{virtualMachineId}][%d] virtualMachineGetUsingGETNotFound ", 404)
}

func (o *VirtualMachineGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

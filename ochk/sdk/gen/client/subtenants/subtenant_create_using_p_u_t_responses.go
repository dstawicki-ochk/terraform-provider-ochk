// Code generated by go-swagger; DO NOT EDIT.

package subtenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SubtenantCreateUsingPUTReader is a Reader for the SubtenantCreateUsingPUT structure.
type SubtenantCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtenantCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtenantCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSubtenantCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubtenantCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubtenantCreateUsingPUTOK creates a SubtenantCreateUsingPUTOK with default headers values
func NewSubtenantCreateUsingPUTOK() *SubtenantCreateUsingPUTOK {
	return &SubtenantCreateUsingPUTOK{}
}

/*SubtenantCreateUsingPUTOK handles this case with default header values.

OK
*/
type SubtenantCreateUsingPUTOK struct {
	Payload *models.SubtenantCreateResponse
}

func (o *SubtenantCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /subtenants][%d] subtenantCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *SubtenantCreateUsingPUTOK) GetPayload() *models.SubtenantCreateResponse {
	return o.Payload
}

func (o *SubtenantCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantCreateUsingPUTCreated creates a SubtenantCreateUsingPUTCreated with default headers values
func NewSubtenantCreateUsingPUTCreated() *SubtenantCreateUsingPUTCreated {
	return &SubtenantCreateUsingPUTCreated{}
}

/*SubtenantCreateUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type SubtenantCreateUsingPUTCreated struct {
	Payload *models.SubtenantCreateResponse
}

func (o *SubtenantCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /subtenants][%d] subtenantCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *SubtenantCreateUsingPUTCreated) GetPayload() *models.SubtenantCreateResponse {
	return o.Payload
}

func (o *SubtenantCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantCreateUsingPUTBadRequest creates a SubtenantCreateUsingPUTBadRequest with default headers values
func NewSubtenantCreateUsingPUTBadRequest() *SubtenantCreateUsingPUTBadRequest {
	return &SubtenantCreateUsingPUTBadRequest{}
}

/*SubtenantCreateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SubtenantCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SubtenantCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subtenants][%d] subtenantCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *SubtenantCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SubtenantCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

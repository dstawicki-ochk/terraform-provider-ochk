// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_s_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GfwRuleCreateUsingPUTReader is a Reader for the GfwRuleCreateUsingPUT structure.
type GfwRuleCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GfwRuleCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGfwRuleCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewGfwRuleCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGfwRuleCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGfwRuleCreateUsingPUTOK creates a GfwRuleCreateUsingPUTOK with default headers values
func NewGfwRuleCreateUsingPUTOK() *GfwRuleCreateUsingPUTOK {
	return &GfwRuleCreateUsingPUTOK{}
}

/*GfwRuleCreateUsingPUTOK handles this case with default header values.

OK
*/
type GfwRuleCreateUsingPUTOK struct {
	Payload *models.CreateGFWRuleResponse
}

func (o *GfwRuleCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] gfwRuleCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *GfwRuleCreateUsingPUTOK) GetPayload() *models.CreateGFWRuleResponse {
	return o.Payload
}

func (o *GfwRuleCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateGFWRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGfwRuleCreateUsingPUTCreated creates a GfwRuleCreateUsingPUTCreated with default headers values
func NewGfwRuleCreateUsingPUTCreated() *GfwRuleCreateUsingPUTCreated {
	return &GfwRuleCreateUsingPUTCreated{}
}

/*GfwRuleCreateUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type GfwRuleCreateUsingPUTCreated struct {
	Payload *models.CreateGFWRuleResponse
}

func (o *GfwRuleCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] gfwRuleCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *GfwRuleCreateUsingPUTCreated) GetPayload() *models.CreateGFWRuleResponse {
	return o.Payload
}

func (o *GfwRuleCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateGFWRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGfwRuleCreateUsingPUTBadRequest creates a GfwRuleCreateUsingPUTBadRequest with default headers values
func NewGfwRuleCreateUsingPUTBadRequest() *GfwRuleCreateUsingPUTBadRequest {
	return &GfwRuleCreateUsingPUTBadRequest{}
}

/*GfwRuleCreateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GfwRuleCreateUsingPUTBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *GfwRuleCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/routers/{routerId}/rules/s-n][%d] gfwRuleCreateUsingPUTBadRequest  %+v", 400, o.Payload)
}

func (o *GfwRuleCreateUsingPUTBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *GfwRuleCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

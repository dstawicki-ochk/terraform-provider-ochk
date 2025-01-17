// Code generated by go-swagger; DO NOT EDIT.

package gateway_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GatewayPolicyListUsingGETReader is a Reader for the GatewayPolicyListUsingGET structure.
type GatewayPolicyListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GatewayPolicyListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGatewayPolicyListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGatewayPolicyListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGatewayPolicyListUsingGETOK creates a GatewayPolicyListUsingGETOK with default headers values
func NewGatewayPolicyListUsingGETOK() *GatewayPolicyListUsingGETOK {
	return &GatewayPolicyListUsingGETOK{}
}

/*GatewayPolicyListUsingGETOK handles this case with default header values.

OK
*/
type GatewayPolicyListUsingGETOK struct {
	Payload *models.GatewayPolicyListResponse
}

func (o *GatewayPolicyListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/firewall/gateway-policies][%d] gatewayPolicyListUsingGETOK  %+v", 200, o.Payload)
}

func (o *GatewayPolicyListUsingGETOK) GetPayload() *models.GatewayPolicyListResponse {
	return o.Payload
}

func (o *GatewayPolicyListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayPolicyListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGatewayPolicyListUsingGETBadRequest creates a GatewayPolicyListUsingGETBadRequest with default headers values
func NewGatewayPolicyListUsingGETBadRequest() *GatewayPolicyListUsingGETBadRequest {
	return &GatewayPolicyListUsingGETBadRequest{}
}

/*GatewayPolicyListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GatewayPolicyListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *GatewayPolicyListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/firewall/gateway-policies][%d] gatewayPolicyListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GatewayPolicyListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *GatewayPolicyListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

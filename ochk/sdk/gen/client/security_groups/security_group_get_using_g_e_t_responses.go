// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SecurityGroupGetUsingGETReader is a Reader for the SecurityGroupGetUsingGET structure.
type SecurityGroupGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityGroupGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecurityGroupGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecurityGroupGetUsingGETOK creates a SecurityGroupGetUsingGETOK with default headers values
func NewSecurityGroupGetUsingGETOK() *SecurityGroupGetUsingGETOK {
	return &SecurityGroupGetUsingGETOK{}
}

/*SecurityGroupGetUsingGETOK handles this case with default header values.

OK
*/
type SecurityGroupGetUsingGETOK struct {
	Payload *models.SecurityGroupGetResponse
}

func (o *SecurityGroupGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/security-groups/{groupId}][%d] securityGroupGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *SecurityGroupGetUsingGETOK) GetPayload() *models.SecurityGroupGetResponse {
	return o.Payload
}

func (o *SecurityGroupGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupGetUsingGETBadRequest creates a SecurityGroupGetUsingGETBadRequest with default headers values
func NewSecurityGroupGetUsingGETBadRequest() *SecurityGroupGetUsingGETBadRequest {
	return &SecurityGroupGetUsingGETBadRequest{}
}

/*SecurityGroupGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SecurityGroupGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SecurityGroupGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/security-groups/{groupId}][%d] securityGroupGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SecurityGroupGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SecurityGroupGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupGetUsingGETNotFound creates a SecurityGroupGetUsingGETNotFound with default headers values
func NewSecurityGroupGetUsingGETNotFound() *SecurityGroupGetUsingGETNotFound {
	return &SecurityGroupGetUsingGETNotFound{}
}

/*SecurityGroupGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type SecurityGroupGetUsingGETNotFound struct {
}

func (o *SecurityGroupGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/security-groups/{groupId}][%d] securityGroupGetUsingGETNotFound ", 404)
}

func (o *SecurityGroupGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_rotation_scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// KeyRotationScheduleGetUsingGETReader is a Reader for the KeyRotationScheduleGetUsingGET structure.
type KeyRotationScheduleGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyRotationScheduleGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyRotationScheduleGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyRotationScheduleGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeyRotationScheduleGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyRotationScheduleGetUsingGETOK creates a KeyRotationScheduleGetUsingGETOK with default headers values
func NewKeyRotationScheduleGetUsingGETOK() *KeyRotationScheduleGetUsingGETOK {
	return &KeyRotationScheduleGetUsingGETOK{}
}

/*KeyRotationScheduleGetUsingGETOK handles this case with default header values.

OK
*/
type KeyRotationScheduleGetUsingGETOK struct {
	Payload *models.KeyRotationScheduleGetResponse
}

func (o *KeyRotationScheduleGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /kms/schedule/{keyId}][%d] keyRotationScheduleGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *KeyRotationScheduleGetUsingGETOK) GetPayload() *models.KeyRotationScheduleGetResponse {
	return o.Payload
}

func (o *KeyRotationScheduleGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyRotationScheduleGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyRotationScheduleGetUsingGETBadRequest creates a KeyRotationScheduleGetUsingGETBadRequest with default headers values
func NewKeyRotationScheduleGetUsingGETBadRequest() *KeyRotationScheduleGetUsingGETBadRequest {
	return &KeyRotationScheduleGetUsingGETBadRequest{}
}

/*KeyRotationScheduleGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyRotationScheduleGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *KeyRotationScheduleGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /kms/schedule/{keyId}][%d] keyRotationScheduleGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *KeyRotationScheduleGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *KeyRotationScheduleGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyRotationScheduleGetUsingGETNotFound creates a KeyRotationScheduleGetUsingGETNotFound with default headers values
func NewKeyRotationScheduleGetUsingGETNotFound() *KeyRotationScheduleGetUsingGETNotFound {
	return &KeyRotationScheduleGetUsingGETNotFound{}
}

/*KeyRotationScheduleGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type KeyRotationScheduleGetUsingGETNotFound struct {
}

func (o *KeyRotationScheduleGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /kms/schedule/{keyId}][%d] keyRotationScheduleGetUsingGETNotFound ", 404)
}

func (o *KeyRotationScheduleGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

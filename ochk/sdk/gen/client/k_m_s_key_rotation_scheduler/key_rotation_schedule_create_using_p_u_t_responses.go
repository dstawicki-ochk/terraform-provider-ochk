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

// KeyRotationScheduleCreateUsingPUTReader is a Reader for the KeyRotationScheduleCreateUsingPUT structure.
type KeyRotationScheduleCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyRotationScheduleCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyRotationScheduleCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewKeyRotationScheduleCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyRotationScheduleCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyRotationScheduleCreateUsingPUTOK creates a KeyRotationScheduleCreateUsingPUTOK with default headers values
func NewKeyRotationScheduleCreateUsingPUTOK() *KeyRotationScheduleCreateUsingPUTOK {
	return &KeyRotationScheduleCreateUsingPUTOK{}
}

/*KeyRotationScheduleCreateUsingPUTOK handles this case with default header values.

OK
*/
type KeyRotationScheduleCreateUsingPUTOK struct {
	Payload *models.CreateKeyRotationScheduleResponse
}

func (o *KeyRotationScheduleCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /kms/schedule][%d] keyRotationScheduleCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *KeyRotationScheduleCreateUsingPUTOK) GetPayload() *models.CreateKeyRotationScheduleResponse {
	return o.Payload
}

func (o *KeyRotationScheduleCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateKeyRotationScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyRotationScheduleCreateUsingPUTCreated creates a KeyRotationScheduleCreateUsingPUTCreated with default headers values
func NewKeyRotationScheduleCreateUsingPUTCreated() *KeyRotationScheduleCreateUsingPUTCreated {
	return &KeyRotationScheduleCreateUsingPUTCreated{}
}

/*KeyRotationScheduleCreateUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type KeyRotationScheduleCreateUsingPUTCreated struct {
	Payload *models.CreateKeyRotationScheduleResponse
}

func (o *KeyRotationScheduleCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /kms/schedule][%d] keyRotationScheduleCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *KeyRotationScheduleCreateUsingPUTCreated) GetPayload() *models.CreateKeyRotationScheduleResponse {
	return o.Payload
}

func (o *KeyRotationScheduleCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateKeyRotationScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyRotationScheduleCreateUsingPUTBadRequest creates a KeyRotationScheduleCreateUsingPUTBadRequest with default headers values
func NewKeyRotationScheduleCreateUsingPUTBadRequest() *KeyRotationScheduleCreateUsingPUTBadRequest {
	return &KeyRotationScheduleCreateUsingPUTBadRequest{}
}

/*KeyRotationScheduleCreateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyRotationScheduleCreateUsingPUTBadRequest struct {
}

func (o *KeyRotationScheduleCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /kms/schedule][%d] keyRotationScheduleCreateUsingPUTBadRequest ", 400)
}

func (o *KeyRotationScheduleCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
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

// KeyRotationScheduleDeleteUsingDELETEReader is a Reader for the KeyRotationScheduleDeleteUsingDELETE structure.
type KeyRotationScheduleDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyRotationScheduleDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyRotationScheduleDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyRotationScheduleDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyRotationScheduleDeleteUsingDELETEOK creates a KeyRotationScheduleDeleteUsingDELETEOK with default headers values
func NewKeyRotationScheduleDeleteUsingDELETEOK() *KeyRotationScheduleDeleteUsingDELETEOK {
	return &KeyRotationScheduleDeleteUsingDELETEOK{}
}

/*KeyRotationScheduleDeleteUsingDELETEOK handles this case with default header values.

OK
*/
type KeyRotationScheduleDeleteUsingDELETEOK struct {
	Payload *models.DeleteKeyRotationScheduleResponse
}

func (o *KeyRotationScheduleDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /kms/schedule/{keyId}][%d] keyRotationScheduleDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *KeyRotationScheduleDeleteUsingDELETEOK) GetPayload() *models.DeleteKeyRotationScheduleResponse {
	return o.Payload
}

func (o *KeyRotationScheduleDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteKeyRotationScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyRotationScheduleDeleteUsingDELETEBadRequest creates a KeyRotationScheduleDeleteUsingDELETEBadRequest with default headers values
func NewKeyRotationScheduleDeleteUsingDELETEBadRequest() *KeyRotationScheduleDeleteUsingDELETEBadRequest {
	return &KeyRotationScheduleDeleteUsingDELETEBadRequest{}
}

/*KeyRotationScheduleDeleteUsingDELETEBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyRotationScheduleDeleteUsingDELETEBadRequest struct {
}

func (o *KeyRotationScheduleDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kms/schedule/{keyId}][%d] keyRotationScheduleDeleteUsingDELETEBadRequest ", 400)
}

func (o *KeyRotationScheduleDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
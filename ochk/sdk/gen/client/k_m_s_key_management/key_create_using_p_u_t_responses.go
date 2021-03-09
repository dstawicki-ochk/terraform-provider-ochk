// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// KeyCreateUsingPUTReader is a Reader for the KeyCreateUsingPUT structure.
type KeyCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewKeyCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyCreateUsingPUTOK creates a KeyCreateUsingPUTOK with default headers values
func NewKeyCreateUsingPUTOK() *KeyCreateUsingPUTOK {
	return &KeyCreateUsingPUTOK{}
}

/*KeyCreateUsingPUTOK handles this case with default header values.

OK
*/
type KeyCreateUsingPUTOK struct {
	Payload *models.CreateKmsKeyResponse
}

func (o *KeyCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *KeyCreateUsingPUTOK) GetPayload() *models.CreateKmsKeyResponse {
	return o.Payload
}

func (o *KeyCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyCreateUsingPUTCreated creates a KeyCreateUsingPUTCreated with default headers values
func NewKeyCreateUsingPUTCreated() *KeyCreateUsingPUTCreated {
	return &KeyCreateUsingPUTCreated{}
}

/*KeyCreateUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type KeyCreateUsingPUTCreated struct {
	Payload *models.CreateKmsKeyResponse
}

func (o *KeyCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *KeyCreateUsingPUTCreated) GetPayload() *models.CreateKmsKeyResponse {
	return o.Payload
}

func (o *KeyCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateKmsKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyCreateUsingPUTBadRequest creates a KeyCreateUsingPUTBadRequest with default headers values
func NewKeyCreateUsingPUTBadRequest() *KeyCreateUsingPUTBadRequest {
	return &KeyCreateUsingPUTBadRequest{}
}

/*KeyCreateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyCreateUsingPUTBadRequest struct {
}

func (o *KeyCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /kms/key][%d] keyCreateUsingPUTBadRequest ", 400)
}

func (o *KeyCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

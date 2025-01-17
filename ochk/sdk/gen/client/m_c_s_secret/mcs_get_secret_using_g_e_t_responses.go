// Code generated by go-swagger; DO NOT EDIT.

package m_c_s_secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// McsGetSecretUsingGETReader is a Reader for the McsGetSecretUsingGET structure.
type McsGetSecretUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *McsGetSecretUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMcsGetSecretUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMcsGetSecretUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMcsGetSecretUsingGETOK creates a McsGetSecretUsingGETOK with default headers values
func NewMcsGetSecretUsingGETOK() *McsGetSecretUsingGETOK {
	return &McsGetSecretUsingGETOK{}
}

/*McsGetSecretUsingGETOK handles this case with default header values.

OK
*/
type McsGetSecretUsingGETOK struct {
	Payload *models.McsSecretGetResponse
}

func (o *McsGetSecretUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /mcs/secret][%d] mcsGetSecretUsingGETOK  %+v", 200, o.Payload)
}

func (o *McsGetSecretUsingGETOK) GetPayload() *models.McsSecretGetResponse {
	return o.Payload
}

func (o *McsGetSecretUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.McsSecretGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMcsGetSecretUsingGETBadRequest creates a McsGetSecretUsingGETBadRequest with default headers values
func NewMcsGetSecretUsingGETBadRequest() *McsGetSecretUsingGETBadRequest {
	return &McsGetSecretUsingGETBadRequest{}
}

/*McsGetSecretUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type McsGetSecretUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *McsGetSecretUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /mcs/secret][%d] mcsGetSecretUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *McsGetSecretUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *McsGetSecretUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

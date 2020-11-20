// Code generated by go-swagger; DO NOT EDIT.

package ip_collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// IPCollectionCreateUsingPUTReader is a Reader for the IPCollectionCreateUsingPUT structure.
type IPCollectionCreateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPCollectionCreateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPCollectionCreateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewIPCollectionCreateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIPCollectionCreateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPCollectionCreateUsingPUTOK creates a IPCollectionCreateUsingPUTOK with default headers values
func NewIPCollectionCreateUsingPUTOK() *IPCollectionCreateUsingPUTOK {
	return &IPCollectionCreateUsingPUTOK{}
}

/*IPCollectionCreateUsingPUTOK handles this case with default header values.

OK
*/
type IPCollectionCreateUsingPUTOK struct {
	Payload *models.IPCollectionCreateResponse
}

func (o *IPCollectionCreateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /ipcs][%d] ipCollectionCreateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *IPCollectionCreateUsingPUTOK) GetPayload() *models.IPCollectionCreateResponse {
	return o.Payload
}

func (o *IPCollectionCreateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPCollectionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPCollectionCreateUsingPUTCreated creates a IPCollectionCreateUsingPUTCreated with default headers values
func NewIPCollectionCreateUsingPUTCreated() *IPCollectionCreateUsingPUTCreated {
	return &IPCollectionCreateUsingPUTCreated{}
}

/*IPCollectionCreateUsingPUTCreated handles this case with default header values.

Entity has been created
*/
type IPCollectionCreateUsingPUTCreated struct {
	Payload *models.IPCollectionCreateResponse
}

func (o *IPCollectionCreateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /ipcs][%d] ipCollectionCreateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *IPCollectionCreateUsingPUTCreated) GetPayload() *models.IPCollectionCreateResponse {
	return o.Payload
}

func (o *IPCollectionCreateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPCollectionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPCollectionCreateUsingPUTBadRequest creates a IPCollectionCreateUsingPUTBadRequest with default headers values
func NewIPCollectionCreateUsingPUTBadRequest() *IPCollectionCreateUsingPUTBadRequest {
	return &IPCollectionCreateUsingPUTBadRequest{}
}

/*IPCollectionCreateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type IPCollectionCreateUsingPUTBadRequest struct {
}

func (o *IPCollectionCreateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipcs][%d] ipCollectionCreateUsingPUTBadRequest ", 400)
}

func (o *IPCollectionCreateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// CustomServiceUpdateUsingPUTReader is a Reader for the CustomServiceUpdateUsingPUT structure.
type CustomServiceUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomServiceUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomServiceUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCustomServiceUpdateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomServiceUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomServiceUpdateUsingPUTOK creates a CustomServiceUpdateUsingPUTOK with default headers values
func NewCustomServiceUpdateUsingPUTOK() *CustomServiceUpdateUsingPUTOK {
	return &CustomServiceUpdateUsingPUTOK{}
}

/*CustomServiceUpdateUsingPUTOK handles this case with default header values.

OK
*/
type CustomServiceUpdateUsingPUTOK struct {
	Payload *models.UpdateCustomServiceResponse
}

func (o *CustomServiceUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/custom-services/{serviceId}][%d] customServiceUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *CustomServiceUpdateUsingPUTOK) GetPayload() *models.UpdateCustomServiceResponse {
	return o.Payload
}

func (o *CustomServiceUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateCustomServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomServiceUpdateUsingPUTCreated creates a CustomServiceUpdateUsingPUTCreated with default headers values
func NewCustomServiceUpdateUsingPUTCreated() *CustomServiceUpdateUsingPUTCreated {
	return &CustomServiceUpdateUsingPUTCreated{}
}

/*CustomServiceUpdateUsingPUTCreated handles this case with default header values.

Entity has been updated
*/
type CustomServiceUpdateUsingPUTCreated struct {
	Payload *models.UpdateCustomServiceResponse
}

func (o *CustomServiceUpdateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /network/custom-services/{serviceId}][%d] customServiceUpdateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *CustomServiceUpdateUsingPUTCreated) GetPayload() *models.UpdateCustomServiceResponse {
	return o.Payload
}

func (o *CustomServiceUpdateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateCustomServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomServiceUpdateUsingPUTBadRequest creates a CustomServiceUpdateUsingPUTBadRequest with default headers values
func NewCustomServiceUpdateUsingPUTBadRequest() *CustomServiceUpdateUsingPUTBadRequest {
	return &CustomServiceUpdateUsingPUTBadRequest{}
}

/*CustomServiceUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type CustomServiceUpdateUsingPUTBadRequest struct {
}

func (o *CustomServiceUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/custom-services/{serviceId}][%d] customServiceUpdateUsingPUTBadRequest ", 400)
}

func (o *CustomServiceUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

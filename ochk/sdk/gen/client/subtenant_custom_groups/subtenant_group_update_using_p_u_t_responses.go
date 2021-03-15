// Code generated by go-swagger; DO NOT EDIT.

package subtenant_custom_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SubtenantGroupUpdateUsingPUTReader is a Reader for the SubtenantGroupUpdateUsingPUT structure.
type SubtenantGroupUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtenantGroupUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtenantGroupUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubtenantGroupUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubtenantGroupUpdateUsingPUTOK creates a SubtenantGroupUpdateUsingPUTOK with default headers values
func NewSubtenantGroupUpdateUsingPUTOK() *SubtenantGroupUpdateUsingPUTOK {
	return &SubtenantGroupUpdateUsingPUTOK{}
}

/*SubtenantGroupUpdateUsingPUTOK handles this case with default header values.

Entity has been updated
*/
type SubtenantGroupUpdateUsingPUTOK struct {
	Payload *models.UpdateGroupResponse
}

func (o *SubtenantGroupUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /subtenants/{subtenantId}/groups/{groupId}][%d] subtenantGroupUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *SubtenantGroupUpdateUsingPUTOK) GetPayload() *models.UpdateGroupResponse {
	return o.Payload
}

func (o *SubtenantGroupUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantGroupUpdateUsingPUTBadRequest creates a SubtenantGroupUpdateUsingPUTBadRequest with default headers values
func NewSubtenantGroupUpdateUsingPUTBadRequest() *SubtenantGroupUpdateUsingPUTBadRequest {
	return &SubtenantGroupUpdateUsingPUTBadRequest{}
}

/*SubtenantGroupUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SubtenantGroupUpdateUsingPUTBadRequest struct {
}

func (o *SubtenantGroupUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subtenants/{subtenantId}/groups/{groupId}][%d] subtenantGroupUpdateUsingPUTBadRequest ", 400)
}

func (o *SubtenantGroupUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
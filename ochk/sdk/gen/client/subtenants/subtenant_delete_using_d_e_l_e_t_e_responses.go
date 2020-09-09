// Code generated by go-swagger; DO NOT EDIT.

package subtenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SubtenantDeleteUsingDELETEReader is a Reader for the SubtenantDeleteUsingDELETE structure.
type SubtenantDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtenantDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtenantDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubtenantDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubtenantDeleteUsingDELETEOK creates a SubtenantDeleteUsingDELETEOK with default headers values
func NewSubtenantDeleteUsingDELETEOK() *SubtenantDeleteUsingDELETEOK {
	return &SubtenantDeleteUsingDELETEOK{}
}

/*SubtenantDeleteUsingDELETEOK handles this case with default header values.

OK
*/
type SubtenantDeleteUsingDELETEOK struct {
	Payload *models.SubtenantDeleteResponse
}

func (o *SubtenantDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}][%d] subtenantDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *SubtenantDeleteUsingDELETEOK) GetPayload() *models.SubtenantDeleteResponse {
	return o.Payload
}

func (o *SubtenantDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtenantDeleteUsingDELETEBadRequest creates a SubtenantDeleteUsingDELETEBadRequest with default headers values
func NewSubtenantDeleteUsingDELETEBadRequest() *SubtenantDeleteUsingDELETEBadRequest {
	return &SubtenantDeleteUsingDELETEBadRequest{}
}

/*SubtenantDeleteUsingDELETEBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SubtenantDeleteUsingDELETEBadRequest struct {
}

func (o *SubtenantDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}][%d] subtenantDeleteUsingDELETEBadRequest ", 400)
}

func (o *SubtenantDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

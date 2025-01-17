// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SnapshotGetUsingGETReader is a Reader for the SnapshotGetUsingGET structure.
type SnapshotGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSnapshotGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSnapshotGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSnapshotGetUsingGETOK creates a SnapshotGetUsingGETOK with default headers values
func NewSnapshotGetUsingGETOK() *SnapshotGetUsingGETOK {
	return &SnapshotGetUsingGETOK{}
}

/*SnapshotGetUsingGETOK handles this case with default header values.

OK
*/
type SnapshotGetUsingGETOK struct {
	Payload *models.SnapshotGetResponse
}

func (o *SnapshotGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vcs/snapshots/{snapshotId}][%d] snapshotGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *SnapshotGetUsingGETOK) GetPayload() *models.SnapshotGetResponse {
	return o.Payload
}

func (o *SnapshotGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotGetUsingGETBadRequest creates a SnapshotGetUsingGETBadRequest with default headers values
func NewSnapshotGetUsingGETBadRequest() *SnapshotGetUsingGETBadRequest {
	return &SnapshotGetUsingGETBadRequest{}
}

/*SnapshotGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SnapshotGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *SnapshotGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/snapshots/{snapshotId}][%d] snapshotGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *SnapshotGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *SnapshotGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotGetUsingGETNotFound creates a SnapshotGetUsingGETNotFound with default headers values
func NewSnapshotGetUsingGETNotFound() *SnapshotGetUsingGETNotFound {
	return &SnapshotGetUsingGETNotFound{}
}

/*SnapshotGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type SnapshotGetUsingGETNotFound struct {
}

func (o *SnapshotGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/snapshots/{snapshotId}][%d] snapshotGetUsingGETNotFound ", 404)
}

func (o *SnapshotGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

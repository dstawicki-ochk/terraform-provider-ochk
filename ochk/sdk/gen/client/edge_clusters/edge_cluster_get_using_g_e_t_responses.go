// Code generated by go-swagger; DO NOT EDIT.

package edge_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// EdgeClusterGetUsingGETReader is a Reader for the EdgeClusterGetUsingGET structure.
type EdgeClusterGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeClusterGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeClusterGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeClusterGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeClusterGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEdgeClusterGetUsingGETOK creates a EdgeClusterGetUsingGETOK with default headers values
func NewEdgeClusterGetUsingGETOK() *EdgeClusterGetUsingGETOK {
	return &EdgeClusterGetUsingGETOK{}
}

/*EdgeClusterGetUsingGETOK handles this case with default header values.

OK
*/
type EdgeClusterGetUsingGETOK struct {
	Payload *models.EdgeClusterGetResponse
}

func (o *EdgeClusterGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/edge-clusters/{edgeClusterId}][%d] edgeClusterGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *EdgeClusterGetUsingGETOK) GetPayload() *models.EdgeClusterGetResponse {
	return o.Payload
}

func (o *EdgeClusterGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeClusterGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeClusterGetUsingGETBadRequest creates a EdgeClusterGetUsingGETBadRequest with default headers values
func NewEdgeClusterGetUsingGETBadRequest() *EdgeClusterGetUsingGETBadRequest {
	return &EdgeClusterGetUsingGETBadRequest{}
}

/*EdgeClusterGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type EdgeClusterGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *EdgeClusterGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/edge-clusters/{edgeClusterId}][%d] edgeClusterGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeClusterGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *EdgeClusterGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeClusterGetUsingGETNotFound creates a EdgeClusterGetUsingGETNotFound with default headers values
func NewEdgeClusterGetUsingGETNotFound() *EdgeClusterGetUsingGETNotFound {
	return &EdgeClusterGetUsingGETNotFound{}
}

/*EdgeClusterGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type EdgeClusterGetUsingGETNotFound struct {
}

func (o *EdgeClusterGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/edge-clusters/{edgeClusterId}][%d] edgeClusterGetUsingGETNotFound ", 404)
}

func (o *EdgeClusterGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

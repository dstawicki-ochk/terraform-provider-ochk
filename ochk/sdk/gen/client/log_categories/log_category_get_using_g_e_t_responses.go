// Code generated by go-swagger; DO NOT EDIT.

package log_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// LogCategoryGetUsingGETReader is a Reader for the LogCategoryGetUsingGET structure.
type LogCategoryGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogCategoryGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogCategoryGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogCategoryGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLogCategoryGetUsingGETOK creates a LogCategoryGetUsingGETOK with default headers values
func NewLogCategoryGetUsingGETOK() *LogCategoryGetUsingGETOK {
	return &LogCategoryGetUsingGETOK{}
}

/*LogCategoryGetUsingGETOK handles this case with default header values.

OK
*/
type LogCategoryGetUsingGETOK struct {
	Payload *models.LogCategoryGetResponse
}

func (o *LogCategoryGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] logCategoryGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *LogCategoryGetUsingGETOK) GetPayload() *models.LogCategoryGetResponse {
	return o.Payload
}

func (o *LogCategoryGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogCategoryGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogCategoryGetUsingGETBadRequest creates a LogCategoryGetUsingGETBadRequest with default headers values
func NewLogCategoryGetUsingGETBadRequest() *LogCategoryGetUsingGETBadRequest {
	return &LogCategoryGetUsingGETBadRequest{}
}

/*LogCategoryGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type LogCategoryGetUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *LogCategoryGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /log/categories/{logCategoryId}][%d] logCategoryGetUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *LogCategoryGetUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *LogCategoryGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

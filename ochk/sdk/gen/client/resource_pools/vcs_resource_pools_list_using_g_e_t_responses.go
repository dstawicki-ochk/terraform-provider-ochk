// Code generated by go-swagger; DO NOT EDIT.

package resource_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VcsResourcePoolsListUsingGETReader is a Reader for the VcsResourcePoolsListUsingGET structure.
type VcsResourcePoolsListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsResourcePoolsListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsResourcePoolsListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsResourcePoolsListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVcsResourcePoolsListUsingGETOK creates a VcsResourcePoolsListUsingGETOK with default headers values
func NewVcsResourcePoolsListUsingGETOK() *VcsResourcePoolsListUsingGETOK {
	return &VcsResourcePoolsListUsingGETOK{}
}

/*VcsResourcePoolsListUsingGETOK handles this case with default header values.

OK
*/
type VcsResourcePoolsListUsingGETOK struct {
	Payload *models.VcsResourcePoolsListResponse
}

func (o *VcsResourcePoolsListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vcs/resourcepools][%d] vcsResourcePoolsListUsingGETOK  %+v", 200, o.Payload)
}

func (o *VcsResourcePoolsListUsingGETOK) GetPayload() *models.VcsResourcePoolsListResponse {
	return o.Payload
}

func (o *VcsResourcePoolsListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsResourcePoolsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsResourcePoolsListUsingGETBadRequest creates a VcsResourcePoolsListUsingGETBadRequest with default headers values
func NewVcsResourcePoolsListUsingGETBadRequest() *VcsResourcePoolsListUsingGETBadRequest {
	return &VcsResourcePoolsListUsingGETBadRequest{}
}

/*VcsResourcePoolsListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsResourcePoolsListUsingGETBadRequest struct {
}

func (o *VcsResourcePoolsListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vcs/resourcepools][%d] vcsResourcePoolsListUsingGETBadRequest ", 400)
}

func (o *VcsResourcePoolsListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

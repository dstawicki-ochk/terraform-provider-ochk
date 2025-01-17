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

// GroupMemberListUsingGETReader is a Reader for the GroupMemberListUsingGET structure.
type GroupMemberListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupMemberListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupMemberListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGroupMemberListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupMemberListUsingGETOK creates a GroupMemberListUsingGETOK with default headers values
func NewGroupMemberListUsingGETOK() *GroupMemberListUsingGETOK {
	return &GroupMemberListUsingGETOK{}
}

/*GroupMemberListUsingGETOK handles this case with default header values.

OK
*/
type GroupMemberListUsingGETOK struct {
	Payload *models.GroupListResponse
}

func (o *GroupMemberListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{parentGroupId}/members/groups][%d] groupMemberListUsingGETOK  %+v", 200, o.Payload)
}

func (o *GroupMemberListUsingGETOK) GetPayload() *models.GroupListResponse {
	return o.Payload
}

func (o *GroupMemberListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupMemberListUsingGETBadRequest creates a GroupMemberListUsingGETBadRequest with default headers values
func NewGroupMemberListUsingGETBadRequest() *GroupMemberListUsingGETBadRequest {
	return &GroupMemberListUsingGETBadRequest{}
}

/*GroupMemberListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GroupMemberListUsingGETBadRequest struct {
	Payload *models.ProxyResponseMessage
}

func (o *GroupMemberListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/groups/{parentGroupId}/members/groups][%d] groupMemberListUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GroupMemberListUsingGETBadRequest) GetPayload() *models.ProxyResponseMessage {
	return o.Payload
}

func (o *GroupMemberListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

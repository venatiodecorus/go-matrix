// Code generated by go-swagger; DO NOT EDIT.

package room_participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// GetMembersByRoomReader is a Reader for the GetMembersByRoom structure.
type GetMembersByRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMembersByRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMembersByRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetMembersByRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMembersByRoomOK creates a GetMembersByRoomOK with default headers values
func NewGetMembersByRoomOK() *GetMembersByRoomOK {
	return &GetMembersByRoomOK{}
}

/*GetMembersByRoomOK handles this case with default header values.

A list of members of the room. If you are joined to the room then
this will be the current members of the room. If you have left the
room then this will be the members of the room when you left.
*/
type GetMembersByRoomOK struct {
	Payload *models.GetMembersByRoomOKBody
}

func (o *GetMembersByRoomOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/members][%d] getMembersByRoomOK  %+v", 200, o.Payload)
}

func (o *GetMembersByRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMembersByRoomOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMembersByRoomForbidden creates a GetMembersByRoomForbidden with default headers values
func NewGetMembersByRoomForbidden() *GetMembersByRoomForbidden {
	return &GetMembersByRoomForbidden{}
}

/*GetMembersByRoomForbidden handles this case with default header values.

You aren't a member of the room and weren't previously a member of the room.

*/
type GetMembersByRoomForbidden struct {
}

func (o *GetMembersByRoomForbidden) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/members][%d] getMembersByRoomForbidden ", 403)
}

func (o *GetMembersByRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

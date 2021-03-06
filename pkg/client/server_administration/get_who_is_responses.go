// Code generated by go-swagger; DO NOT EDIT.

package server_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// GetWhoIsReader is a Reader for the GetWhoIs structure.
type GetWhoIsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWhoIsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWhoIsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWhoIsOK creates a GetWhoIsOK with default headers values
func NewGetWhoIsOK() *GetWhoIsOK {
	return &GetWhoIsOK{}
}

/*GetWhoIsOK handles this case with default header values.

The lookup was successful.
*/
type GetWhoIsOK struct {
	Payload *models.GetWhoIsOKBody
}

func (o *GetWhoIsOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/admin/whois/{userId}][%d] getWhoIsOK  %+v", 200, o.Payload)
}

func (o *GetWhoIsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetWhoIsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

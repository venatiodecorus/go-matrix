// Code generated by go-swagger; DO NOT EDIT.

package user_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRoomTagParams creates a new DeleteRoomTagParams object
// with the default values initialized.
func NewDeleteRoomTagParams() *DeleteRoomTagParams {
	var ()
	return &DeleteRoomTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoomTagParamsWithTimeout creates a new DeleteRoomTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoomTagParamsWithTimeout(timeout time.Duration) *DeleteRoomTagParams {
	var ()
	return &DeleteRoomTagParams{

		timeout: timeout,
	}
}

// NewDeleteRoomTagParamsWithContext creates a new DeleteRoomTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoomTagParamsWithContext(ctx context.Context) *DeleteRoomTagParams {
	var ()
	return &DeleteRoomTagParams{

		Context: ctx,
	}
}

// NewDeleteRoomTagParamsWithHTTPClient creates a new DeleteRoomTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoomTagParamsWithHTTPClient(client *http.Client) *DeleteRoomTagParams {
	var ()
	return &DeleteRoomTagParams{
		HTTPClient: client,
	}
}

/*DeleteRoomTagParams contains all the parameters to send to the API endpoint
for the delete room tag operation typically these are written to a http.Request
*/
type DeleteRoomTagParams struct {

	/*RoomID
	  The ID of the room to remove a tag from.

	*/
	RoomID string
	/*Tag
	  The tag to remove.

	*/
	Tag string
	/*UserID
	  The id of the user to remove a tag for. The access token must be
	authorized to make requests for this user ID.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete room tag params
func (o *DeleteRoomTagParams) WithTimeout(timeout time.Duration) *DeleteRoomTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete room tag params
func (o *DeleteRoomTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete room tag params
func (o *DeleteRoomTagParams) WithContext(ctx context.Context) *DeleteRoomTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete room tag params
func (o *DeleteRoomTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete room tag params
func (o *DeleteRoomTagParams) WithHTTPClient(client *http.Client) *DeleteRoomTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete room tag params
func (o *DeleteRoomTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoomID adds the roomID to the delete room tag params
func (o *DeleteRoomTagParams) WithRoomID(roomID string) *DeleteRoomTagParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the delete room tag params
func (o *DeleteRoomTagParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WithTag adds the tag to the delete room tag params
func (o *DeleteRoomTagParams) WithTag(tag string) *DeleteRoomTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the delete room tag params
func (o *DeleteRoomTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WithUserID adds the userID to the delete room tag params
func (o *DeleteRoomTagParams) WithUserID(userID string) *DeleteRoomTagParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete room tag params
func (o *DeleteRoomTagParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoomTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roomId
	if err := r.SetPathParam("roomId", o.RoomID); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

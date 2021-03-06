// Code generated by go-swagger; DO NOT EDIT.

package presence

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

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// NewSetPresenceParams creates a new SetPresenceParams object
// with the default values initialized.
func NewSetPresenceParams() *SetPresenceParams {
	var ()
	return &SetPresenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPresenceParamsWithTimeout creates a new SetPresenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPresenceParamsWithTimeout(timeout time.Duration) *SetPresenceParams {
	var ()
	return &SetPresenceParams{

		timeout: timeout,
	}
}

// NewSetPresenceParamsWithContext creates a new SetPresenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPresenceParamsWithContext(ctx context.Context) *SetPresenceParams {
	var ()
	return &SetPresenceParams{

		Context: ctx,
	}
}

// NewSetPresenceParamsWithHTTPClient creates a new SetPresenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPresenceParamsWithHTTPClient(client *http.Client) *SetPresenceParams {
	var ()
	return &SetPresenceParams{
		HTTPClient: client,
	}
}

/*SetPresenceParams contains all the parameters to send to the API endpoint
for the set presence operation typically these are written to a http.Request
*/
type SetPresenceParams struct {

	/*PresenceState
	  The updated presence state.

	*/
	PresenceState *models.SetPresenceParamsBody
	/*UserID
	  The user whose presence state to update.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set presence params
func (o *SetPresenceParams) WithTimeout(timeout time.Duration) *SetPresenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set presence params
func (o *SetPresenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set presence params
func (o *SetPresenceParams) WithContext(ctx context.Context) *SetPresenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set presence params
func (o *SetPresenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set presence params
func (o *SetPresenceParams) WithHTTPClient(client *http.Client) *SetPresenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set presence params
func (o *SetPresenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPresenceState adds the presenceState to the set presence params
func (o *SetPresenceParams) WithPresenceState(presenceState *models.SetPresenceParamsBody) *SetPresenceParams {
	o.SetPresenceState(presenceState)
	return o
}

// SetPresenceState adds the presenceState to the set presence params
func (o *SetPresenceParams) SetPresenceState(presenceState *models.SetPresenceParamsBody) {
	o.PresenceState = presenceState
}

// WithUserID adds the userID to the set presence params
func (o *SetPresenceParams) WithUserID(userID string) *SetPresenceParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set presence params
func (o *SetPresenceParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetPresenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PresenceState != nil {
		if err := r.SetBodyParam(o.PresenceState); err != nil {
			return err
		}
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

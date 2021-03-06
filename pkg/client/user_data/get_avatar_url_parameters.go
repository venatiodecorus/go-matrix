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

// NewGetAvatarURLParams creates a new GetAvatarURLParams object
// with the default values initialized.
func NewGetAvatarURLParams() *GetAvatarURLParams {
	var ()
	return &GetAvatarURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvatarURLParamsWithTimeout creates a new GetAvatarURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAvatarURLParamsWithTimeout(timeout time.Duration) *GetAvatarURLParams {
	var ()
	return &GetAvatarURLParams{

		timeout: timeout,
	}
}

// NewGetAvatarURLParamsWithContext creates a new GetAvatarURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAvatarURLParamsWithContext(ctx context.Context) *GetAvatarURLParams {
	var ()
	return &GetAvatarURLParams{

		Context: ctx,
	}
}

// NewGetAvatarURLParamsWithHTTPClient creates a new GetAvatarURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAvatarURLParamsWithHTTPClient(client *http.Client) *GetAvatarURLParams {
	var ()
	return &GetAvatarURLParams{
		HTTPClient: client,
	}
}

/*GetAvatarURLParams contains all the parameters to send to the API endpoint
for the get avatar Url operation typically these are written to a http.Request
*/
type GetAvatarURLParams struct {

	/*UserID
	  The user whose avatar URL to get.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get avatar Url params
func (o *GetAvatarURLParams) WithTimeout(timeout time.Duration) *GetAvatarURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get avatar Url params
func (o *GetAvatarURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get avatar Url params
func (o *GetAvatarURLParams) WithContext(ctx context.Context) *GetAvatarURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get avatar Url params
func (o *GetAvatarURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get avatar Url params
func (o *GetAvatarURLParams) WithHTTPClient(client *http.Client) *GetAvatarURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get avatar Url params
func (o *GetAvatarURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get avatar Url params
func (o *GetAvatarURLParams) WithUserID(userID string) *GetAvatarURLParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get avatar Url params
func (o *GetAvatarURLParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvatarURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

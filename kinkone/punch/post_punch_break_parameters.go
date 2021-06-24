// Code generated by go-swagger; DO NOT EDIT.

package punch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostPunchBreakParams creates a new PostPunchBreakParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPunchBreakParams() *PostPunchBreakParams {
	return &PostPunchBreakParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPunchBreakParamsWithTimeout creates a new PostPunchBreakParams object
// with the ability to set a timeout on a request.
func NewPostPunchBreakParamsWithTimeout(timeout time.Duration) *PostPunchBreakParams {
	return &PostPunchBreakParams{
		timeout: timeout,
	}
}

// NewPostPunchBreakParamsWithContext creates a new PostPunchBreakParams object
// with the ability to set a context for a request.
func NewPostPunchBreakParamsWithContext(ctx context.Context) *PostPunchBreakParams {
	return &PostPunchBreakParams{
		Context: ctx,
	}
}

// NewPostPunchBreakParamsWithHTTPClient creates a new PostPunchBreakParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPunchBreakParamsWithHTTPClient(client *http.Client) *PostPunchBreakParams {
	return &PostPunchBreakParams{
		HTTPClient: client,
	}
}

/* PostPunchBreakParams contains all the parameters to send to the API endpoint
   for the post punch break operation.

   Typically these are written to a http.Request.
*/
type PostPunchBreakParams struct {

	/* Allnight.

	   徹夜フラグ（0: 通常勤務, 1: 徹夜勤務）

	   Default: "0"
	*/
	Allnight *string

	/* Email.

	   メールアドレス
	*/
	Email string

	/* Latitude.

	   緯度
	*/
	Latitude *string

	/* Longitude.

	   経度
	*/
	Longitude *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post punch break params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPunchBreakParams) WithDefaults() *PostPunchBreakParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post punch break params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPunchBreakParams) SetDefaults() {
	var (
		allnightDefault = string("0")
	)

	val := PostPunchBreakParams{
		Allnight: &allnightDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post punch break params
func (o *PostPunchBreakParams) WithTimeout(timeout time.Duration) *PostPunchBreakParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post punch break params
func (o *PostPunchBreakParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post punch break params
func (o *PostPunchBreakParams) WithContext(ctx context.Context) *PostPunchBreakParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post punch break params
func (o *PostPunchBreakParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post punch break params
func (o *PostPunchBreakParams) WithHTTPClient(client *http.Client) *PostPunchBreakParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post punch break params
func (o *PostPunchBreakParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllnight adds the allnight to the post punch break params
func (o *PostPunchBreakParams) WithAllnight(allnight *string) *PostPunchBreakParams {
	o.SetAllnight(allnight)
	return o
}

// SetAllnight adds the allnight to the post punch break params
func (o *PostPunchBreakParams) SetAllnight(allnight *string) {
	o.Allnight = allnight
}

// WithEmail adds the email to the post punch break params
func (o *PostPunchBreakParams) WithEmail(email string) *PostPunchBreakParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the post punch break params
func (o *PostPunchBreakParams) SetEmail(email string) {
	o.Email = email
}

// WithLatitude adds the latitude to the post punch break params
func (o *PostPunchBreakParams) WithLatitude(latitude *string) *PostPunchBreakParams {
	o.SetLatitude(latitude)
	return o
}

// SetLatitude adds the latitude to the post punch break params
func (o *PostPunchBreakParams) SetLatitude(latitude *string) {
	o.Latitude = latitude
}

// WithLongitude adds the longitude to the post punch break params
func (o *PostPunchBreakParams) WithLongitude(longitude *string) *PostPunchBreakParams {
	o.SetLongitude(longitude)
	return o
}

// SetLongitude adds the longitude to the post punch break params
func (o *PostPunchBreakParams) SetLongitude(longitude *string) {
	o.Longitude = longitude
}

// WriteToRequest writes these params to a swagger request
func (o *PostPunchBreakParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Allnight != nil {

		// form param allnight
		var frAllnight string
		if o.Allnight != nil {
			frAllnight = *o.Allnight
		}
		fAllnight := frAllnight
		if fAllnight != "" {
			if err := r.SetFormParam("allnight", fAllnight); err != nil {
				return err
			}
		}
	}

	// form param email
	frEmail := o.Email
	fEmail := frEmail
	if fEmail != "" {
		if err := r.SetFormParam("email", fEmail); err != nil {
			return err
		}
	}

	if o.Latitude != nil {

		// form param latitude
		var frLatitude string
		if o.Latitude != nil {
			frLatitude = *o.Latitude
		}
		fLatitude := frLatitude
		if fLatitude != "" {
			if err := r.SetFormParam("latitude", fLatitude); err != nil {
				return err
			}
		}
	}

	if o.Longitude != nil {

		// form param longitude
		var frLongitude string
		if o.Longitude != nil {
			frLongitude = *o.Longitude
		}
		fLongitude := frLongitude
		if fLongitude != "" {
			if err := r.SetFormParam("longitude", fLongitude); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
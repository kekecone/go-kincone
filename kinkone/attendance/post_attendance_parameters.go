// Code generated by go-swagger; DO NOT EDIT.

package attendance

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
	"github.com/go-openapi/swag"
)

// NewPostAttendanceParams creates a new PostAttendanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAttendanceParams() *PostAttendanceParams {
	return &PostAttendanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAttendanceParamsWithTimeout creates a new PostAttendanceParams object
// with the ability to set a timeout on a request.
func NewPostAttendanceParamsWithTimeout(timeout time.Duration) *PostAttendanceParams {
	return &PostAttendanceParams{
		timeout: timeout,
	}
}

// NewPostAttendanceParamsWithContext creates a new PostAttendanceParams object
// with the ability to set a context for a request.
func NewPostAttendanceParamsWithContext(ctx context.Context) *PostAttendanceParams {
	return &PostAttendanceParams{
		Context: ctx,
	}
}

// NewPostAttendanceParamsWithHTTPClient creates a new PostAttendanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAttendanceParamsWithHTTPClient(client *http.Client) *PostAttendanceParams {
	return &PostAttendanceParams{
		HTTPClient: client,
	}
}

/* PostAttendanceParams contains all the parameters to send to the API endpoint
   for the post attendance operation.

   Typically these are written to a http.Request.
*/
type PostAttendanceParams struct {

	/* AutoGiveBreaks.

	   1を指定した場合は、就業形態の「休憩の付与方法」が "時間帯を指定する" の場合に休憩時間が自動的に計算されます。
	*/
	AutoGiveBreaks *int64

	/* Email.

	   メールアドレス
	*/
	Email string

	/* End.

	   退勤時刻（hh:mm）
	*/
	End *string

	/* Note.

	   備考
	*/
	Note *string

	/* Start.

	   出勤時刻（hh:mm）
	*/
	Start *string

	/* WorkDay.

	   出勤日（YYYY-MM-DD）
	*/
	WorkDay string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post attendance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAttendanceParams) WithDefaults() *PostAttendanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post attendance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAttendanceParams) SetDefaults() {
	var (
		autoGiveBreaksDefault = int64(0)
	)

	val := PostAttendanceParams{
		AutoGiveBreaks: &autoGiveBreaksDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post attendance params
func (o *PostAttendanceParams) WithTimeout(timeout time.Duration) *PostAttendanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post attendance params
func (o *PostAttendanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post attendance params
func (o *PostAttendanceParams) WithContext(ctx context.Context) *PostAttendanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post attendance params
func (o *PostAttendanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post attendance params
func (o *PostAttendanceParams) WithHTTPClient(client *http.Client) *PostAttendanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post attendance params
func (o *PostAttendanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAutoGiveBreaks adds the autoGiveBreaks to the post attendance params
func (o *PostAttendanceParams) WithAutoGiveBreaks(autoGiveBreaks *int64) *PostAttendanceParams {
	o.SetAutoGiveBreaks(autoGiveBreaks)
	return o
}

// SetAutoGiveBreaks adds the autoGiveBreaks to the post attendance params
func (o *PostAttendanceParams) SetAutoGiveBreaks(autoGiveBreaks *int64) {
	o.AutoGiveBreaks = autoGiveBreaks
}

// WithEmail adds the email to the post attendance params
func (o *PostAttendanceParams) WithEmail(email string) *PostAttendanceParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the post attendance params
func (o *PostAttendanceParams) SetEmail(email string) {
	o.Email = email
}

// WithEnd adds the end to the post attendance params
func (o *PostAttendanceParams) WithEnd(end *string) *PostAttendanceParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the post attendance params
func (o *PostAttendanceParams) SetEnd(end *string) {
	o.End = end
}

// WithNote adds the note to the post attendance params
func (o *PostAttendanceParams) WithNote(note *string) *PostAttendanceParams {
	o.SetNote(note)
	return o
}

// SetNote adds the note to the post attendance params
func (o *PostAttendanceParams) SetNote(note *string) {
	o.Note = note
}

// WithStart adds the start to the post attendance params
func (o *PostAttendanceParams) WithStart(start *string) *PostAttendanceParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the post attendance params
func (o *PostAttendanceParams) SetStart(start *string) {
	o.Start = start
}

// WithWorkDay adds the workDay to the post attendance params
func (o *PostAttendanceParams) WithWorkDay(workDay string) *PostAttendanceParams {
	o.SetWorkDay(workDay)
	return o
}

// SetWorkDay adds the workDay to the post attendance params
func (o *PostAttendanceParams) SetWorkDay(workDay string) {
	o.WorkDay = workDay
}

// WriteToRequest writes these params to a swagger request
func (o *PostAttendanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AutoGiveBreaks != nil {

		// form param auto_give_breaks
		var frAutoGiveBreaks int64
		if o.AutoGiveBreaks != nil {
			frAutoGiveBreaks = *o.AutoGiveBreaks
		}
		fAutoGiveBreaks := swag.FormatInt64(frAutoGiveBreaks)
		if fAutoGiveBreaks != "" {
			if err := r.SetFormParam("auto_give_breaks", fAutoGiveBreaks); err != nil {
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

	if o.End != nil {

		// form param end
		var frEnd string
		if o.End != nil {
			frEnd = *o.End
		}
		fEnd := frEnd
		if fEnd != "" {
			if err := r.SetFormParam("end", fEnd); err != nil {
				return err
			}
		}
	}

	if o.Note != nil {

		// form param note
		var frNote string
		if o.Note != nil {
			frNote = *o.Note
		}
		fNote := frNote
		if fNote != "" {
			if err := r.SetFormParam("note", fNote); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// form param start
		var frStart string
		if o.Start != nil {
			frStart = *o.Start
		}
		fStart := frStart
		if fStart != "" {
			if err := r.SetFormParam("start", fStart); err != nil {
				return err
			}
		}
	}

	// form param work_day
	frWorkDay := o.WorkDay
	fWorkDay := frWorkDay
	if fWorkDay != "" {
		if err := r.SetFormParam("work_day", fWorkDay); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

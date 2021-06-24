// Code generated by go-swagger; DO NOT EDIT.

package expenses

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

// NewPostExpenseParams creates a new PostExpenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExpenseParams() *PostExpenseParams {
	return &PostExpenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExpenseParamsWithTimeout creates a new PostExpenseParams object
// with the ability to set a timeout on a request.
func NewPostExpenseParamsWithTimeout(timeout time.Duration) *PostExpenseParams {
	return &PostExpenseParams{
		timeout: timeout,
	}
}

// NewPostExpenseParamsWithContext creates a new PostExpenseParams object
// with the ability to set a context for a request.
func NewPostExpenseParamsWithContext(ctx context.Context) *PostExpenseParams {
	return &PostExpenseParams{
		Context: ctx,
	}
}

// NewPostExpenseParamsWithHTTPClient creates a new PostExpenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExpenseParamsWithHTTPClient(client *http.Client) *PostExpenseParams {
	return &PostExpenseParams{
		HTTPClient: client,
	}
}

/* PostExpenseParams contains all the parameters to send to the API endpoint
   for the post expense operation.

   Typically these are written to a http.Request.
*/
type PostExpenseParams struct {

	/* Date.

	   利用日（YYYY-MM-DD）
	*/
	Date string

	/* Email.

	   メールアドレス
	*/
	Email string

	/* Expense.

	   交通費
	*/
	Expense int64

	/* In.

	   利用区間（入場駅）
	*/
	In *string

	/* Note.

	   訪問先/備考
	*/
	Note *string

	/* Out.

	   利用区間（出場駅）
	*/
	Out *string

	/* Purpose.

	   用途（交通費の明細に用途欄を追加する設定になっている場合のみ使用出来ます）
	*/
	Purpose *string

	/* Type.

	   交通手段（1: 電車, 2: バス, 3: タクシー, 4: 新幹線/特急, 5: 飛行機, 6: 車, 7: 船/フェリー, 8: その他, 0: 物販）
	*/
	Type int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post expense params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExpenseParams) WithDefaults() *PostExpenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post expense params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExpenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post expense params
func (o *PostExpenseParams) WithTimeout(timeout time.Duration) *PostExpenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post expense params
func (o *PostExpenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post expense params
func (o *PostExpenseParams) WithContext(ctx context.Context) *PostExpenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post expense params
func (o *PostExpenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post expense params
func (o *PostExpenseParams) WithHTTPClient(client *http.Client) *PostExpenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post expense params
func (o *PostExpenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the post expense params
func (o *PostExpenseParams) WithDate(date string) *PostExpenseParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the post expense params
func (o *PostExpenseParams) SetDate(date string) {
	o.Date = date
}

// WithEmail adds the email to the post expense params
func (o *PostExpenseParams) WithEmail(email string) *PostExpenseParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the post expense params
func (o *PostExpenseParams) SetEmail(email string) {
	o.Email = email
}

// WithExpense adds the expense to the post expense params
func (o *PostExpenseParams) WithExpense(expense int64) *PostExpenseParams {
	o.SetExpense(expense)
	return o
}

// SetExpense adds the expense to the post expense params
func (o *PostExpenseParams) SetExpense(expense int64) {
	o.Expense = expense
}

// WithIn adds the in to the post expense params
func (o *PostExpenseParams) WithIn(in *string) *PostExpenseParams {
	o.SetIn(in)
	return o
}

// SetIn adds the in to the post expense params
func (o *PostExpenseParams) SetIn(in *string) {
	o.In = in
}

// WithNote adds the note to the post expense params
func (o *PostExpenseParams) WithNote(note *string) *PostExpenseParams {
	o.SetNote(note)
	return o
}

// SetNote adds the note to the post expense params
func (o *PostExpenseParams) SetNote(note *string) {
	o.Note = note
}

// WithOut adds the out to the post expense params
func (o *PostExpenseParams) WithOut(out *string) *PostExpenseParams {
	o.SetOut(out)
	return o
}

// SetOut adds the out to the post expense params
func (o *PostExpenseParams) SetOut(out *string) {
	o.Out = out
}

// WithPurpose adds the purpose to the post expense params
func (o *PostExpenseParams) WithPurpose(purpose *string) *PostExpenseParams {
	o.SetPurpose(purpose)
	return o
}

// SetPurpose adds the purpose to the post expense params
func (o *PostExpenseParams) SetPurpose(purpose *string) {
	o.Purpose = purpose
}

// WithType adds the typeVar to the post expense params
func (o *PostExpenseParams) WithType(typeVar int64) *PostExpenseParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the post expense params
func (o *PostExpenseParams) SetType(typeVar int64) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PostExpenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param date
	frDate := o.Date
	fDate := frDate
	if fDate != "" {
		if err := r.SetFormParam("date", fDate); err != nil {
			return err
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

	// form param expense
	frExpense := o.Expense
	fExpense := swag.FormatInt64(frExpense)
	if fExpense != "" {
		if err := r.SetFormParam("expense", fExpense); err != nil {
			return err
		}
	}

	if o.In != nil {

		// form param in
		var frIn string
		if o.In != nil {
			frIn = *o.In
		}
		fIn := frIn
		if fIn != "" {
			if err := r.SetFormParam("in", fIn); err != nil {
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

	if o.Out != nil {

		// form param out
		var frOut string
		if o.Out != nil {
			frOut = *o.Out
		}
		fOut := frOut
		if fOut != "" {
			if err := r.SetFormParam("out", fOut); err != nil {
				return err
			}
		}
	}

	if o.Purpose != nil {

		// form param purpose
		var frPurpose string
		if o.Purpose != nil {
			frPurpose = *o.Purpose
		}
		fPurpose := frPurpose
		if fPurpose != "" {
			if err := r.SetFormParam("purpose", fPurpose); err != nil {
				return err
			}
		}
	}

	// form param type
	frType := o.Type
	fType := swag.FormatInt64(frType)
	if fType != "" {
		if err := r.SetFormParam("type", fType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
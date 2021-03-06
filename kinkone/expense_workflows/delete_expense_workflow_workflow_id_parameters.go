// Code generated by go-swagger; DO NOT EDIT.

package expense_workflows

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

// NewDeleteExpenseWorkflowWorkflowIDParams creates a new DeleteExpenseWorkflowWorkflowIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteExpenseWorkflowWorkflowIDParams() *DeleteExpenseWorkflowWorkflowIDParams {
	return &DeleteExpenseWorkflowWorkflowIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExpenseWorkflowWorkflowIDParamsWithTimeout creates a new DeleteExpenseWorkflowWorkflowIDParams object
// with the ability to set a timeout on a request.
func NewDeleteExpenseWorkflowWorkflowIDParamsWithTimeout(timeout time.Duration) *DeleteExpenseWorkflowWorkflowIDParams {
	return &DeleteExpenseWorkflowWorkflowIDParams{
		timeout: timeout,
	}
}

// NewDeleteExpenseWorkflowWorkflowIDParamsWithContext creates a new DeleteExpenseWorkflowWorkflowIDParams object
// with the ability to set a context for a request.
func NewDeleteExpenseWorkflowWorkflowIDParamsWithContext(ctx context.Context) *DeleteExpenseWorkflowWorkflowIDParams {
	return &DeleteExpenseWorkflowWorkflowIDParams{
		Context: ctx,
	}
}

// NewDeleteExpenseWorkflowWorkflowIDParamsWithHTTPClient creates a new DeleteExpenseWorkflowWorkflowIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteExpenseWorkflowWorkflowIDParamsWithHTTPClient(client *http.Client) *DeleteExpenseWorkflowWorkflowIDParams {
	return &DeleteExpenseWorkflowWorkflowIDParams{
		HTTPClient: client,
	}
}

/* DeleteExpenseWorkflowWorkflowIDParams contains all the parameters to send to the API endpoint
   for the delete expense workflow workflow ID operation.

   Typically these are written to a http.Request.
*/
type DeleteExpenseWorkflowWorkflowIDParams struct {

	/* Note.

	   ?????????????????????
	*/
	Note *string

	/* WorkflowID.

	   ??????ID
	*/
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete expense workflow workflow ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExpenseWorkflowWorkflowIDParams) WithDefaults() *DeleteExpenseWorkflowWorkflowIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete expense workflow workflow ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExpenseWorkflowWorkflowIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) WithTimeout(timeout time.Duration) *DeleteExpenseWorkflowWorkflowIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) WithContext(ctx context.Context) *DeleteExpenseWorkflowWorkflowIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) WithHTTPClient(client *http.Client) *DeleteExpenseWorkflowWorkflowIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNote adds the note to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) WithNote(note *string) *DeleteExpenseWorkflowWorkflowIDParams {
	o.SetNote(note)
	return o
}

// SetNote adds the note to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) SetNote(note *string) {
	o.Note = note
}

// WithWorkflowID adds the workflowID to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) WithWorkflowID(workflowID string) *DeleteExpenseWorkflowWorkflowIDParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the delete expense workflow workflow ID params
func (o *DeleteExpenseWorkflowWorkflowIDParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExpenseWorkflowWorkflowIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param workflowID
	if err := r.SetPathParam("workflowID", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

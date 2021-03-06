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

// NewPostExpenseWorkflowWorkflowIDParams creates a new PostExpenseWorkflowWorkflowIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExpenseWorkflowWorkflowIDParams() *PostExpenseWorkflowWorkflowIDParams {
	return &PostExpenseWorkflowWorkflowIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExpenseWorkflowWorkflowIDParamsWithTimeout creates a new PostExpenseWorkflowWorkflowIDParams object
// with the ability to set a timeout on a request.
func NewPostExpenseWorkflowWorkflowIDParamsWithTimeout(timeout time.Duration) *PostExpenseWorkflowWorkflowIDParams {
	return &PostExpenseWorkflowWorkflowIDParams{
		timeout: timeout,
	}
}

// NewPostExpenseWorkflowWorkflowIDParamsWithContext creates a new PostExpenseWorkflowWorkflowIDParams object
// with the ability to set a context for a request.
func NewPostExpenseWorkflowWorkflowIDParamsWithContext(ctx context.Context) *PostExpenseWorkflowWorkflowIDParams {
	return &PostExpenseWorkflowWorkflowIDParams{
		Context: ctx,
	}
}

// NewPostExpenseWorkflowWorkflowIDParamsWithHTTPClient creates a new PostExpenseWorkflowWorkflowIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExpenseWorkflowWorkflowIDParamsWithHTTPClient(client *http.Client) *PostExpenseWorkflowWorkflowIDParams {
	return &PostExpenseWorkflowWorkflowIDParams{
		HTTPClient: client,
	}
}

/* PostExpenseWorkflowWorkflowIDParams contains all the parameters to send to the API endpoint
   for the post expense workflow workflow ID operation.

   Typically these are written to a http.Request.
*/
type PostExpenseWorkflowWorkflowIDParams struct {

	/* WorkflowID.

	   ??????ID
	*/
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post expense workflow workflow ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExpenseWorkflowWorkflowIDParams) WithDefaults() *PostExpenseWorkflowWorkflowIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post expense workflow workflow ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExpenseWorkflowWorkflowIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) WithTimeout(timeout time.Duration) *PostExpenseWorkflowWorkflowIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) WithContext(ctx context.Context) *PostExpenseWorkflowWorkflowIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) WithHTTPClient(client *http.Client) *PostExpenseWorkflowWorkflowIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowID adds the workflowID to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) WithWorkflowID(workflowID string) *PostExpenseWorkflowWorkflowIDParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the post expense workflow workflow ID params
func (o *PostExpenseWorkflowWorkflowIDParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *PostExpenseWorkflowWorkflowIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workflowID
	if err := r.SetPathParam("workflowID", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

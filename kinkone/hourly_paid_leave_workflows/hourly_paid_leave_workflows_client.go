// Code generated by go-swagger; DO NOT EDIT.

package hourly_paid_leave_workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hourly paid leave workflows API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hourly paid leave workflows API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteHourlyPaidLeaveWorkflowWorkflowID(params *DeleteHourlyPaidLeaveWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHourlyPaidLeaveWorkflowWorkflowIDOK, error)

	GetHourlyPaidLeaveWorkflows(params *GetHourlyPaidLeaveWorkflowsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHourlyPaidLeaveWorkflowsOK, error)

	PostHourlyPaidLeaveWorkflowWorkflowID(params *PostHourlyPaidLeaveWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHourlyPaidLeaveWorkflowWorkflowIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteHourlyPaidLeaveWorkflowWorkflowID 時間休の申請差戻しs

  指定した申請IDに該当する時間休の申請を差戻します。
*/
func (a *Client) DeleteHourlyPaidLeaveWorkflowWorkflowID(params *DeleteHourlyPaidLeaveWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHourlyPaidLeaveWorkflowWorkflowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHourlyPaidLeaveWorkflowWorkflowIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteHourlyPaidLeaveWorkflowWorkflowID",
		Method:             "DELETE",
		PathPattern:        "/hourly_paid_leave/workflow/{workflowID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHourlyPaidLeaveWorkflowWorkflowIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHourlyPaidLeaveWorkflowWorkflowIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteHourlyPaidLeaveWorkflowWorkflowID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetHourlyPaidLeaveWorkflows 時間休の申請一覧取得s

  指定した期間の時間休の申請一覧を返します。
*/
func (a *Client) GetHourlyPaidLeaveWorkflows(params *GetHourlyPaidLeaveWorkflowsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHourlyPaidLeaveWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHourlyPaidLeaveWorkflowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHourlyPaidLeaveWorkflows",
		Method:             "GET",
		PathPattern:        "/hourly_paid_leave/workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHourlyPaidLeaveWorkflowsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHourlyPaidLeaveWorkflowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHourlyPaidLeaveWorkflows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostHourlyPaidLeaveWorkflowWorkflowID 時間休の申請承認s

  指定した申請IDに該当する時間休の申請を承認します。
*/
func (a *Client) PostHourlyPaidLeaveWorkflowWorkflowID(params *PostHourlyPaidLeaveWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHourlyPaidLeaveWorkflowWorkflowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHourlyPaidLeaveWorkflowWorkflowIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostHourlyPaidLeaveWorkflowWorkflowID",
		Method:             "POST",
		PathPattern:        "/hourly_paid_leave/workflow/{workflowID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHourlyPaidLeaveWorkflowWorkflowIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHourlyPaidLeaveWorkflowWorkflowIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHourlyPaidLeaveWorkflowWorkflowID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

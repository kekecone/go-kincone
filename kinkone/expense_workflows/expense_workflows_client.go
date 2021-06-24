// Code generated by go-swagger; DO NOT EDIT.

package expense_workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new expense workflows API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for expense workflows API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteExpenseWorkflowWorkflowID(params *DeleteExpenseWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExpenseWorkflowWorkflowIDOK, error)

	GetExpenseWorkflows(params *GetExpenseWorkflowsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExpenseWorkflowsOK, error)

	PostExpenseWorkflowWorkflowID(params *PostExpenseWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostExpenseWorkflowWorkflowIDOK, error)

	PutExpenseWorkflowWorkflowID(params *PutExpenseWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutExpenseWorkflowWorkflowIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteExpenseWorkflowWorkflowID 交通費申請差戻しs

  指定した申請IDに該当する交通費の申請を差戻します。
*/
func (a *Client) DeleteExpenseWorkflowWorkflowID(params *DeleteExpenseWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExpenseWorkflowWorkflowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExpenseWorkflowWorkflowIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteExpenseWorkflowWorkflowID",
		Method:             "DELETE",
		PathPattern:        "/expense/workflow/{workflowID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExpenseWorkflowWorkflowIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteExpenseWorkflowWorkflowIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteExpenseWorkflowWorkflowID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetExpenseWorkflows 交通費申請一覧取得s

  指定した期間の交通費申請一覧を返します。
*/
func (a *Client) GetExpenseWorkflows(params *GetExpenseWorkflowsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExpenseWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExpenseWorkflowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetExpenseWorkflows",
		Method:             "GET",
		PathPattern:        "/expense/workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExpenseWorkflowsReader{formats: a.formats},
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
	success, ok := result.(*GetExpenseWorkflowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetExpenseWorkflows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostExpenseWorkflowWorkflowID 交通費申請承認s

  指定した申請IDに該当する交通費の申請を承認します。
*/
func (a *Client) PostExpenseWorkflowWorkflowID(params *PostExpenseWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostExpenseWorkflowWorkflowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostExpenseWorkflowWorkflowIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostExpenseWorkflowWorkflowID",
		Method:             "POST",
		PathPattern:        "/expense/workflow/{workflowID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostExpenseWorkflowWorkflowIDReader{formats: a.formats},
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
	success, ok := result.(*PostExpenseWorkflowWorkflowIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostExpenseWorkflowWorkflowID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutExpenseWorkflowWorkflowID 交通費申請承認取消s

  指定した申請IDに該当する交通費の承認を取り消します。
*/
func (a *Client) PutExpenseWorkflowWorkflowID(params *PutExpenseWorkflowWorkflowIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutExpenseWorkflowWorkflowIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutExpenseWorkflowWorkflowIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutExpenseWorkflowWorkflowID",
		Method:             "PUT",
		PathPattern:        "/expense/workflow/{workflowID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutExpenseWorkflowWorkflowIDReader{formats: a.formats},
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
	success, ok := result.(*PutExpenseWorkflowWorkflowIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutExpenseWorkflowWorkflowID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

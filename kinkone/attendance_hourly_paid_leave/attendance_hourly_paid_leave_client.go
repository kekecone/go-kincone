// Code generated by go-swagger; DO NOT EDIT.

package attendance_hourly_paid_leave

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new attendance hourly paid leave API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for attendance hourly paid leave API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAttendanceHourlyPaidLeave(params *DeleteAttendanceHourlyPaidLeaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAttendanceHourlyPaidLeaveOK, error)

	PostAttendanceHourlyPaidLeave(params *PostAttendanceHourlyPaidLeaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAttendanceHourlyPaidLeaveOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAttendanceHourlyPaidLeave 時間休の削除s

  メールアドレスと日付を指定して勤怠明細に登録されている時間休を削除します。
*/
func (a *Client) DeleteAttendanceHourlyPaidLeave(params *DeleteAttendanceHourlyPaidLeaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAttendanceHourlyPaidLeaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAttendanceHourlyPaidLeaveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAttendanceHourlyPaidLeave",
		Method:             "DELETE",
		PathPattern:        "/attendance/hourly_paid_leave",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAttendanceHourlyPaidLeaveReader{formats: a.formats},
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
	success, ok := result.(*DeleteAttendanceHourlyPaidLeaveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAttendanceHourlyPaidLeave: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAttendanceHourlyPaidLeave 時間休の登録s 更新

  メールアドレスと日付を指定して勤怠明細に時間休を登録します。時間休のワークフロー利用が ON になっている場合は指定できません。
*/
func (a *Client) PostAttendanceHourlyPaidLeave(params *PostAttendanceHourlyPaidLeaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAttendanceHourlyPaidLeaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAttendanceHourlyPaidLeaveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAttendanceHourlyPaidLeave",
		Method:             "POST",
		PathPattern:        "/attendance/hourly_paid_leave",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAttendanceHourlyPaidLeaveReader{formats: a.formats},
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
	success, ok := result.(*PostAttendanceHourlyPaidLeaveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAttendanceHourlyPaidLeave: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

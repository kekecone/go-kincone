// Code generated by go-swagger; DO NOT EDIT.

package attendance_workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kekecone/go-kincone/kinkone"
)

// PostAttendanceWorkflowWorkflowIDReader is a Reader for the PostAttendanceWorkflowWorkflowID structure.
type PostAttendanceWorkflowWorkflowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAttendanceWorkflowWorkflowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAttendanceWorkflowWorkflowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostAttendanceWorkflowWorkflowIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAttendanceWorkflowWorkflowIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAttendanceWorkflowWorkflowIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAttendanceWorkflowWorkflowIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAttendanceWorkflowWorkflowIDOK creates a PostAttendanceWorkflowWorkflowIDOK with default headers values
func NewPostAttendanceWorkflowWorkflowIDOK() *PostAttendanceWorkflowWorkflowIDOK {
	return &PostAttendanceWorkflowWorkflowIDOK{}
}

/* PostAttendanceWorkflowWorkflowIDOK describes a response with status code 200, with default header values.

正常時
*/
type PostAttendanceWorkflowWorkflowIDOK struct {
	Payload *PostAttendanceWorkflowWorkflowIDOKBody
}

func (o *PostAttendanceWorkflowWorkflowIDOK) Error() string {
	return fmt.Sprintf("[POST /attendance/workflow/{workflowID}][%d] postAttendanceWorkflowWorkflowIdOK  %+v", 200, o.Payload)
}
func (o *PostAttendanceWorkflowWorkflowIDOK) GetPayload() *PostAttendanceWorkflowWorkflowIDOKBody {
	return o.Payload
}

func (o *PostAttendanceWorkflowWorkflowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAttendanceWorkflowWorkflowIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAttendanceWorkflowWorkflowIDUnauthorized creates a PostAttendanceWorkflowWorkflowIDUnauthorized with default headers values
func NewPostAttendanceWorkflowWorkflowIDUnauthorized() *PostAttendanceWorkflowWorkflowIDUnauthorized {
	return &PostAttendanceWorkflowWorkflowIDUnauthorized{}
}

/* PostAttendanceWorkflowWorkflowIDUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type PostAttendanceWorkflowWorkflowIDUnauthorized struct {
	Payload *kinkone.ErrorModel
}

func (o *PostAttendanceWorkflowWorkflowIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /attendance/workflow/{workflowID}][%d] postAttendanceWorkflowWorkflowIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PostAttendanceWorkflowWorkflowIDUnauthorized) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PostAttendanceWorkflowWorkflowIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAttendanceWorkflowWorkflowIDForbidden creates a PostAttendanceWorkflowWorkflowIDForbidden with default headers values
func NewPostAttendanceWorkflowWorkflowIDForbidden() *PostAttendanceWorkflowWorkflowIDForbidden {
	return &PostAttendanceWorkflowWorkflowIDForbidden{}
}

/* PostAttendanceWorkflowWorkflowIDForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type PostAttendanceWorkflowWorkflowIDForbidden struct {
	Payload *kinkone.ErrorModel
}

func (o *PostAttendanceWorkflowWorkflowIDForbidden) Error() string {
	return fmt.Sprintf("[POST /attendance/workflow/{workflowID}][%d] postAttendanceWorkflowWorkflowIdForbidden  %+v", 403, o.Payload)
}
func (o *PostAttendanceWorkflowWorkflowIDForbidden) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PostAttendanceWorkflowWorkflowIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAttendanceWorkflowWorkflowIDUnprocessableEntity creates a PostAttendanceWorkflowWorkflowIDUnprocessableEntity with default headers values
func NewPostAttendanceWorkflowWorkflowIDUnprocessableEntity() *PostAttendanceWorkflowWorkflowIDUnprocessableEntity {
	return &PostAttendanceWorkflowWorkflowIDUnprocessableEntity{}
}

/* PostAttendanceWorkflowWorkflowIDUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type PostAttendanceWorkflowWorkflowIDUnprocessableEntity struct {
	Payload *kinkone.ErrorModel
}

func (o *PostAttendanceWorkflowWorkflowIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /attendance/workflow/{workflowID}][%d] postAttendanceWorkflowWorkflowIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PostAttendanceWorkflowWorkflowIDUnprocessableEntity) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PostAttendanceWorkflowWorkflowIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAttendanceWorkflowWorkflowIDInternalServerError creates a PostAttendanceWorkflowWorkflowIDInternalServerError with default headers values
func NewPostAttendanceWorkflowWorkflowIDInternalServerError() *PostAttendanceWorkflowWorkflowIDInternalServerError {
	return &PostAttendanceWorkflowWorkflowIDInternalServerError{}
}

/* PostAttendanceWorkflowWorkflowIDInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type PostAttendanceWorkflowWorkflowIDInternalServerError struct {
	Payload *kinkone.ErrorModel
}

func (o *PostAttendanceWorkflowWorkflowIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /attendance/workflow/{workflowID}][%d] postAttendanceWorkflowWorkflowIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAttendanceWorkflowWorkflowIDInternalServerError) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PostAttendanceWorkflowWorkflowIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostAttendanceWorkflowWorkflowIDOKBody post attendance workflow workflow ID o k body
swagger:model PostAttendanceWorkflowWorkflowIDOKBody
*/
type PostAttendanceWorkflowWorkflowIDOKBody struct {

	// response
	Response *PostAttendanceWorkflowWorkflowIDOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this post attendance workflow workflow ID o k body
func (o *PostAttendanceWorkflowWorkflowIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAttendanceWorkflowWorkflowIDOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAttendanceWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post attendance workflow workflow ID o k body based on the context it is used
func (o *PostAttendanceWorkflowWorkflowIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAttendanceWorkflowWorkflowIDOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAttendanceWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAttendanceWorkflowWorkflowIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAttendanceWorkflowWorkflowIDOKBody) UnmarshalBinary(b []byte) error {
	var res PostAttendanceWorkflowWorkflowIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostAttendanceWorkflowWorkflowIDOKBodyResponse APIレスポンス
swagger:model PostAttendanceWorkflowWorkflowIDOKBodyResponse
*/
type PostAttendanceWorkflowWorkflowIDOKBodyResponse struct {

	// 処理結果
	Result bool `json:"result,omitempty"`
}

// Validate validates this post attendance workflow workflow ID o k body response
func (o *PostAttendanceWorkflowWorkflowIDOKBodyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post attendance workflow workflow ID o k body response based on context it is used
func (o *PostAttendanceWorkflowWorkflowIDOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAttendanceWorkflowWorkflowIDOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAttendanceWorkflowWorkflowIDOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res PostAttendanceWorkflowWorkflowIDOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package punch

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

	"github.com/kekecone/go-kincone/kinkone/model"
)

// PostPunchInReader is a Reader for the PostPunchIn structure.
type PostPunchInReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPunchInReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPunchInOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostPunchInUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPunchInForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostPunchInUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPunchInInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostPunchInOK creates a PostPunchInOK with default headers values
func NewPostPunchInOK() *PostPunchInOK {
	return &PostPunchInOK{}
}

/* PostPunchInOK describes a response with status code 200, with default header values.

正常時
*/
type PostPunchInOK struct {
	Payload *PostPunchInOKBody
}

func (o *PostPunchInOK) Error() string {
	return fmt.Sprintf("[POST /punch/in][%d] postPunchInOK  %+v", 200, o.Payload)
}
func (o *PostPunchInOK) GetPayload() *PostPunchInOKBody {
	return o.Payload
}

func (o *PostPunchInOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPunchInOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPunchInUnauthorized creates a PostPunchInUnauthorized with default headers values
func NewPostPunchInUnauthorized() *PostPunchInUnauthorized {
	return &PostPunchInUnauthorized{}
}

/* PostPunchInUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type PostPunchInUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *PostPunchInUnauthorized) Error() string {
	return fmt.Sprintf("[POST /punch/in][%d] postPunchInUnauthorized  %+v", 401, o.Payload)
}
func (o *PostPunchInUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostPunchInUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPunchInForbidden creates a PostPunchInForbidden with default headers values
func NewPostPunchInForbidden() *PostPunchInForbidden {
	return &PostPunchInForbidden{}
}

/* PostPunchInForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type PostPunchInForbidden struct {
	Payload *model.ErrorModel
}

func (o *PostPunchInForbidden) Error() string {
	return fmt.Sprintf("[POST /punch/in][%d] postPunchInForbidden  %+v", 403, o.Payload)
}
func (o *PostPunchInForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostPunchInForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPunchInUnprocessableEntity creates a PostPunchInUnprocessableEntity with default headers values
func NewPostPunchInUnprocessableEntity() *PostPunchInUnprocessableEntity {
	return &PostPunchInUnprocessableEntity{}
}

/* PostPunchInUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type PostPunchInUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *PostPunchInUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /punch/in][%d] postPunchInUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PostPunchInUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostPunchInUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPunchInInternalServerError creates a PostPunchInInternalServerError with default headers values
func NewPostPunchInInternalServerError() *PostPunchInInternalServerError {
	return &PostPunchInInternalServerError{}
}

/* PostPunchInInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type PostPunchInInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *PostPunchInInternalServerError) Error() string {
	return fmt.Sprintf("[POST /punch/in][%d] postPunchInInternalServerError  %+v", 500, o.Payload)
}
func (o *PostPunchInInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostPunchInInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostPunchInOKBody post punch in o k body
swagger:model PostPunchInOKBody
*/
type PostPunchInOKBody struct {

	// response
	Response *PostPunchInOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this post punch in o k body
func (o *PostPunchInOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPunchInOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPunchInOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post punch in o k body based on the context it is used
func (o *PostPunchInOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPunchInOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPunchInOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPunchInOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPunchInOKBody) UnmarshalBinary(b []byte) error {
	var res PostPunchInOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPunchInOKBodyResponse APIレスポンス
swagger:model PostPunchInOKBodyResponse
*/
type PostPunchInOKBodyResponse struct {

	// attendance
	Attendance *model.AttendanceModel `json:"attendance,omitempty"`
}

// Validate validates this post punch in o k body response
func (o *PostPunchInOKBodyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttendance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPunchInOKBodyResponse) validateAttendance(formats strfmt.Registry) error {
	if swag.IsZero(o.Attendance) { // not required
		return nil
	}

	if o.Attendance != nil {
		if err := o.Attendance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPunchInOK" + "." + "response" + "." + "attendance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post punch in o k body response based on the context it is used
func (o *PostPunchInOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttendance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPunchInOKBodyResponse) contextValidateAttendance(ctx context.Context, formats strfmt.Registry) error {

	if o.Attendance != nil {
		if err := o.Attendance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPunchInOK" + "." + "response" + "." + "attendance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPunchInOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPunchInOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res PostPunchInOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

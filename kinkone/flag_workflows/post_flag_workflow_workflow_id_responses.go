// Code generated by go-swagger; DO NOT EDIT.

package flag_workflows

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

// PostFlagWorkflowWorkflowIDReader is a Reader for the PostFlagWorkflowWorkflowID structure.
type PostFlagWorkflowWorkflowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlagWorkflowWorkflowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlagWorkflowWorkflowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostFlagWorkflowWorkflowIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlagWorkflowWorkflowIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostFlagWorkflowWorkflowIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlagWorkflowWorkflowIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlagWorkflowWorkflowIDOK creates a PostFlagWorkflowWorkflowIDOK with default headers values
func NewPostFlagWorkflowWorkflowIDOK() *PostFlagWorkflowWorkflowIDOK {
	return &PostFlagWorkflowWorkflowIDOK{}
}

/* PostFlagWorkflowWorkflowIDOK describes a response with status code 200, with default header values.

正常時
*/
type PostFlagWorkflowWorkflowIDOK struct {
	Payload *PostFlagWorkflowWorkflowIDOKBody
}

func (o *PostFlagWorkflowWorkflowIDOK) Error() string {
	return fmt.Sprintf("[POST /flag/workflow/{workflowID}][%d] postFlagWorkflowWorkflowIdOK  %+v", 200, o.Payload)
}
func (o *PostFlagWorkflowWorkflowIDOK) GetPayload() *PostFlagWorkflowWorkflowIDOKBody {
	return o.Payload
}

func (o *PostFlagWorkflowWorkflowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostFlagWorkflowWorkflowIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlagWorkflowWorkflowIDUnauthorized creates a PostFlagWorkflowWorkflowIDUnauthorized with default headers values
func NewPostFlagWorkflowWorkflowIDUnauthorized() *PostFlagWorkflowWorkflowIDUnauthorized {
	return &PostFlagWorkflowWorkflowIDUnauthorized{}
}

/* PostFlagWorkflowWorkflowIDUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type PostFlagWorkflowWorkflowIDUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *PostFlagWorkflowWorkflowIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /flag/workflow/{workflowID}][%d] postFlagWorkflowWorkflowIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PostFlagWorkflowWorkflowIDUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostFlagWorkflowWorkflowIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlagWorkflowWorkflowIDForbidden creates a PostFlagWorkflowWorkflowIDForbidden with default headers values
func NewPostFlagWorkflowWorkflowIDForbidden() *PostFlagWorkflowWorkflowIDForbidden {
	return &PostFlagWorkflowWorkflowIDForbidden{}
}

/* PostFlagWorkflowWorkflowIDForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type PostFlagWorkflowWorkflowIDForbidden struct {
	Payload *model.ErrorModel
}

func (o *PostFlagWorkflowWorkflowIDForbidden) Error() string {
	return fmt.Sprintf("[POST /flag/workflow/{workflowID}][%d] postFlagWorkflowWorkflowIdForbidden  %+v", 403, o.Payload)
}
func (o *PostFlagWorkflowWorkflowIDForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostFlagWorkflowWorkflowIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlagWorkflowWorkflowIDUnprocessableEntity creates a PostFlagWorkflowWorkflowIDUnprocessableEntity with default headers values
func NewPostFlagWorkflowWorkflowIDUnprocessableEntity() *PostFlagWorkflowWorkflowIDUnprocessableEntity {
	return &PostFlagWorkflowWorkflowIDUnprocessableEntity{}
}

/* PostFlagWorkflowWorkflowIDUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type PostFlagWorkflowWorkflowIDUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *PostFlagWorkflowWorkflowIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /flag/workflow/{workflowID}][%d] postFlagWorkflowWorkflowIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PostFlagWorkflowWorkflowIDUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostFlagWorkflowWorkflowIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlagWorkflowWorkflowIDInternalServerError creates a PostFlagWorkflowWorkflowIDInternalServerError with default headers values
func NewPostFlagWorkflowWorkflowIDInternalServerError() *PostFlagWorkflowWorkflowIDInternalServerError {
	return &PostFlagWorkflowWorkflowIDInternalServerError{}
}

/* PostFlagWorkflowWorkflowIDInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type PostFlagWorkflowWorkflowIDInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *PostFlagWorkflowWorkflowIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /flag/workflow/{workflowID}][%d] postFlagWorkflowWorkflowIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PostFlagWorkflowWorkflowIDInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PostFlagWorkflowWorkflowIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostFlagWorkflowWorkflowIDOKBody post flag workflow workflow ID o k body
swagger:model PostFlagWorkflowWorkflowIDOKBody
*/
type PostFlagWorkflowWorkflowIDOKBody struct {

	// response
	Response *PostFlagWorkflowWorkflowIDOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this post flag workflow workflow ID o k body
func (o *PostFlagWorkflowWorkflowIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFlagWorkflowWorkflowIDOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postFlagWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post flag workflow workflow ID o k body based on the context it is used
func (o *PostFlagWorkflowWorkflowIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFlagWorkflowWorkflowIDOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postFlagWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostFlagWorkflowWorkflowIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFlagWorkflowWorkflowIDOKBody) UnmarshalBinary(b []byte) error {
	var res PostFlagWorkflowWorkflowIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostFlagWorkflowWorkflowIDOKBodyResponse APIレスポンス
swagger:model PostFlagWorkflowWorkflowIDOKBodyResponse
*/
type PostFlagWorkflowWorkflowIDOKBodyResponse struct {

	// 処理結果
	Result bool `json:"result,omitempty"`
}

// Validate validates this post flag workflow workflow ID o k body response
func (o *PostFlagWorkflowWorkflowIDOKBodyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post flag workflow workflow ID o k body response based on context it is used
func (o *PostFlagWorkflowWorkflowIDOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFlagWorkflowWorkflowIDOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFlagWorkflowWorkflowIDOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res PostFlagWorkflowWorkflowIDOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

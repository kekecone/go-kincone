// Code generated by go-swagger; DO NOT EDIT.

package expense_workflows

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

// PutExpenseWorkflowWorkflowIDReader is a Reader for the PutExpenseWorkflowWorkflowID structure.
type PutExpenseWorkflowWorkflowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutExpenseWorkflowWorkflowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutExpenseWorkflowWorkflowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutExpenseWorkflowWorkflowIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutExpenseWorkflowWorkflowIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutExpenseWorkflowWorkflowIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutExpenseWorkflowWorkflowIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutExpenseWorkflowWorkflowIDOK creates a PutExpenseWorkflowWorkflowIDOK with default headers values
func NewPutExpenseWorkflowWorkflowIDOK() *PutExpenseWorkflowWorkflowIDOK {
	return &PutExpenseWorkflowWorkflowIDOK{}
}

/* PutExpenseWorkflowWorkflowIDOK describes a response with status code 200, with default header values.

正常時
*/
type PutExpenseWorkflowWorkflowIDOK struct {
	Payload *PutExpenseWorkflowWorkflowIDOKBody
}

func (o *PutExpenseWorkflowWorkflowIDOK) Error() string {
	return fmt.Sprintf("[PUT /expense/workflow/{workflowID}][%d] putExpenseWorkflowWorkflowIdOK  %+v", 200, o.Payload)
}
func (o *PutExpenseWorkflowWorkflowIDOK) GetPayload() *PutExpenseWorkflowWorkflowIDOKBody {
	return o.Payload
}

func (o *PutExpenseWorkflowWorkflowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutExpenseWorkflowWorkflowIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExpenseWorkflowWorkflowIDUnauthorized creates a PutExpenseWorkflowWorkflowIDUnauthorized with default headers values
func NewPutExpenseWorkflowWorkflowIDUnauthorized() *PutExpenseWorkflowWorkflowIDUnauthorized {
	return &PutExpenseWorkflowWorkflowIDUnauthorized{}
}

/* PutExpenseWorkflowWorkflowIDUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type PutExpenseWorkflowWorkflowIDUnauthorized struct {
	Payload *kinkone.ErrorModel
}

func (o *PutExpenseWorkflowWorkflowIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /expense/workflow/{workflowID}][%d] putExpenseWorkflowWorkflowIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PutExpenseWorkflowWorkflowIDUnauthorized) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PutExpenseWorkflowWorkflowIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExpenseWorkflowWorkflowIDForbidden creates a PutExpenseWorkflowWorkflowIDForbidden with default headers values
func NewPutExpenseWorkflowWorkflowIDForbidden() *PutExpenseWorkflowWorkflowIDForbidden {
	return &PutExpenseWorkflowWorkflowIDForbidden{}
}

/* PutExpenseWorkflowWorkflowIDForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type PutExpenseWorkflowWorkflowIDForbidden struct {
	Payload *kinkone.ErrorModel
}

func (o *PutExpenseWorkflowWorkflowIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /expense/workflow/{workflowID}][%d] putExpenseWorkflowWorkflowIdForbidden  %+v", 403, o.Payload)
}
func (o *PutExpenseWorkflowWorkflowIDForbidden) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PutExpenseWorkflowWorkflowIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExpenseWorkflowWorkflowIDUnprocessableEntity creates a PutExpenseWorkflowWorkflowIDUnprocessableEntity with default headers values
func NewPutExpenseWorkflowWorkflowIDUnprocessableEntity() *PutExpenseWorkflowWorkflowIDUnprocessableEntity {
	return &PutExpenseWorkflowWorkflowIDUnprocessableEntity{}
}

/* PutExpenseWorkflowWorkflowIDUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type PutExpenseWorkflowWorkflowIDUnprocessableEntity struct {
	Payload *kinkone.ErrorModel
}

func (o *PutExpenseWorkflowWorkflowIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /expense/workflow/{workflowID}][%d] putExpenseWorkflowWorkflowIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PutExpenseWorkflowWorkflowIDUnprocessableEntity) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PutExpenseWorkflowWorkflowIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExpenseWorkflowWorkflowIDInternalServerError creates a PutExpenseWorkflowWorkflowIDInternalServerError with default headers values
func NewPutExpenseWorkflowWorkflowIDInternalServerError() *PutExpenseWorkflowWorkflowIDInternalServerError {
	return &PutExpenseWorkflowWorkflowIDInternalServerError{}
}

/* PutExpenseWorkflowWorkflowIDInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type PutExpenseWorkflowWorkflowIDInternalServerError struct {
	Payload *kinkone.ErrorModel
}

func (o *PutExpenseWorkflowWorkflowIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /expense/workflow/{workflowID}][%d] putExpenseWorkflowWorkflowIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PutExpenseWorkflowWorkflowIDInternalServerError) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *PutExpenseWorkflowWorkflowIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutExpenseWorkflowWorkflowIDOKBody put expense workflow workflow ID o k body
swagger:model PutExpenseWorkflowWorkflowIDOKBody
*/
type PutExpenseWorkflowWorkflowIDOKBody struct {

	// response
	Response *PutExpenseWorkflowWorkflowIDOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this put expense workflow workflow ID o k body
func (o *PutExpenseWorkflowWorkflowIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutExpenseWorkflowWorkflowIDOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putExpenseWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put expense workflow workflow ID o k body based on the context it is used
func (o *PutExpenseWorkflowWorkflowIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutExpenseWorkflowWorkflowIDOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putExpenseWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutExpenseWorkflowWorkflowIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutExpenseWorkflowWorkflowIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutExpenseWorkflowWorkflowIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutExpenseWorkflowWorkflowIDOKBodyResponse APIレスポンス
swagger:model PutExpenseWorkflowWorkflowIDOKBodyResponse
*/
type PutExpenseWorkflowWorkflowIDOKBodyResponse struct {

	// 処理結果
	Result bool `json:"result,omitempty"`
}

// Validate validates this put expense workflow workflow ID o k body response
func (o *PutExpenseWorkflowWorkflowIDOKBodyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put expense workflow workflow ID o k body response based on context it is used
func (o *PutExpenseWorkflowWorkflowIDOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutExpenseWorkflowWorkflowIDOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutExpenseWorkflowWorkflowIDOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res PutExpenseWorkflowWorkflowIDOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

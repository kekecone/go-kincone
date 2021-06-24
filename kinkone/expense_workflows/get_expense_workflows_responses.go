// Code generated by go-swagger; DO NOT EDIT.

package expense_workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kekecone/go-kincone/kinkone/model"
)

// GetExpenseWorkflowsReader is a Reader for the GetExpenseWorkflows structure.
type GetExpenseWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExpenseWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExpenseWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExpenseWorkflowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExpenseWorkflowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetExpenseWorkflowsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExpenseWorkflowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExpenseWorkflowsOK creates a GetExpenseWorkflowsOK with default headers values
func NewGetExpenseWorkflowsOK() *GetExpenseWorkflowsOK {
	return &GetExpenseWorkflowsOK{}
}

/* GetExpenseWorkflowsOK describes a response with status code 200, with default header values.

正常時
*/
type GetExpenseWorkflowsOK struct {
	Payload *GetExpenseWorkflowsOKBody
}

func (o *GetExpenseWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /expense/workflows][%d] getExpenseWorkflowsOK  %+v", 200, o.Payload)
}
func (o *GetExpenseWorkflowsOK) GetPayload() *GetExpenseWorkflowsOKBody {
	return o.Payload
}

func (o *GetExpenseWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetExpenseWorkflowsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExpenseWorkflowsUnauthorized creates a GetExpenseWorkflowsUnauthorized with default headers values
func NewGetExpenseWorkflowsUnauthorized() *GetExpenseWorkflowsUnauthorized {
	return &GetExpenseWorkflowsUnauthorized{}
}

/* GetExpenseWorkflowsUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type GetExpenseWorkflowsUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *GetExpenseWorkflowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /expense/workflows][%d] getExpenseWorkflowsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetExpenseWorkflowsUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetExpenseWorkflowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExpenseWorkflowsForbidden creates a GetExpenseWorkflowsForbidden with default headers values
func NewGetExpenseWorkflowsForbidden() *GetExpenseWorkflowsForbidden {
	return &GetExpenseWorkflowsForbidden{}
}

/* GetExpenseWorkflowsForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type GetExpenseWorkflowsForbidden struct {
	Payload *model.ErrorModel
}

func (o *GetExpenseWorkflowsForbidden) Error() string {
	return fmt.Sprintf("[GET /expense/workflows][%d] getExpenseWorkflowsForbidden  %+v", 403, o.Payload)
}
func (o *GetExpenseWorkflowsForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetExpenseWorkflowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExpenseWorkflowsUnprocessableEntity creates a GetExpenseWorkflowsUnprocessableEntity with default headers values
func NewGetExpenseWorkflowsUnprocessableEntity() *GetExpenseWorkflowsUnprocessableEntity {
	return &GetExpenseWorkflowsUnprocessableEntity{}
}

/* GetExpenseWorkflowsUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type GetExpenseWorkflowsUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *GetExpenseWorkflowsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /expense/workflows][%d] getExpenseWorkflowsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetExpenseWorkflowsUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetExpenseWorkflowsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExpenseWorkflowsInternalServerError creates a GetExpenseWorkflowsInternalServerError with default headers values
func NewGetExpenseWorkflowsInternalServerError() *GetExpenseWorkflowsInternalServerError {
	return &GetExpenseWorkflowsInternalServerError{}
}

/* GetExpenseWorkflowsInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type GetExpenseWorkflowsInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *GetExpenseWorkflowsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /expense/workflows][%d] getExpenseWorkflowsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetExpenseWorkflowsInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetExpenseWorkflowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetExpenseWorkflowsOKBody get expense workflows o k body
swagger:model GetExpenseWorkflowsOKBody
*/
type GetExpenseWorkflowsOKBody struct {

	// response
	Response *GetExpenseWorkflowsOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this get expense workflows o k body
func (o *GetExpenseWorkflowsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetExpenseWorkflowsOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getExpenseWorkflowsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get expense workflows o k body based on the context it is used
func (o *GetExpenseWorkflowsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetExpenseWorkflowsOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getExpenseWorkflowsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetExpenseWorkflowsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetExpenseWorkflowsOKBody) UnmarshalBinary(b []byte) error {
	var res GetExpenseWorkflowsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetExpenseWorkflowsOKBodyResponse APIレスポンス
swagger:model GetExpenseWorkflowsOKBodyResponse
*/
type GetExpenseWorkflowsOKBodyResponse struct {

	// workflows
	Workflows []*model.ExpenseWorkflowModel `json:"workflows"`
}

// Validate validates this get expense workflows o k body response
func (o *GetExpenseWorkflowsOKBodyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetExpenseWorkflowsOKBodyResponse) validateWorkflows(formats strfmt.Registry) error {
	if swag.IsZero(o.Workflows) { // not required
		return nil
	}

	for i := 0; i < len(o.Workflows); i++ {
		if swag.IsZero(o.Workflows[i]) { // not required
			continue
		}

		if o.Workflows[i] != nil {
			if err := o.Workflows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getExpenseWorkflowsOK" + "." + "response" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get expense workflows o k body response based on the context it is used
func (o *GetExpenseWorkflowsOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateWorkflows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetExpenseWorkflowsOKBodyResponse) contextValidateWorkflows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Workflows); i++ {

		if o.Workflows[i] != nil {
			if err := o.Workflows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getExpenseWorkflowsOK" + "." + "response" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetExpenseWorkflowsOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetExpenseWorkflowsOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res GetExpenseWorkflowsOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

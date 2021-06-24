// Code generated by go-swagger; DO NOT EDIT.

package flag_workflows

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

	"github.com/kekecone/go-kincone/kinkone"
)

// GetFlagWorkflowsReader is a Reader for the GetFlagWorkflows structure.
type GetFlagWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlagWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlagWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFlagWorkflowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlagWorkflowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetFlagWorkflowsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlagWorkflowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlagWorkflowsOK creates a GetFlagWorkflowsOK with default headers values
func NewGetFlagWorkflowsOK() *GetFlagWorkflowsOK {
	return &GetFlagWorkflowsOK{}
}

/* GetFlagWorkflowsOK describes a response with status code 200, with default header values.

正常時
*/
type GetFlagWorkflowsOK struct {
	Payload *GetFlagWorkflowsOKBody
}

func (o *GetFlagWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /flag/workflows][%d] getFlagWorkflowsOK  %+v", 200, o.Payload)
}
func (o *GetFlagWorkflowsOK) GetPayload() *GetFlagWorkflowsOKBody {
	return o.Payload
}

func (o *GetFlagWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFlagWorkflowsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlagWorkflowsUnauthorized creates a GetFlagWorkflowsUnauthorized with default headers values
func NewGetFlagWorkflowsUnauthorized() *GetFlagWorkflowsUnauthorized {
	return &GetFlagWorkflowsUnauthorized{}
}

/* GetFlagWorkflowsUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type GetFlagWorkflowsUnauthorized struct {
	Payload *kinkone.ErrorModel
}

func (o *GetFlagWorkflowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flag/workflows][%d] getFlagWorkflowsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFlagWorkflowsUnauthorized) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *GetFlagWorkflowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlagWorkflowsForbidden creates a GetFlagWorkflowsForbidden with default headers values
func NewGetFlagWorkflowsForbidden() *GetFlagWorkflowsForbidden {
	return &GetFlagWorkflowsForbidden{}
}

/* GetFlagWorkflowsForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type GetFlagWorkflowsForbidden struct {
	Payload *kinkone.ErrorModel
}

func (o *GetFlagWorkflowsForbidden) Error() string {
	return fmt.Sprintf("[GET /flag/workflows][%d] getFlagWorkflowsForbidden  %+v", 403, o.Payload)
}
func (o *GetFlagWorkflowsForbidden) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *GetFlagWorkflowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlagWorkflowsUnprocessableEntity creates a GetFlagWorkflowsUnprocessableEntity with default headers values
func NewGetFlagWorkflowsUnprocessableEntity() *GetFlagWorkflowsUnprocessableEntity {
	return &GetFlagWorkflowsUnprocessableEntity{}
}

/* GetFlagWorkflowsUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type GetFlagWorkflowsUnprocessableEntity struct {
	Payload *kinkone.ErrorModel
}

func (o *GetFlagWorkflowsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /flag/workflows][%d] getFlagWorkflowsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetFlagWorkflowsUnprocessableEntity) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *GetFlagWorkflowsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlagWorkflowsInternalServerError creates a GetFlagWorkflowsInternalServerError with default headers values
func NewGetFlagWorkflowsInternalServerError() *GetFlagWorkflowsInternalServerError {
	return &GetFlagWorkflowsInternalServerError{}
}

/* GetFlagWorkflowsInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type GetFlagWorkflowsInternalServerError struct {
	Payload *kinkone.ErrorModel
}

func (o *GetFlagWorkflowsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /flag/workflows][%d] getFlagWorkflowsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFlagWorkflowsInternalServerError) GetPayload() *kinkone.ErrorModel {
	return o.Payload
}

func (o *GetFlagWorkflowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kinkone.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFlagWorkflowsOKBody get flag workflows o k body
swagger:model GetFlagWorkflowsOKBody
*/
type GetFlagWorkflowsOKBody struct {

	// response
	Response *GetFlagWorkflowsOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this get flag workflows o k body
func (o *GetFlagWorkflowsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFlagWorkflowsOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFlagWorkflowsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get flag workflows o k body based on the context it is used
func (o *GetFlagWorkflowsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFlagWorkflowsOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFlagWorkflowsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFlagWorkflowsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFlagWorkflowsOKBody) UnmarshalBinary(b []byte) error {
	var res GetFlagWorkflowsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFlagWorkflowsOKBodyResponse APIレスポンス
swagger:model GetFlagWorkflowsOKBodyResponse
*/
type GetFlagWorkflowsOKBodyResponse struct {

	// workflows
	Workflows []*kinkone.FlagWorkflowModel `json:"workflows"`
}

// Validate validates this get flag workflows o k body response
func (o *GetFlagWorkflowsOKBodyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFlagWorkflowsOKBodyResponse) validateWorkflows(formats strfmt.Registry) error {
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
					return ve.ValidateName("getFlagWorkflowsOK" + "." + "response" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get flag workflows o k body response based on the context it is used
func (o *GetFlagWorkflowsOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateWorkflows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFlagWorkflowsOKBodyResponse) contextValidateWorkflows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Workflows); i++ {

		if o.Workflows[i] != nil {
			if err := o.Workflows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getFlagWorkflowsOK" + "." + "response" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFlagWorkflowsOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFlagWorkflowsOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res GetFlagWorkflowsOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

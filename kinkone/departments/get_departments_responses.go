// Code generated by go-swagger; DO NOT EDIT.

package departments

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

// GetDepartmentsReader is a Reader for the GetDepartments structure.
type GetDepartmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDepartmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDepartmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDepartmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDepartmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetDepartmentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDepartmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDepartmentsOK creates a GetDepartmentsOK with default headers values
func NewGetDepartmentsOK() *GetDepartmentsOK {
	return &GetDepartmentsOK{}
}

/* GetDepartmentsOK describes a response with status code 200, with default header values.

正常時
*/
type GetDepartmentsOK struct {
	Payload *GetDepartmentsOKBody
}

func (o *GetDepartmentsOK) Error() string {
	return fmt.Sprintf("[GET /departments][%d] getDepartmentsOK  %+v", 200, o.Payload)
}
func (o *GetDepartmentsOK) GetPayload() *GetDepartmentsOKBody {
	return o.Payload
}

func (o *GetDepartmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDepartmentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDepartmentsUnauthorized creates a GetDepartmentsUnauthorized with default headers values
func NewGetDepartmentsUnauthorized() *GetDepartmentsUnauthorized {
	return &GetDepartmentsUnauthorized{}
}

/* GetDepartmentsUnauthorized describes a response with status code 401, with default header values.

APIトークンの値を確認して下さい。
*/
type GetDepartmentsUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *GetDepartmentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /departments][%d] getDepartmentsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDepartmentsUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetDepartmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDepartmentsForbidden creates a GetDepartmentsForbidden with default headers values
func NewGetDepartmentsForbidden() *GetDepartmentsForbidden {
	return &GetDepartmentsForbidden{}
}

/* GetDepartmentsForbidden describes a response with status code 403, with default header values.

APIにアクセスする権限がありません。
*/
type GetDepartmentsForbidden struct {
	Payload *model.ErrorModel
}

func (o *GetDepartmentsForbidden) Error() string {
	return fmt.Sprintf("[GET /departments][%d] getDepartmentsForbidden  %+v", 403, o.Payload)
}
func (o *GetDepartmentsForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetDepartmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDepartmentsUnprocessableEntity creates a GetDepartmentsUnprocessableEntity with default headers values
func NewGetDepartmentsUnprocessableEntity() *GetDepartmentsUnprocessableEntity {
	return &GetDepartmentsUnprocessableEntity{}
}

/* GetDepartmentsUnprocessableEntity describes a response with status code 422, with default header values.

パラメータの値を確認して下さい。
*/
type GetDepartmentsUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *GetDepartmentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /departments][%d] getDepartmentsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetDepartmentsUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetDepartmentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDepartmentsInternalServerError creates a GetDepartmentsInternalServerError with default headers values
func NewGetDepartmentsInternalServerError() *GetDepartmentsInternalServerError {
	return &GetDepartmentsInternalServerError{}
}

/* GetDepartmentsInternalServerError describes a response with status code 500, with default header values.

APIサーバーの内部的なエラーです。
*/
type GetDepartmentsInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *GetDepartmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /departments][%d] getDepartmentsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDepartmentsInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetDepartmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDepartmentsOKBody get departments o k body
swagger:model GetDepartmentsOKBody
*/
type GetDepartmentsOKBody struct {

	// response
	Response *GetDepartmentsOKBodyResponse `json:"response,omitempty"`

	// 結果コード
	ResultCode string `json:"result_code,omitempty"`

	// 結果ステータス
	Status bool `json:"status,omitempty"`
}

// Validate validates this get departments o k body
func (o *GetDepartmentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDepartmentsOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDepartmentsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get departments o k body based on the context it is used
func (o *GetDepartmentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDepartmentsOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getDepartmentsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDepartmentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDepartmentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDepartmentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDepartmentsOKBodyResponse APIレスポンス
swagger:model GetDepartmentsOKBodyResponse
*/
type GetDepartmentsOKBodyResponse struct {

	// departments
	Departments []*model.DepartmentModel `json:"departments"`
}

// Validate validates this get departments o k body response
func (o *GetDepartmentsOKBodyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDepartments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDepartmentsOKBodyResponse) validateDepartments(formats strfmt.Registry) error {
	if swag.IsZero(o.Departments) { // not required
		return nil
	}

	for i := 0; i < len(o.Departments); i++ {
		if swag.IsZero(o.Departments[i]) { // not required
			continue
		}

		if o.Departments[i] != nil {
			if err := o.Departments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDepartmentsOK" + "." + "response" + "." + "departments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get departments o k body response based on the context it is used
func (o *GetDepartmentsOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDepartments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDepartmentsOKBodyResponse) contextValidateDepartments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Departments); i++ {

		if o.Departments[i] != nil {
			if err := o.Departments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDepartmentsOK" + "." + "response" + "." + "departments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDepartmentsOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDepartmentsOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res GetDepartmentsOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

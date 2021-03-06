// Code generated by go-swagger; DO NOT EDIT.

package compensation_workflows

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

// GetCompensationWorkflowsReader is a Reader for the GetCompensationWorkflows structure.
type GetCompensationWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompensationWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCompensationWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCompensationWorkflowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCompensationWorkflowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCompensationWorkflowsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCompensationWorkflowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCompensationWorkflowsOK creates a GetCompensationWorkflowsOK with default headers values
func NewGetCompensationWorkflowsOK() *GetCompensationWorkflowsOK {
	return &GetCompensationWorkflowsOK{}
}

/* GetCompensationWorkflowsOK describes a response with status code 200, with default header values.

?????????
*/
type GetCompensationWorkflowsOK struct {
	Payload *GetCompensationWorkflowsOKBody
}

func (o *GetCompensationWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /compensation/workflows][%d] getCompensationWorkflowsOK  %+v", 200, o.Payload)
}
func (o *GetCompensationWorkflowsOK) GetPayload() *GetCompensationWorkflowsOKBody {
	return o.Payload
}

func (o *GetCompensationWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCompensationWorkflowsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompensationWorkflowsUnauthorized creates a GetCompensationWorkflowsUnauthorized with default headers values
func NewGetCompensationWorkflowsUnauthorized() *GetCompensationWorkflowsUnauthorized {
	return &GetCompensationWorkflowsUnauthorized{}
}

/* GetCompensationWorkflowsUnauthorized describes a response with status code 401, with default header values.

API?????????????????????????????????????????????
*/
type GetCompensationWorkflowsUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *GetCompensationWorkflowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /compensation/workflows][%d] getCompensationWorkflowsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCompensationWorkflowsUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetCompensationWorkflowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompensationWorkflowsForbidden creates a GetCompensationWorkflowsForbidden with default headers values
func NewGetCompensationWorkflowsForbidden() *GetCompensationWorkflowsForbidden {
	return &GetCompensationWorkflowsForbidden{}
}

/* GetCompensationWorkflowsForbidden describes a response with status code 403, with default header values.

API????????????????????????????????????????????????
*/
type GetCompensationWorkflowsForbidden struct {
	Payload *model.ErrorModel
}

func (o *GetCompensationWorkflowsForbidden) Error() string {
	return fmt.Sprintf("[GET /compensation/workflows][%d] getCompensationWorkflowsForbidden  %+v", 403, o.Payload)
}
func (o *GetCompensationWorkflowsForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetCompensationWorkflowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompensationWorkflowsUnprocessableEntity creates a GetCompensationWorkflowsUnprocessableEntity with default headers values
func NewGetCompensationWorkflowsUnprocessableEntity() *GetCompensationWorkflowsUnprocessableEntity {
	return &GetCompensationWorkflowsUnprocessableEntity{}
}

/* GetCompensationWorkflowsUnprocessableEntity describes a response with status code 422, with default header values.

????????????????????????????????????????????????
*/
type GetCompensationWorkflowsUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *GetCompensationWorkflowsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /compensation/workflows][%d] getCompensationWorkflowsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetCompensationWorkflowsUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetCompensationWorkflowsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompensationWorkflowsInternalServerError creates a GetCompensationWorkflowsInternalServerError with default headers values
func NewGetCompensationWorkflowsInternalServerError() *GetCompensationWorkflowsInternalServerError {
	return &GetCompensationWorkflowsInternalServerError{}
}

/* GetCompensationWorkflowsInternalServerError describes a response with status code 500, with default header values.

API?????????????????????????????????????????????
*/
type GetCompensationWorkflowsInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *GetCompensationWorkflowsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /compensation/workflows][%d] getCompensationWorkflowsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCompensationWorkflowsInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *GetCompensationWorkflowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCompensationWorkflowsOKBody get compensation workflows o k body
swagger:model GetCompensationWorkflowsOKBody
*/
type GetCompensationWorkflowsOKBody struct {

	// response
	Response *GetCompensationWorkflowsOKBodyResponse `json:"response,omitempty"`

	// ???????????????
	ResultCode string `json:"result_code,omitempty"`

	// ?????????????????????
	Status bool `json:"status,omitempty"`
}

// Validate validates this get compensation workflows o k body
func (o *GetCompensationWorkflowsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCompensationWorkflowsOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCompensationWorkflowsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get compensation workflows o k body based on the context it is used
func (o *GetCompensationWorkflowsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCompensationWorkflowsOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCompensationWorkflowsOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCompensationWorkflowsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCompensationWorkflowsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCompensationWorkflowsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCompensationWorkflowsOKBodyResponse API???????????????
swagger:model GetCompensationWorkflowsOKBodyResponse
*/
type GetCompensationWorkflowsOKBodyResponse struct {

	// workflows
	Workflows []*model.CompensationWorkflowModel `json:"workflows"`
}

// Validate validates this get compensation workflows o k body response
func (o *GetCompensationWorkflowsOKBodyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCompensationWorkflowsOKBodyResponse) validateWorkflows(formats strfmt.Registry) error {
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
					return ve.ValidateName("getCompensationWorkflowsOK" + "." + "response" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get compensation workflows o k body response based on the context it is used
func (o *GetCompensationWorkflowsOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateWorkflows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCompensationWorkflowsOKBodyResponse) contextValidateWorkflows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Workflows); i++ {

		if o.Workflows[i] != nil {
			if err := o.Workflows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCompensationWorkflowsOK" + "." + "response" + "." + "workflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCompensationWorkflowsOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCompensationWorkflowsOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res GetCompensationWorkflowsOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

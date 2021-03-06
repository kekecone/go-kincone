// Code generated by go-swagger; DO NOT EDIT.

package hourly_paid_leave_workflows

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

// DeleteHourlyPaidLeaveWorkflowWorkflowIDReader is a Reader for the DeleteHourlyPaidLeaveWorkflowWorkflowID structure.
type DeleteHourlyPaidLeaveWorkflowWorkflowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHourlyPaidLeaveWorkflowWorkflowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteHourlyPaidLeaveWorkflowWorkflowIDOK creates a DeleteHourlyPaidLeaveWorkflowWorkflowIDOK with default headers values
func NewDeleteHourlyPaidLeaveWorkflowWorkflowIDOK() *DeleteHourlyPaidLeaveWorkflowWorkflowIDOK {
	return &DeleteHourlyPaidLeaveWorkflowWorkflowIDOK{}
}

/* DeleteHourlyPaidLeaveWorkflowWorkflowIDOK describes a response with status code 200, with default header values.

?????????
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDOK struct {
	Payload *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOK) Error() string {
	return fmt.Sprintf("[DELETE /hourly_paid_leave/workflow/{workflowID}][%d] deleteHourlyPaidLeaveWorkflowWorkflowIdOK  %+v", 200, o.Payload)
}
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOK) GetPayload() *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody {
	return o.Payload
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized creates a DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized with default headers values
func NewDeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized() *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized {
	return &DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized{}
}

/* DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized describes a response with status code 401, with default header values.

API?????????????????????????????????????????????
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /hourly_paid_leave/workflow/{workflowID}][%d] deleteHourlyPaidLeaveWorkflowWorkflowIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden creates a DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden with default headers values
func NewDeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden() *DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden {
	return &DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden{}
}

/* DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden describes a response with status code 403, with default header values.

API????????????????????????????????????????????????
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden struct {
	Payload *model.ErrorModel
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /hourly_paid_leave/workflow/{workflowID}][%d] deleteHourlyPaidLeaveWorkflowWorkflowIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity creates a DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity with default headers values
func NewDeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity() *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity {
	return &DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity{}
}

/* DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity describes a response with status code 422, with default header values.

????????????????????????????????????????????????
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /hourly_paid_leave/workflow/{workflowID}][%d] deleteHourlyPaidLeaveWorkflowWorkflowIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError creates a DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError with default headers values
func NewDeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError() *DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError {
	return &DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError{}
}

/* DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError describes a response with status code 500, with default header values.

API?????????????????????????????????????????????
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /hourly_paid_leave/workflow/{workflowID}][%d] deleteHourlyPaidLeaveWorkflowWorkflowIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody delete hourly paid leave workflow workflow ID o k body
swagger:model DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody struct {

	// response
	Response *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse `json:"response,omitempty"`

	// ???????????????
	ResultCode string `json:"result_code,omitempty"`

	// ?????????????????????
	Status bool `json:"status,omitempty"`
}

// Validate validates this delete hourly paid leave workflow workflow ID o k body
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteHourlyPaidLeaveWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delete hourly paid leave workflow workflow ID o k body based on the context it is used
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteHourlyPaidLeaveWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse API???????????????
swagger:model DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse
*/
type DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse struct {

	// ????????????
	Result bool `json:"result,omitempty"`
}

// Validate validates this delete hourly paid leave workflow workflow ID o k body response
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete hourly paid leave workflow workflow ID o k body response based on context it is used
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res DeleteHourlyPaidLeaveWorkflowWorkflowIDOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

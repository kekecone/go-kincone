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

	"github.com/kekecone/go-kincone/kinkone/model"
)

// PutAttendanceWorkflowWorkflowIDReader is a Reader for the PutAttendanceWorkflowWorkflowID structure.
type PutAttendanceWorkflowWorkflowIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAttendanceWorkflowWorkflowIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAttendanceWorkflowWorkflowIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutAttendanceWorkflowWorkflowIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAttendanceWorkflowWorkflowIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAttendanceWorkflowWorkflowIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAttendanceWorkflowWorkflowIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAttendanceWorkflowWorkflowIDOK creates a PutAttendanceWorkflowWorkflowIDOK with default headers values
func NewPutAttendanceWorkflowWorkflowIDOK() *PutAttendanceWorkflowWorkflowIDOK {
	return &PutAttendanceWorkflowWorkflowIDOK{}
}

/* PutAttendanceWorkflowWorkflowIDOK describes a response with status code 200, with default header values.

?????????
*/
type PutAttendanceWorkflowWorkflowIDOK struct {
	Payload *PutAttendanceWorkflowWorkflowIDOKBody
}

func (o *PutAttendanceWorkflowWorkflowIDOK) Error() string {
	return fmt.Sprintf("[PUT /attendance/workflow/{workflowID}][%d] putAttendanceWorkflowWorkflowIdOK  %+v", 200, o.Payload)
}
func (o *PutAttendanceWorkflowWorkflowIDOK) GetPayload() *PutAttendanceWorkflowWorkflowIDOKBody {
	return o.Payload
}

func (o *PutAttendanceWorkflowWorkflowIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutAttendanceWorkflowWorkflowIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAttendanceWorkflowWorkflowIDUnauthorized creates a PutAttendanceWorkflowWorkflowIDUnauthorized with default headers values
func NewPutAttendanceWorkflowWorkflowIDUnauthorized() *PutAttendanceWorkflowWorkflowIDUnauthorized {
	return &PutAttendanceWorkflowWorkflowIDUnauthorized{}
}

/* PutAttendanceWorkflowWorkflowIDUnauthorized describes a response with status code 401, with default header values.

API?????????????????????????????????????????????
*/
type PutAttendanceWorkflowWorkflowIDUnauthorized struct {
	Payload *model.ErrorModel
}

func (o *PutAttendanceWorkflowWorkflowIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /attendance/workflow/{workflowID}][%d] putAttendanceWorkflowWorkflowIdUnauthorized  %+v", 401, o.Payload)
}
func (o *PutAttendanceWorkflowWorkflowIDUnauthorized) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PutAttendanceWorkflowWorkflowIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAttendanceWorkflowWorkflowIDForbidden creates a PutAttendanceWorkflowWorkflowIDForbidden with default headers values
func NewPutAttendanceWorkflowWorkflowIDForbidden() *PutAttendanceWorkflowWorkflowIDForbidden {
	return &PutAttendanceWorkflowWorkflowIDForbidden{}
}

/* PutAttendanceWorkflowWorkflowIDForbidden describes a response with status code 403, with default header values.

API????????????????????????????????????????????????
*/
type PutAttendanceWorkflowWorkflowIDForbidden struct {
	Payload *model.ErrorModel
}

func (o *PutAttendanceWorkflowWorkflowIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /attendance/workflow/{workflowID}][%d] putAttendanceWorkflowWorkflowIdForbidden  %+v", 403, o.Payload)
}
func (o *PutAttendanceWorkflowWorkflowIDForbidden) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PutAttendanceWorkflowWorkflowIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAttendanceWorkflowWorkflowIDUnprocessableEntity creates a PutAttendanceWorkflowWorkflowIDUnprocessableEntity with default headers values
func NewPutAttendanceWorkflowWorkflowIDUnprocessableEntity() *PutAttendanceWorkflowWorkflowIDUnprocessableEntity {
	return &PutAttendanceWorkflowWorkflowIDUnprocessableEntity{}
}

/* PutAttendanceWorkflowWorkflowIDUnprocessableEntity describes a response with status code 422, with default header values.

????????????????????????????????????????????????
*/
type PutAttendanceWorkflowWorkflowIDUnprocessableEntity struct {
	Payload *model.ErrorModel
}

func (o *PutAttendanceWorkflowWorkflowIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /attendance/workflow/{workflowID}][%d] putAttendanceWorkflowWorkflowIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PutAttendanceWorkflowWorkflowIDUnprocessableEntity) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PutAttendanceWorkflowWorkflowIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAttendanceWorkflowWorkflowIDInternalServerError creates a PutAttendanceWorkflowWorkflowIDInternalServerError with default headers values
func NewPutAttendanceWorkflowWorkflowIDInternalServerError() *PutAttendanceWorkflowWorkflowIDInternalServerError {
	return &PutAttendanceWorkflowWorkflowIDInternalServerError{}
}

/* PutAttendanceWorkflowWorkflowIDInternalServerError describes a response with status code 500, with default header values.

API?????????????????????????????????????????????
*/
type PutAttendanceWorkflowWorkflowIDInternalServerError struct {
	Payload *model.ErrorModel
}

func (o *PutAttendanceWorkflowWorkflowIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /attendance/workflow/{workflowID}][%d] putAttendanceWorkflowWorkflowIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PutAttendanceWorkflowWorkflowIDInternalServerError) GetPayload() *model.ErrorModel {
	return o.Payload
}

func (o *PutAttendanceWorkflowWorkflowIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutAttendanceWorkflowWorkflowIDOKBody put attendance workflow workflow ID o k body
swagger:model PutAttendanceWorkflowWorkflowIDOKBody
*/
type PutAttendanceWorkflowWorkflowIDOKBody struct {

	// response
	Response *PutAttendanceWorkflowWorkflowIDOKBodyResponse `json:"response,omitempty"`

	// ???????????????
	ResultCode string `json:"result_code,omitempty"`

	// ?????????????????????
	Status bool `json:"status,omitempty"`
}

// Validate validates this put attendance workflow workflow ID o k body
func (o *PutAttendanceWorkflowWorkflowIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutAttendanceWorkflowWorkflowIDOKBody) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(o.Response) { // not required
		return nil
	}

	if o.Response != nil {
		if err := o.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putAttendanceWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put attendance workflow workflow ID o k body based on the context it is used
func (o *PutAttendanceWorkflowWorkflowIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutAttendanceWorkflowWorkflowIDOKBody) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if o.Response != nil {
		if err := o.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putAttendanceWorkflowWorkflowIdOK" + "." + "response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutAttendanceWorkflowWorkflowIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAttendanceWorkflowWorkflowIDOKBody) UnmarshalBinary(b []byte) error {
	var res PutAttendanceWorkflowWorkflowIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutAttendanceWorkflowWorkflowIDOKBodyResponse API???????????????
swagger:model PutAttendanceWorkflowWorkflowIDOKBodyResponse
*/
type PutAttendanceWorkflowWorkflowIDOKBodyResponse struct {

	// ????????????
	Result bool `json:"result,omitempty"`
}

// Validate validates this put attendance workflow workflow ID o k body response
func (o *PutAttendanceWorkflowWorkflowIDOKBodyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put attendance workflow workflow ID o k body response based on context it is used
func (o *PutAttendanceWorkflowWorkflowIDOKBodyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutAttendanceWorkflowWorkflowIDOKBodyResponse) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutAttendanceWorkflowWorkflowIDOKBodyResponse) UnmarshalBinary(b []byte) error {
	var res PutAttendanceWorkflowWorkflowIDOKBodyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package kinkone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttendanceWorkflowModel 勤怠申請
//
// swagger:model attendanceWorkflowModel
type AttendanceWorkflowModel struct {

	// 欠勤回数
	AbsentDays int64 `json:"absent_days,omitempty"`

	// 実労働時間合計
	ActualWorkHours string `json:"actual_work_hours,omitempty"`

	// 月間所定労働時間
	AgreedRegularWorkHours string `json:"agreed_regular_work_hours,omitempty"`

	// 申請日
	AppliedDate string `json:"applied_date,omitempty"`

	// 出勤回数
	AttendDays int64 `json:"attend_days,omitempty"`

	// 差戻しコメント
	BackComment string `json:"back_comment,omitempty"`

	// 休憩時間合計
	BreakHours string `json:"break_hours,omitempty"`

	// 遅刻回数
	ComeLateDays int64 `json:"come_late_days,omitempty"`

	// 遅刻時間
	ComeLateHours string `json:"come_late_hours,omitempty"`

	// 申請コメント
	Comment string `json:"comment,omitempty"`

	// 承認日
	ConfirmedDate string `json:"confirmed_date,omitempty"`

	// 勤怠明細
	Details []*AttendanceModel `json:"details"`

	// メールアドレス
	Email string `json:"email,omitempty"`

	// 期間終了日
	EndDate string `json:"end_date,omitempty"`

	// 姓
	FamilyName string `json:"family_name,omitempty"`

	// 名
	FirstName string `json:"first_name,omitempty"`

	// 固定残業時間
	FixedOvertimeWorkHours string `json:"fixed_overtime_work_hours,omitempty"`

	// 適用労働時間合計
	FixedWorkHours string `json:"fixed_work_hours,omitempty"`

	// 時間休取得時間合計
	HourlyPaidLeave string `json:"hourly_paid_leave,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// 不完全なデータ有り
	Invalid bool `json:"invalid,omitempty"`

	// 深夜労働時間合計
	LatenightWorkHours string `json:"latenight_work_hours,omitempty"`

	// 早退回数
	LeaveEarlyDays int64 `json:"leave_early_days,omitempty"`

	// 早退時間
	LeaveEarlyHours string `json:"leave_early_hours,omitempty"`

	// 期間終了時点の時間休残時間
	LeftHourlyPaidHolidayAtEndDate string `json:"left_hourly_paid_holiday_at_end_date,omitempty"`

	// 期間終了時点の有給休暇残日数
	LeftPaidHolidayAtEndDate string `json:"left_paid_holiday_at_end_date,omitempty"`

	// 法定休日出勤回数
	LegalHolidayWorkDays int64 `json:"legal_holiday_work_days,omitempty"`

	// 法定休日実労働時間合計
	LegalHolidayWorkHours string `json:"legal_holiday_work_hours,omitempty"`

	// 月間法定労働時間
	MonthlyStatutoryWorkHours string `json:"monthly_statutory_work_hours,omitempty"`

	// 法定外残業時間（60時間超）
	NonStatutoryOvertimeHoursOver60h string `json:"non_statutory_overtime_hours_over_60h,omitempty"`

	// 法定外残業時間（合計）
	NonStatutoryOvertimeHoursTotal string `json:"non_statutory_overtime_hours_total,omitempty"`

	// 法定外残業時間（合計）（固定残業時間適用）
	NonStatutoryOvertimeHoursTotalAppliedFixed string `json:"non_statutory_overtime_hours_total_applied_fixed,omitempty"`

	// 法定外残業時間（60時間まで）
	NonStatutoryOvertimeHoursUnder60h string `json:"non_statutory_overtime_hours_under_60h,omitempty"`

	// 時間外労働時間合計
	OvertimeWorkHours string `json:"overtime_work_hours,omitempty"`

	// 残業時間（固定残業時間適用）
	OvertimeWorkHoursAppliedFixed string `json:"overtime_work_hours_applied_fixed,omitempty"`

	// 有給取得日数合計
	PaidHoliday string `json:"paid_holiday,omitempty"`

	// 所定労働時間合計
	RegularWorkHours string `json:"regular_work_hours,omitempty"`

	// 期間開始日
	StartDate string `json:"start_date,omitempty"`

	// 申請ステータス（1: 差戻し, 2: 申請済み, 3: 承認）
	Status int64 `json:"status,omitempty"`

	// 法定内残業時間
	StatutoryOvertime string `json:"statutory_overtime,omitempty"`

	// 法定内残業時間（固定残業時間適用）
	StatutoryOvertimeAppliedFixed string `json:"statutory_overtime_applied_fixed,omitempty"`

	// 労働時間合計
	TotalWorkHours string `json:"total_work_hours,omitempty"`

	// 従業員番号
	UserNumber string `json:"user_number,omitempty"`
}

// Validate validates this attendance workflow model
func (m *AttendanceWorkflowModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttendanceWorkflowModel) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	for i := 0; i < len(m.Details); i++ {
		if swag.IsZero(m.Details[i]) { // not required
			continue
		}

		if m.Details[i] != nil {
			if err := m.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this attendance workflow model based on the context it is used
func (m *AttendanceWorkflowModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttendanceWorkflowModel) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Details); i++ {

		if m.Details[i] != nil {
			if err := m.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttendanceWorkflowModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttendanceWorkflowModel) UnmarshalBinary(b []byte) error {
	var res AttendanceWorkflowModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

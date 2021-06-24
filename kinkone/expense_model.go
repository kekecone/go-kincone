// Code generated by go-swagger; DO NOT EDIT.

package kinkone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExpenseModel 交通費明細
//
// swagger:model expenseModel
type ExpenseModel struct {

	// 勘定科目名
	Account string `json:"account,omitempty"`

	// 勘定科目コード
	AccountCode string `json:"account_code,omitempty"`

	// 利用日
	Date string `json:"date,omitempty"`

	// 交通費
	Expense int64 `json:"expense,omitempty"`

	// 利用区間（入場駅）
	In string `json:"in,omitempty"`

	// 訪問先/備考
	Note string `json:"note,omitempty"`

	// 利用区間（出場駅）
	Out string `json:"out,omitempty"`

	// 用途
	Purpose string `json:"purpose,omitempty"`

	// 補助科目名
	SubAccount string `json:"sub_account,omitempty"`

	// 補助科目コード
	SubAccountCode string `json:"sub_account_code,omitempty"`

	// 交通手段（1: 電車, 2: バス, 3: タクシー, 4: 新幹線/特急, 5: 飛行機, 6: 車, 7: 船/フェリー, 8: その他, 0: 物販）
	Type int64 `json:"type,omitempty"`
}

// Validate validates this expense model
func (m *ExpenseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this expense model based on context it is used
func (m *ExpenseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExpenseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpenseModel) UnmarshalBinary(b []byte) error {
	var res ExpenseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

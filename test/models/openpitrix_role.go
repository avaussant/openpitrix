// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRole openpitrix role
// swagger:model openpitrixRole
type OpenpitrixRole struct {

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// owner path
	OwnerPath string `json:"owner_path,omitempty"`

	// portal
	Portal string `json:"portal,omitempty"`

	// role id
	RoleID string `json:"role_id,omitempty"`

	// role name
	RoleName string `json:"role_name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// update time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// user id
	UserID []string `json:"user_id"`
}

// Validate validates this openpitrix role
func (m *OpenpitrixRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixRole) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRole) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
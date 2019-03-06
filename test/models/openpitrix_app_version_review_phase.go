// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAppVersionReviewPhase openpitrix app version review phase
// swagger:model openpitrixAppVersionReviewPhase
type OpenpitrixAppVersionReviewPhase struct {

	// message
	Message string `json:"message,omitempty"`

	// operator
	Operator string `json:"operator,omitempty"`

	// review access
	ReviewAccess string `json:"review_access,omitempty"`

	// review time
	ReviewTime strfmt.DateTime `json:"review_time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`
}

// Validate validates this openpitrix app version review phase
func (m *OpenpitrixAppVersionReviewPhase) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAppVersionReviewPhase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAppVersionReviewPhase) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAppVersionReviewPhase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRepo openpitrix repo
// swagger:model openpitrixRepo
type OpenpitrixRepo struct {

	// app default status
	AppDefaultStatus string `json:"app_default_status,omitempty"`

	// category set
	CategorySet OpenpitrixRepoCategorySet `json:"category_set"`

	// controller
	Controller int32 `json:"controller,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// credential
	Credential string `json:"credential,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels OpenpitrixRepoLabels `json:"labels"`

	// name
	Name string `json:"name,omitempty"`

	// owner path
	OwnerPath string `json:"owner_path,omitempty"`

	// providers
	Providers []string `json:"providers"`

	// repo id
	RepoID string `json:"repo_id,omitempty"`

	// selectors
	Selectors OpenpitrixRepoSelectors `json:"selectors"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// visibility
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this openpitrix repo
func (m *OpenpitrixRepo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviders(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixRepo) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRepo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRepo) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRepo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

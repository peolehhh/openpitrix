// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyAppVersionResponse openpitrix modify app version response
// swagger:model openpitrixModifyAppVersionResponse
type OpenpitrixModifyAppVersionResponse struct {

	// required, version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix modify app version response
func (m *OpenpitrixModifyAppVersionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyAppVersionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyAppVersionResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyAppVersionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

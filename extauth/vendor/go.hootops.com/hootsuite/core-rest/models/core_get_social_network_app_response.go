// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CoreGetSocialNetworkAppResponse core get social network app response
// swagger:model coreGetSocialNetworkAppResponse
type CoreGetSocialNetworkAppResponse struct {

	// auth1
	Auth1 string `json:"auth1,omitempty"`

	// auth2
	Auth2 string `json:"auth2,omitempty"`

	// auth3
	Auth3 string `json:"auth3,omitempty"`

	// created date
	CreatedDate string `json:"createdDate,omitempty"`

	// created user
	CreatedUser uint64 `json:"createdUser,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// modified date
	ModifiedDate string `json:"modifiedDate,omitempty"`

	// modified user
	ModifiedUser uint64 `json:"modifiedUser,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID uint64 `json:"organizationId,omitempty"`
}

// Validate validates this core get social network app response
func (m *CoreGetSocialNetworkAppResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreGetSocialNetworkAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreGetSocialNetworkAppResponse) UnmarshalBinary(b []byte) error {
	var res CoreGetSocialNetworkAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
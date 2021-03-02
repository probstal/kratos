// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationFlow registration flow
//
// swagger:model registrationFlow
type RegistrationFlow struct {

	// active
	Active CredentialsType `json:"active,omitempty"`

	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in,
	// a new flow has to be initiated.
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expires_at"`

	// id
	// Required: true
	// Format: uuid4
	ID *UUID `json:"id"`

	// IssuedAt is the time (UTC) when the flow occurred.
	// Required: true
	// Format: date-time
	IssuedAt *strfmt.DateTime `json:"issued_at"`

	// messages
	Messages Messages `json:"messages,omitempty"`

	// Methods contains context for all enabled registration methods. If a registration flow has been
	// processed, but for example the password is incorrect, this will contain error messages.
	// Required: true
	Methods map[string]RegistrationFlowMethod `json:"methods"`

	// RequestURL is the initial URL that was requested from ORY Kratos. It can be used
	// to forward information contained in the URL's path or query for example.
	// Required: true
	RequestURL *string `json:"request_url"`

	// type
	Type Type `json:"type,omitempty"`
}

// Validate validates this registration flow
func (m *RegistrationFlow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationFlow) validateActive(formats strfmt.Registry) error {
	if swag.IsZero(m.Active) { // not required
		return nil
	}

	if err := m.Active.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active")
		}
		return err
	}

	return nil
}

func (m *RegistrationFlow) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expires_at", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationFlow) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *RegistrationFlow) validateIssuedAt(formats strfmt.Registry) error {

	if err := validate.Required("issued_at", "body", m.IssuedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("issued_at", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationFlow) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	if err := m.Messages.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("messages")
		}
		return err
	}

	return nil
}

func (m *RegistrationFlow) validateMethods(formats strfmt.Registry) error {

	if err := validate.Required("methods", "body", m.Methods); err != nil {
		return err
	}

	for k := range m.Methods {

		if err := validate.Required("methods"+"."+k, "body", m.Methods[k]); err != nil {
			return err
		}
		if val, ok := m.Methods[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RegistrationFlow) validateRequestURL(formats strfmt.Registry) error {

	if err := validate.Required("request_url", "body", m.RequestURL); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationFlow) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this registration flow based on the context it is used
func (m *RegistrationFlow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationFlow) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Active.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active")
		}
		return err
	}

	return nil
}

func (m *RegistrationFlow) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {
		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *RegistrationFlow) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Messages.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("messages")
		}
		return err
	}

	return nil
}

func (m *RegistrationFlow) contextValidateMethods(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("methods", "body", m.Methods); err != nil {
		return err
	}

	for k := range m.Methods {

		if val, ok := m.Methods[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RegistrationFlow) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationFlow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationFlow) UnmarshalBinary(b []byte) error {
	var res RegistrationFlow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
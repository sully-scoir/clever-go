// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DistrictsCreated districts created
//
// swagger:model districts.created
type DistrictsCreated struct {
	createdField strfmt.DateTime

	idField string

	// data
	Data *DistrictObject `json:"data,omitempty"`
}

// Created gets the created of this subtype
func (m *DistrictsCreated) Created() strfmt.DateTime {
	return m.createdField
}

// SetCreated sets the created of this subtype
func (m *DistrictsCreated) SetCreated(val strfmt.DateTime) {
	m.createdField = val
}

// ID gets the id of this subtype
func (m *DistrictsCreated) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *DistrictsCreated) SetID(val string) {
	m.idField = val
}

// Type gets the type of this subtype
func (m *DistrictsCreated) Type() string {
	return "districts.created"
}

// SetType sets the type of this subtype
func (m *DistrictsCreated) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DistrictsCreated) UnmarshalJSON(raw []byte) error {
	var data struct {

		// data
		Data *DistrictObject `json:"data,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Created strfmt.DateTime `json:"created,omitempty"`

		ID string `json:"id,omitempty"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DistrictsCreated

	result.createdField = base.Created

	result.idField = base.ID

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Data = data.Data

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DistrictsCreated) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// data
		Data *DistrictObject `json:"data,omitempty"`
	}{

		Data: m.Data,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Created strfmt.DateTime `json:"created,omitempty"`

		ID string `json:"id,omitempty"`

		Type string `json:"type"`
	}{

		Created: m.Created(),

		ID: m.ID(),

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this districts created
func (m *DistrictsCreated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrictsCreated) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created()) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "datetime", m.Created().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DistrictsCreated) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this districts created based on the context it is used
func (m *DistrictsCreated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrictsCreated) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistrictsCreated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrictsCreated) UnmarshalBinary(b []byte) error {
	var res DistrictsCreated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

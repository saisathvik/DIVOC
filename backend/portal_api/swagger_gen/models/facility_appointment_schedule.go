// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FacilityAppointmentSchedule facility appointment schedule
//
// swagger:model FacilityAppointmentSchedule
type FacilityAppointmentSchedule struct {

	// days
	Days []*FacilityAppointmentScheduleDaysItems0 `json:"days"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`
}

// Validate validates this facility appointment schedule
func (m *FacilityAppointmentSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityAppointmentSchedule) validateDays(formats strfmt.Registry) error {

	if swag.IsZero(m.Days) { // not required
		return nil
	}

	for i := 0; i < len(m.Days); i++ {
		if swag.IsZero(m.Days[i]) { // not required
			continue
		}

		if m.Days[i] != nil {
			if err := m.Days[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("days" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacilityAppointmentSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityAppointmentSchedule) UnmarshalBinary(b []byte) error {
	var res FacilityAppointmentSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FacilityAppointmentScheduleDaysItems0 facility appointment schedule days items0
//
// swagger:model FacilityAppointmentScheduleDaysItems0
type FacilityAppointmentScheduleDaysItems0 struct {

	// day (mon, tue, wed, thr, fri, sat, sun)
	Day string `json:"day,omitempty"`

	// Maximum appointment per day
	// Minimum: 0
	MaxAppointments *int64 `json:"maxAppointments,omitempty"`
}

// Validate validates this facility appointment schedule days items0
func (m *FacilityAppointmentScheduleDaysItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxAppointments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityAppointmentScheduleDaysItems0) validateMaxAppointments(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxAppointments) { // not required
		return nil
	}

	if err := validate.MinimumInt("maxAppointments", "body", int64(*m.MaxAppointments), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacilityAppointmentScheduleDaysItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityAppointmentScheduleDaysItems0) UnmarshalBinary(b []byte) error {
	var res FacilityAppointmentScheduleDaysItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

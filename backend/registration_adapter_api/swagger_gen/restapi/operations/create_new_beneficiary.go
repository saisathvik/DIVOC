// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNewBeneficiaryHandlerFunc turns a function with the right signature into a create new beneficiary handler
type CreateNewBeneficiaryHandlerFunc func(CreateNewBeneficiaryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateNewBeneficiaryHandlerFunc) Handle(params CreateNewBeneficiaryParams) middleware.Responder {
	return fn(params)
}

// CreateNewBeneficiaryHandler interface for that can handle valid create new beneficiary params
type CreateNewBeneficiaryHandler interface {
	Handle(CreateNewBeneficiaryParams) middleware.Responder
}

// NewCreateNewBeneficiary creates a new http.Handler for the create new beneficiary operation
func NewCreateNewBeneficiary(ctx *middleware.Context, handler CreateNewBeneficiaryHandler) *CreateNewBeneficiary {
	return &CreateNewBeneficiary{Context: ctx, Handler: handler}
}

/*CreateNewBeneficiary swagger:route POST /registration/beneficiaries/new createNewBeneficiary

Create a new benficiary

*/
type CreateNewBeneficiary struct {
	Context *middleware.Context
	Handler CreateNewBeneficiaryHandler
}

func (o *CreateNewBeneficiary) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateNewBeneficiaryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateNewBeneficiaryBody create new beneficiary body
//
// swagger:model CreateNewBeneficiaryBody
type CreateNewBeneficiaryBody struct {

	// birth year
	// Required: true
	BirthYear *string `json:"birth_year"`

	// comorbidities
	// Required: true
	Comorbidities []string `json:"comorbidities"`

	// consent
	// Required: true
	Consent *string `json:"consent"`

	// district id
	// Required: true
	DistrictID *string `json:"district_id"`

	// gender
	// Required: true
	Gender *string `json:"gender"`

	// name
	// Required: true
	Name *string `json:"name"`

	// photo id number
	// Required: true
	PhotoIDNumber *string `json:"photo_id_number"`

	// photo id type
	// Required: true
	PhotoIDType *string `json:"photo_id_type"`

	// pincode
	// Required: true
	Pincode *string `json:"pincode"`

	// state id
	// Required: true
	StateID *string `json:"state_id"`
}

// Validate validates this create new beneficiary body
func (o *CreateNewBeneficiaryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBirthYear(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateComorbidities(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConsent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDistrictID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePhotoIDNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePhotoIDType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePincode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNewBeneficiaryBody) validateBirthYear(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"birth_year", "body", o.BirthYear); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validateComorbidities(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"comorbidities", "body", o.Comorbidities); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validateConsent(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"consent", "body", o.Consent); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validateDistrictID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"district_id", "body", o.DistrictID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"gender", "body", o.Gender); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validatePhotoIDNumber(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"photo_id_number", "body", o.PhotoIDNumber); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validatePhotoIDType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"photo_id_type", "body", o.PhotoIDType); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validatePincode(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"pincode", "body", o.Pincode); err != nil {
		return err
	}

	return nil
}

func (o *CreateNewBeneficiaryBody) validateStateID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"state_id", "body", o.StateID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNewBeneficiaryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNewBeneficiaryBody) UnmarshalBinary(b []byte) error {
	var res CreateNewBeneficiaryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateNewBeneficiaryOKBody create new beneficiary o k body
//
// swagger:model CreateNewBeneficiaryOKBody
type CreateNewBeneficiaryOKBody struct {

	// beneficiary reference id
	BeneficiaryReferenceID string `json:"beneficiary_reference_id,omitempty"`
}

// Validate validates this create new beneficiary o k body
func (o *CreateNewBeneficiaryOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNewBeneficiaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNewBeneficiaryOKBody) UnmarshalBinary(b []byte) error {
	var res CreateNewBeneficiaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

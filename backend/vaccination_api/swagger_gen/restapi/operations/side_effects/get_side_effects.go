// Code generated by go-swagger; DO NOT EDIT.

package side_effects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/api/swagger_gen/models"
)

// GetSideEffectsHandlerFunc turns a function with the right signature into a get side effects handler
type GetSideEffectsHandlerFunc func(GetSideEffectsParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSideEffectsHandlerFunc) Handle(params GetSideEffectsParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetSideEffectsHandler interface for that can handle valid get side effects params
type GetSideEffectsHandler interface {
	Handle(GetSideEffectsParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetSideEffects creates a new http.Handler for the get side effects operation
func NewGetSideEffects(ctx *middleware.Context, handler GetSideEffectsHandler) *GetSideEffects {
	return &GetSideEffects{Context: ctx, Handler: handler}
}

/*GetSideEffects swagger:route GET /sideEffects sideEffects getSideEffects

Get Side Effects

*/
type GetSideEffects struct {
	Context *middleware.Context
	Handler GetSideEffectsHandler
}

func (o *GetSideEffects) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSideEffectsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

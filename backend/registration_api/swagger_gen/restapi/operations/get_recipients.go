// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/registration-api/swagger_gen/models"
)

// GetRecipientsHandlerFunc turns a function with the right signature into a get recipients handler
type GetRecipientsHandlerFunc func(GetRecipientsParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRecipientsHandlerFunc) Handle(params GetRecipientsParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetRecipientsHandler interface for that can handle valid get recipients params
type GetRecipientsHandler interface {
	Handle(GetRecipientsParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetRecipients creates a new http.Handler for the get recipients operation
func NewGetRecipients(ctx *middleware.Context, handler GetRecipientsHandler) *GetRecipients {
	return &GetRecipients{Context: ctx, Handler: handler}
}

/* GetRecipients swagger:route GET /recipients getRecipients

Get all the recipients

*/
type GetRecipients struct {
	Context *middleware.Context
	Handler GetRecipientsHandler
}

func (o *GetRecipients) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRecipientsParams()

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

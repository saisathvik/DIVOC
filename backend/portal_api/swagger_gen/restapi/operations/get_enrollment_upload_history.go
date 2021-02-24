// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// GetEnrollmentUploadHistoryHandlerFunc turns a function with the right signature into a get enrollment upload history handler
type GetEnrollmentUploadHistoryHandlerFunc func(GetEnrollmentUploadHistoryParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEnrollmentUploadHistoryHandlerFunc) Handle(params GetEnrollmentUploadHistoryParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetEnrollmentUploadHistoryHandler interface for that can handle valid get enrollment upload history params
type GetEnrollmentUploadHistoryHandler interface {
	Handle(GetEnrollmentUploadHistoryParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetEnrollmentUploadHistory creates a new http.Handler for the get enrollment upload history operation
func NewGetEnrollmentUploadHistory(ctx *middleware.Context, handler GetEnrollmentUploadHistoryHandler) *GetEnrollmentUploadHistory {
	return &GetEnrollmentUploadHistory{Context: ctx, Handler: handler}
}

/* GetEnrollmentUploadHistory swagger:route GET /enrollments/uploads getEnrollmentUploadHistory

Get Enrollments uploads

*/
type GetEnrollmentUploadHistory struct {
	Context *middleware.Context
	Handler GetEnrollmentUploadHistoryHandler
}

func (o *GetEnrollmentUploadHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEnrollmentUploadHistoryParams()
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

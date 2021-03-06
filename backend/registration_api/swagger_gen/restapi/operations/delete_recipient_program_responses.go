// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteRecipientProgramOKCode is the HTTP code returned for type DeleteRecipientProgramOK
const DeleteRecipientProgramOKCode int = 200

/*DeleteRecipientProgramOK OK

swagger:response deleteRecipientProgramOK
*/
type DeleteRecipientProgramOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteRecipientProgramOK creates DeleteRecipientProgramOK with default headers values
func NewDeleteRecipientProgramOK() *DeleteRecipientProgramOK {

	return &DeleteRecipientProgramOK{}
}

// WithPayload adds the payload to the delete recipient program o k response
func (o *DeleteRecipientProgramOK) WithPayload(payload interface{}) *DeleteRecipientProgramOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete recipient program o k response
func (o *DeleteRecipientProgramOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRecipientProgramOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteRecipientProgramBadRequestCode is the HTTP code returned for type DeleteRecipientProgramBadRequest
const DeleteRecipientProgramBadRequestCode int = 400

/*DeleteRecipientProgramBadRequest Bad Request

swagger:response deleteRecipientProgramBadRequest
*/
type DeleteRecipientProgramBadRequest struct {
}

// NewDeleteRecipientProgramBadRequest creates DeleteRecipientProgramBadRequest with default headers values
func NewDeleteRecipientProgramBadRequest() *DeleteRecipientProgramBadRequest {

	return &DeleteRecipientProgramBadRequest{}
}

// WriteResponse to the client
func (o *DeleteRecipientProgramBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteRecipientProgramUnauthorizedCode is the HTTP code returned for type DeleteRecipientProgramUnauthorized
const DeleteRecipientProgramUnauthorizedCode int = 401

/*DeleteRecipientProgramUnauthorized Unauthorized

swagger:response deleteRecipientProgramUnauthorized
*/
type DeleteRecipientProgramUnauthorized struct {
}

// NewDeleteRecipientProgramUnauthorized creates DeleteRecipientProgramUnauthorized with default headers values
func NewDeleteRecipientProgramUnauthorized() *DeleteRecipientProgramUnauthorized {

	return &DeleteRecipientProgramUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRecipientProgramUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteRecipientProgramInternalServerErrorCode is the HTTP code returned for type DeleteRecipientProgramInternalServerError
const DeleteRecipientProgramInternalServerErrorCode int = 500

/*DeleteRecipientProgramInternalServerError Internal Error

swagger:response deleteRecipientProgramInternalServerError
*/
type DeleteRecipientProgramInternalServerError struct {
}

// NewDeleteRecipientProgramInternalServerError creates DeleteRecipientProgramInternalServerError with default headers values
func NewDeleteRecipientProgramInternalServerError() *DeleteRecipientProgramInternalServerError {

	return &DeleteRecipientProgramInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteRecipientProgramInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

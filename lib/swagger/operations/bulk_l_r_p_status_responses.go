package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*BulkLRPStatusOK OK

swagger:response bulkLRPStatusOK
*/
type BulkLRPStatusOK struct {

	// In: body
	Payload BulkLRPStatusOKBodyBody `json:"body,omitempty"`
}

// NewBulkLRPStatusOK creates BulkLRPStatusOK with default headers values
func NewBulkLRPStatusOK() *BulkLRPStatusOK {
	return &BulkLRPStatusOK{}
}

// WithPayload adds the payload to the bulk l r p status o k response
func (o *BulkLRPStatusOK) WithPayload(payload BulkLRPStatusOKBodyBody) *BulkLRPStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bulk l r p status o k response
func (o *BulkLRPStatusOK) SetPayload(payload BulkLRPStatusOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BulkLRPStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*BulkLRPStatusBadRequest Failed parsing guids

swagger:response bulkLRPStatusBadRequest
*/
type BulkLRPStatusBadRequest struct {
}

// NewBulkLRPStatusBadRequest creates BulkLRPStatusBadRequest with default headers values
func NewBulkLRPStatusBadRequest() *BulkLRPStatusBadRequest {
	return &BulkLRPStatusBadRequest{}
}

// WriteResponse to the client
func (o *BulkLRPStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

/*BulkLRPStatusInternalServerError Failed constructing throttler

swagger:response bulkLRPStatusInternalServerError
*/
type BulkLRPStatusInternalServerError struct {
}

// NewBulkLRPStatusInternalServerError creates BulkLRPStatusInternalServerError with default headers values
func NewBulkLRPStatusInternalServerError() *BulkLRPStatusInternalServerError {
	return &BulkLRPStatusInternalServerError{}
}

// WriteResponse to the client
func (o *BulkLRPStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}
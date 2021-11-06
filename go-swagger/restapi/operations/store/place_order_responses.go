// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger/models"
)

// PlaceOrderOKCode is the HTTP code returned for type PlaceOrderOK
const PlaceOrderOKCode int = 200

/*PlaceOrderOK successful operation

swagger:response placeOrderOK
*/
type PlaceOrderOK struct {

	/*
	  In: Body
	*/
	Payload *models.Order `json:"body,omitempty"`
}

// NewPlaceOrderOK creates PlaceOrderOK with default headers values
func NewPlaceOrderOK() *PlaceOrderOK {

	return &PlaceOrderOK{}
}

// WithPayload adds the payload to the place order o k response
func (o *PlaceOrderOK) WithPayload(payload *models.Order) *PlaceOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the place order o k response
func (o *PlaceOrderOK) SetPayload(payload *models.Order) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PlaceOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PlaceOrderBadRequestCode is the HTTP code returned for type PlaceOrderBadRequest
const PlaceOrderBadRequestCode int = 400

/*PlaceOrderBadRequest Invalid Order

swagger:response placeOrderBadRequest
*/
type PlaceOrderBadRequest struct {
}

// NewPlaceOrderBadRequest creates PlaceOrderBadRequest with default headers values
func NewPlaceOrderBadRequest() *PlaceOrderBadRequest {

	return &PlaceOrderBadRequest{}
}

// WriteResponse to the client
func (o *PlaceOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

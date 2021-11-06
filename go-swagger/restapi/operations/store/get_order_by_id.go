// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetOrderByIDHandlerFunc turns a function with the right signature into a get order by Id handler
type GetOrderByIDHandlerFunc func(GetOrderByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOrderByIDHandlerFunc) Handle(params GetOrderByIDParams) middleware.Responder {
	return fn(params)
}

// GetOrderByIDHandler interface for that can handle valid get order by Id params
type GetOrderByIDHandler interface {
	Handle(GetOrderByIDParams) middleware.Responder
}

// NewGetOrderByID creates a new http.Handler for the get order by Id operation
func NewGetOrderByID(ctx *middleware.Context, handler GetOrderByIDHandler) *GetOrderByID {
	return &GetOrderByID{Context: ctx, Handler: handler}
}

/* GetOrderByID swagger:route GET /store/order/{orderId} store getOrderById

Find purchase order by ID

For valid response try integer IDs with value >= 1 and <= 10.         Other values will generated exceptions

*/
type GetOrderByID struct {
	Context *middleware.Context
	Handler GetOrderByIDHandler
}

func (o *GetOrderByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetOrderByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

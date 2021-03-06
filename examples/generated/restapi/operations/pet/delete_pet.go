// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/examples/generated/models"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// DeletePetHandlerFunc turns a function with the right signature into a delete pet handler
type DeletePetHandlerFunc func(DeletePetParams, *models.User) error

func (fn DeletePetHandlerFunc) Handle(params DeletePetParams, principal *models.User) error {
	return fn(params, principal)
}

// DeletePetHandler interface for that can handle valid delete pet params
type DeletePetHandler interface {
	Handle(DeletePetParams, *models.User) error
}

// NewDeletePet creates a new http.Handler for the delete pet operation
func NewDeletePet(ctx *middleware.Context, handler DeletePetHandler) *DeletePet {
	return &DeletePet{Context: ctx, Handler: handler}
}

/*
Deletes a pet
*/
type DeletePet struct {
	Context *middleware.Context
	Params  DeletePetParams
	Handler DeletePetHandler
}

func (o *DeletePet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	err = o.Handler.Handle(o.Params, principal) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	o.Context.Respond(rw, r, route.Produces, route, nil)

}

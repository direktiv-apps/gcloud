// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams) middleware.Responder {
	return fn(params)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/* Post swagger:route POST / post

Post post API

*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBody post body
//
// swagger:model PostBody
type PostBody struct {

	// Service account name
	// Example: sa@myproject.iam.gserviceaccount.com
	// Required: true
	Account *string `json:"account"`

	// List of commands to run. Use `--format=json` to get JSON results.
	// Example: gcloud compute instances list --format=json
	Commands []string `json:"commands"`

	// If set to true all commands are getting executed and errors ignored.
	// Example: true
	Continue *bool `json:"continue,omitempty"`

	// Base64 encoded JSON access file (IAM). If not provided the function uses `key.json`.
	Key string `json:"key,omitempty"`

	// Specifies the project name.
	// Example: my-project-234
	// Required: true
	Project *string `json:"project"`
}

// Validate validates this post body
func (o *PostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBody) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"account", "body", o.Account); err != nil {
		return err
	}

	return nil
}

func (o *PostBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post body based on context it is used
func (o *PostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBody) UnmarshalBinary(b []byte) error {
	var res PostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBody post o k body
//
// swagger:model PostOKBody
type PostOKBody struct {

	// gcloud
	Gcloud []*PostOKBodyGcloudItems0 `json:"gcloud"`
}

// Validate validates this post o k body
func (o *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGcloud(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) validateGcloud(formats strfmt.Registry) error {
	if swag.IsZero(o.Gcloud) { // not required
		return nil
	}

	for i := 0; i < len(o.Gcloud); i++ {
		if swag.IsZero(o.Gcloud[i]) { // not required
			continue
		}

		if o.Gcloud[i] != nil {
			if err := o.Gcloud[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "gcloud" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "gcloud" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (o *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGcloud(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) contextValidateGcloud(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Gcloud); i++ {

		if o.Gcloud[i] != nil {
			if err := o.Gcloud[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOK" + "." + "gcloud" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOK" + "." + "gcloud" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOKBodyGcloudItems0 post o k body gcloud items0
//
// swagger:model PostOKBodyGcloudItems0
type PostOKBodyGcloudItems0 struct {

	// result
	// Required: true
	Result interface{} `json:"result"`

	// success
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this post o k body gcloud items0
func (o *PostOKBodyGcloudItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBodyGcloudItems0) validateResult(formats strfmt.Registry) error {

	if o.Result == nil {
		return errors.Required("result", "body", nil)
	}

	return nil
}

func (o *PostOKBodyGcloudItems0) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", o.Success); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post o k body gcloud items0 based on context it is used
func (o *PostOKBodyGcloudItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBodyGcloudItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBodyGcloudItems0) UnmarshalBinary(b []byte) error {
	var res PostOKBodyGcloudItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

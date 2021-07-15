// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/openshift/assisted-service/models"
)

// NewV2RegisterInfraEnvParams creates a new V2RegisterInfraEnvParams object
// no default values defined in spec.
func NewV2RegisterInfraEnvParams() V2RegisterInfraEnvParams {

	return V2RegisterInfraEnvParams{}
}

// V2RegisterInfraEnvParams contains all the bound params for the v2 register infra env operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2RegisterInfraEnv
type V2RegisterInfraEnvParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The parameters for the generated ISO.
	  Required: true
	  In: body
	*/
	InfraenvCreateParams *models.InfraenvCreateParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2RegisterInfraEnvParams() beforehand.
func (o *V2RegisterInfraEnvParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.InfraenvCreateParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("infraenvCreateParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("infraenvCreateParams", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.InfraenvCreateParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("infraenvCreateParams", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ListSupportedOpenshiftVersionsOKCode is the HTTP code returned for type V2ListSupportedOpenshiftVersionsOK
const V2ListSupportedOpenshiftVersionsOKCode int = 200

/*V2ListSupportedOpenshiftVersionsOK Success.

swagger:response v2ListSupportedOpenshiftVersionsOK
*/
type V2ListSupportedOpenshiftVersionsOK struct {

	/*
	  In: Body
	*/
	Payload models.OpenshiftVersions `json:"body,omitempty"`
}

// NewV2ListSupportedOpenshiftVersionsOK creates V2ListSupportedOpenshiftVersionsOK with default headers values
func NewV2ListSupportedOpenshiftVersionsOK() *V2ListSupportedOpenshiftVersionsOK {

	return &V2ListSupportedOpenshiftVersionsOK{}
}

// WithPayload adds the payload to the v2 list supported openshift versions o k response
func (o *V2ListSupportedOpenshiftVersionsOK) WithPayload(payload models.OpenshiftVersions) *V2ListSupportedOpenshiftVersionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list supported openshift versions o k response
func (o *V2ListSupportedOpenshiftVersionsOK) SetPayload(payload models.OpenshiftVersions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListSupportedOpenshiftVersionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = models.OpenshiftVersions{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2ListSupportedOpenshiftVersionsServiceUnavailableCode is the HTTP code returned for type V2ListSupportedOpenshiftVersionsServiceUnavailable
const V2ListSupportedOpenshiftVersionsServiceUnavailableCode int = 503

/*V2ListSupportedOpenshiftVersionsServiceUnavailable Unavailable.

swagger:response v2ListSupportedOpenshiftVersionsServiceUnavailable
*/
type V2ListSupportedOpenshiftVersionsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListSupportedOpenshiftVersionsServiceUnavailable creates V2ListSupportedOpenshiftVersionsServiceUnavailable with default headers values
func NewV2ListSupportedOpenshiftVersionsServiceUnavailable() *V2ListSupportedOpenshiftVersionsServiceUnavailable {

	return &V2ListSupportedOpenshiftVersionsServiceUnavailable{}
}

// WithPayload adds the payload to the v2 list supported openshift versions service unavailable response
func (o *V2ListSupportedOpenshiftVersionsServiceUnavailable) WithPayload(payload *models.Error) *V2ListSupportedOpenshiftVersionsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list supported openshift versions service unavailable response
func (o *V2ListSupportedOpenshiftVersionsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListSupportedOpenshiftVersionsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

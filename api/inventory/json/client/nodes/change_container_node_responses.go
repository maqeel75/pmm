// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeContainerNodeReader is a Reader for the ChangeContainerNode structure.
type ChangeContainerNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeContainerNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeContainerNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeContainerNodeOK creates a ChangeContainerNodeOK with default headers values
func NewChangeContainerNodeOK() *ChangeContainerNodeOK {
	return &ChangeContainerNodeOK{}
}

/*ChangeContainerNodeOK handles this case with default header values.

(empty)
*/
type ChangeContainerNodeOK struct {
	Payload *ChangeContainerNodeOKBody
}

func (o *ChangeContainerNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/ChangeContainer][%d] changeContainerNodeOK  %+v", 200, o.Payload)
}

func (o *ChangeContainerNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeContainerNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeContainerNodeBody change container node body
swagger:model ChangeContainerNodeBody
*/
type ChangeContainerNodeBody struct {

	// Container name.
	DockerContainerName string `json:"docker_container_name,omitempty"`

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this change container node body
func (o *ChangeContainerNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeContainerNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeContainerNodeBody) UnmarshalBinary(b []byte) error {
	var res ChangeContainerNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeContainerNodeOKBody change container node o k body
swagger:model ChangeContainerNodeOKBody
*/
type ChangeContainerNodeOKBody struct {

	// container
	Container *ChangeContainerNodeOKBodyContainer `json:"container,omitempty"`
}

// Validate validates this change container node o k body
func (o *ChangeContainerNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeContainerNodeOKBody) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeContainerNodeOK" + "." + "container")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeContainerNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeContainerNodeOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeContainerNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeContainerNodeOKBodyContainer ContainerNode represents a Docker container.
swagger:model ChangeContainerNodeOKBodyContainer
*/
type ChangeContainerNodeOKBodyContainer struct {

	// Docker container identifier. If specified, must be a unique Docker container identifier. Can't be changed.
	DockerContainerID string `json:"docker_container_id,omitempty"`

	// Container name. Can be changed.
	DockerContainerName string `json:"docker_container_name,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. If defined, Generic Node with that machine_id must exist. Can't be changed.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier, can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name, can be changed.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this change container node o k body container
func (o *ChangeContainerNodeOKBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeContainerNodeOKBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeContainerNodeOKBodyContainer) UnmarshalBinary(b []byte) error {
	var res ChangeContainerNodeOKBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

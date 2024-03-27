// Code generated by go-swagger; DO NOT EDIT.

package management_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegisterNodeReader is a Reader for the RegisterNode structure.
type RegisterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterNodeOK creates a RegisterNodeOK with default headers values
func NewRegisterNodeOK() *RegisterNodeOK {
	return &RegisterNodeOK{}
}

/*
RegisterNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type RegisterNodeOK struct {
	Payload *RegisterNodeOKBody
}

func (o *RegisterNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/nodes][%d] registerNodeOk  %+v", 200, o.Payload)
}

func (o *RegisterNodeOK) GetPayload() *RegisterNodeOKBody {
	return o.Payload
}

func (o *RegisterNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RegisterNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterNodeDefault creates a RegisterNodeDefault with default headers values
func NewRegisterNodeDefault(code int) *RegisterNodeDefault {
	return &RegisterNodeDefault{
		_statusCode: code,
	}
}

/*
RegisterNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RegisterNodeDefault struct {
	_statusCode int

	Payload *RegisterNodeDefaultBody
}

// Code gets the status code for the register node default response
func (o *RegisterNodeDefault) Code() int {
	return o._statusCode
}

func (o *RegisterNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/nodes][%d] RegisterNode default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterNodeDefault) GetPayload() *RegisterNodeDefaultBody {
	return o.Payload
}

func (o *RegisterNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RegisterNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RegisterNodeBody register node body
swagger:model RegisterNodeBody
*/
type RegisterNodeBody struct {
	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_UNSPECIFIED NODE_TYPE_GENERIC_NODE NODE_TYPE_CONTAINER_NODE NODE_TYPE_REMOTE_NODE NODE_TYPE_REMOTE_RDS_NODE NODE_TYPE_REMOTE_AZURE_DATABASE_NODE]
	NodeType *string `json:"node_type,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels for Node.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// If true, and Node with that name already exist, it will be removed with all dependent Services and Agents.
	Reregister bool `json:"reregister,omitempty"`

	// MetricsMode defines desired metrics mode for agent,
	// it can be pull, push or auto mode chosen by server.
	//
	//  - METRICS_MODE_UNSPECIFIED: Auto
	// Enum: [METRICS_MODE_UNSPECIFIED METRICS_MODE_PULL METRICS_MODE_PUSH]
	MetricsMode *string `json:"metrics_mode,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`

	// Optionally expose the exporter process on all public interfaces
	ExposeExporter bool `json:"expose_exporter,omitempty"`
}

// Validate validates this register node body
func (o *RegisterNodeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var registerNodeBodyTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_UNSPECIFIED","NODE_TYPE_GENERIC_NODE","NODE_TYPE_CONTAINER_NODE","NODE_TYPE_REMOTE_NODE","NODE_TYPE_REMOTE_RDS_NODE","NODE_TYPE_REMOTE_AZURE_DATABASE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registerNodeBodyTypeNodeTypePropEnum = append(registerNodeBodyTypeNodeTypePropEnum, v)
	}
}

const (

	// RegisterNodeBodyNodeTypeNODETYPEUNSPECIFIED captures enum value "NODE_TYPE_UNSPECIFIED"
	RegisterNodeBodyNodeTypeNODETYPEUNSPECIFIED string = "NODE_TYPE_UNSPECIFIED"

	// RegisterNodeBodyNodeTypeNODETYPEGENERICNODE captures enum value "NODE_TYPE_GENERIC_NODE"
	RegisterNodeBodyNodeTypeNODETYPEGENERICNODE string = "NODE_TYPE_GENERIC_NODE"

	// RegisterNodeBodyNodeTypeNODETYPECONTAINERNODE captures enum value "NODE_TYPE_CONTAINER_NODE"
	RegisterNodeBodyNodeTypeNODETYPECONTAINERNODE string = "NODE_TYPE_CONTAINER_NODE"

	// RegisterNodeBodyNodeTypeNODETYPEREMOTENODE captures enum value "NODE_TYPE_REMOTE_NODE"
	RegisterNodeBodyNodeTypeNODETYPEREMOTENODE string = "NODE_TYPE_REMOTE_NODE"

	// RegisterNodeBodyNodeTypeNODETYPEREMOTERDSNODE captures enum value "NODE_TYPE_REMOTE_RDS_NODE"
	RegisterNodeBodyNodeTypeNODETYPEREMOTERDSNODE string = "NODE_TYPE_REMOTE_RDS_NODE"

	// RegisterNodeBodyNodeTypeNODETYPEREMOTEAZUREDATABASENODE captures enum value "NODE_TYPE_REMOTE_AZURE_DATABASE_NODE"
	RegisterNodeBodyNodeTypeNODETYPEREMOTEAZUREDATABASENODE string = "NODE_TYPE_REMOTE_AZURE_DATABASE_NODE"
)

// prop value enum
func (o *RegisterNodeBody) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, registerNodeBodyTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RegisterNodeBody) validateNodeType(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

var registerNodeBodyTypeMetricsModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["METRICS_MODE_UNSPECIFIED","METRICS_MODE_PULL","METRICS_MODE_PUSH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registerNodeBodyTypeMetricsModePropEnum = append(registerNodeBodyTypeMetricsModePropEnum, v)
	}
}

const (

	// RegisterNodeBodyMetricsModeMETRICSMODEUNSPECIFIED captures enum value "METRICS_MODE_UNSPECIFIED"
	RegisterNodeBodyMetricsModeMETRICSMODEUNSPECIFIED string = "METRICS_MODE_UNSPECIFIED"

	// RegisterNodeBodyMetricsModeMETRICSMODEPULL captures enum value "METRICS_MODE_PULL"
	RegisterNodeBodyMetricsModeMETRICSMODEPULL string = "METRICS_MODE_PULL"

	// RegisterNodeBodyMetricsModeMETRICSMODEPUSH captures enum value "METRICS_MODE_PUSH"
	RegisterNodeBodyMetricsModeMETRICSMODEPUSH string = "METRICS_MODE_PUSH"
)

// prop value enum
func (o *RegisterNodeBody) validateMetricsModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, registerNodeBodyTypeMetricsModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RegisterNodeBody) validateMetricsMode(formats strfmt.Registry) error {
	if swag.IsZero(o.MetricsMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateMetricsModeEnum("body"+"."+"metrics_mode", "body", *o.MetricsMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this register node body based on context it is used
func (o *RegisterNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeBody) UnmarshalBinary(b []byte) error {
	var res RegisterNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RegisterNodeDefaultBody register node default body
swagger:model RegisterNodeDefaultBody
*/
type RegisterNodeDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RegisterNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this register node default body
func (o *RegisterNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RegisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RegisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this register node default body based on the context it is used
func (o *RegisterNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RegisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RegisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res RegisterNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RegisterNodeDefaultBodyDetailsItems0 register node default body details items0
swagger:model RegisterNodeDefaultBodyDetailsItems0
*/
type RegisterNodeDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this register node default body details items0
func (o *RegisterNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register node default body details items0 based on context it is used
func (o *RegisterNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RegisterNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RegisterNodeOKBody register node OK body
swagger:model RegisterNodeOKBody
*/
type RegisterNodeOKBody struct {
	// Token represents token for vmagent auth config.
	Token string `json:"token,omitempty"`

	// Warning message.
	Warning string `json:"warning,omitempty"`

	// container node
	ContainerNode *RegisterNodeOKBodyContainerNode `json:"container_node,omitempty"`

	// generic node
	GenericNode *RegisterNodeOKBodyGenericNode `json:"generic_node,omitempty"`

	// pmm agent
	PMMAgent *RegisterNodeOKBodyPMMAgent `json:"pmm_agent,omitempty"`
}

// Validate validates this register node OK body
func (o *RegisterNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainerNode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGenericNode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterNodeOKBody) validateContainerNode(formats strfmt.Registry) error {
	if swag.IsZero(o.ContainerNode) { // not required
		return nil
	}

	if o.ContainerNode != nil {
		if err := o.ContainerNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerNodeOk" + "." + "container_node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerNodeOk" + "." + "container_node")
			}
			return err
		}
	}

	return nil
}

func (o *RegisterNodeOKBody) validateGenericNode(formats strfmt.Registry) error {
	if swag.IsZero(o.GenericNode) { // not required
		return nil
	}

	if o.GenericNode != nil {
		if err := o.GenericNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerNodeOk" + "." + "generic_node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerNodeOk" + "." + "generic_node")
			}
			return err
		}
	}

	return nil
}

func (o *RegisterNodeOKBody) validatePMMAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	if o.PMMAgent != nil {
		if err := o.PMMAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerNodeOk" + "." + "pmm_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerNodeOk" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this register node OK body based on the context it is used
func (o *RegisterNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateContainerNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGenericNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePMMAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterNodeOKBody) contextValidateContainerNode(ctx context.Context, formats strfmt.Registry) error {
	if o.ContainerNode != nil {
		if err := o.ContainerNode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerNodeOk" + "." + "container_node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerNodeOk" + "." + "container_node")
			}
			return err
		}
	}

	return nil
}

func (o *RegisterNodeOKBody) contextValidateGenericNode(ctx context.Context, formats strfmt.Registry) error {
	if o.GenericNode != nil {
		if err := o.GenericNode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerNodeOk" + "." + "generic_node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerNodeOk" + "." + "generic_node")
			}
			return err
		}
	}

	return nil
}

func (o *RegisterNodeOKBody) contextValidatePMMAgent(ctx context.Context, formats strfmt.Registry) error {
	if o.PMMAgent != nil {
		if err := o.PMMAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerNodeOk" + "." + "pmm_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registerNodeOk" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeOKBody) UnmarshalBinary(b []byte) error {
	var res RegisterNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RegisterNodeOKBodyContainerNode ContainerNode represents a Docker container.
swagger:model RegisterNodeOKBodyContainerNode
*/
type RegisterNodeOKBodyContainerNode struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs.
	MachineID string `json:"machine_id,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this register node OK body container node
func (o *RegisterNodeOKBodyContainerNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register node OK body container node based on context it is used
func (o *RegisterNodeOKBodyContainerNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeOKBodyContainerNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeOKBodyContainerNode) UnmarshalBinary(b []byte) error {
	var res RegisterNodeOKBodyContainerNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RegisterNodeOKBodyGenericNode GenericNode represents a bare metal server or virtual machine.
swagger:model RegisterNodeOKBodyGenericNode
*/
type RegisterNodeOKBodyGenericNode struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this register node OK body generic node
func (o *RegisterNodeOKBodyGenericNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register node OK body generic node based on context it is used
func (o *RegisterNodeOKBodyGenericNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeOKBodyGenericNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeOKBodyGenericNode) UnmarshalBinary(b []byte) error {
	var res RegisterNodeOKBodyGenericNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RegisterNodeOKBodyPMMAgent PMMAgent runs on Generic or Container Node.
swagger:model RegisterNodeOKBodyPMMAgent
*/
type RegisterNodeOKBodyPMMAgent struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	Connected bool `json:"connected,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this register node OK body PMM agent
func (o *RegisterNodeOKBodyPMMAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register node OK body PMM agent based on context it is used
func (o *RegisterNodeOKBodyPMMAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterNodeOKBodyPMMAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterNodeOKBodyPMMAgent) UnmarshalBinary(b []byte) error {
	var res RegisterNodeOKBodyPMMAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

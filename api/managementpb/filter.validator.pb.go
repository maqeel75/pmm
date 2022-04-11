// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/filter.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Filter) Validate() error {
	if oneOfNester, ok := this.GetValue().(*Filter_IntValues); ok {
		if oneOfNester.IntValues != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.IntValues); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("IntValues", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*Filter_LongValues); ok {
		if oneOfNester.LongValues != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LongValues); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LongValues", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*Filter_StringValues); ok {
		if oneOfNester.StringValues != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StringValues); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StringValues", err)
			}
		}
	}
	if oneOfNester, ok := this.GetValue().(*Filter_IntRangeValues); ok {
		if oneOfNester.IntRangeValues != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.IntRangeValues); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("IntRangeValues", err)
			}
		}
	}
	return nil
}
func (this *IntValues) Validate() error {
	return nil
}
func (this *LongValues) Validate() error {
	return nil
}
func (this *StringValues) Validate() error {
	return nil
}
func (this *IntRangeValues) Validate() error {
	return nil
}
func (this *FilterParams) Validate() error {
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	return nil
}

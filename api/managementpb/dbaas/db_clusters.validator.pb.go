// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/dbaas/db_clusters.proto

package dbaasv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *PSMDBCluster) Validate() error {
	if this.Operation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operation", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *PXCCluster) Validate() error {
	if this.Operation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operation", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *ListDBClustersRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	return nil
}

func (this *ListDBClustersResponse) Validate() error {
	for _, item := range this.PxcClusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PxcClusters", err)
			}
		}
	}
	for _, item := range this.PsmdbClusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PsmdbClusters", err)
			}
		}
	}
	return nil
}

func (this *RestartDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}

func (this *RestartDBClusterResponse) Validate() error {
	return nil
}

func (this *DeleteDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}

func (this *DeleteDBClusterResponse) Validate() error {
	return nil
}

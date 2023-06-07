// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: managementpb/backup/artifacts.proto

package backupv1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BackupStatus shows the current status of execution of backup.
type BackupStatus int32

const (
	BackupStatus_BACKUP_STATUS_INVALID             BackupStatus = 0
	BackupStatus_BACKUP_STATUS_PENDING             BackupStatus = 1
	BackupStatus_BACKUP_STATUS_IN_PROGRESS         BackupStatus = 2
	BackupStatus_BACKUP_STATUS_PAUSED              BackupStatus = 3
	BackupStatus_BACKUP_STATUS_SUCCESS             BackupStatus = 4
	BackupStatus_BACKUP_STATUS_ERROR               BackupStatus = 5
	BackupStatus_BACKUP_STATUS_DELETING            BackupStatus = 6
	BackupStatus_BACKUP_STATUS_FAILED_TO_DELETE    BackupStatus = 7
	BackupStatus_BACKUP_STATUS_CLEANUP_IN_PROGRESS BackupStatus = 8
)

// Enum value maps for BackupStatus.
var (
	BackupStatus_name = map[int32]string{
		0: "BACKUP_STATUS_INVALID",
		1: "BACKUP_STATUS_PENDING",
		2: "BACKUP_STATUS_IN_PROGRESS",
		3: "BACKUP_STATUS_PAUSED",
		4: "BACKUP_STATUS_SUCCESS",
		5: "BACKUP_STATUS_ERROR",
		6: "BACKUP_STATUS_DELETING",
		7: "BACKUP_STATUS_FAILED_TO_DELETE",
		8: "BACKUP_STATUS_CLEANUP_IN_PROGRESS",
	}
	BackupStatus_value = map[string]int32{
		"BACKUP_STATUS_INVALID":             0,
		"BACKUP_STATUS_PENDING":             1,
		"BACKUP_STATUS_IN_PROGRESS":         2,
		"BACKUP_STATUS_PAUSED":              3,
		"BACKUP_STATUS_SUCCESS":             4,
		"BACKUP_STATUS_ERROR":               5,
		"BACKUP_STATUS_DELETING":            6,
		"BACKUP_STATUS_FAILED_TO_DELETE":    7,
		"BACKUP_STATUS_CLEANUP_IN_PROGRESS": 8,
	}
)

func (x BackupStatus) Enum() *BackupStatus {
	p := new(BackupStatus)
	*p = x
	return p
}

func (x BackupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_backup_artifacts_proto_enumTypes[0].Descriptor()
}

func (BackupStatus) Type() protoreflect.EnumType {
	return &file_managementpb_backup_artifacts_proto_enumTypes[0]
}

func (x BackupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackupStatus.Descriptor instead.
func (BackupStatus) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{0}
}

// Artifact represents single backup artifact.
type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Machine-readable artifact ID.
	ArtifactId string `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	// Artifact name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `protobuf:"bytes,3,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Machine-readable location ID.
	LocationId string `protobuf:"bytes,4,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Location name.
	LocationName string `protobuf:"bytes,5,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	// Machine-readable service ID.
	ServiceId string `protobuf:"bytes,6,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Service name.
	ServiceName string `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Backup data model.
	DataModel DataModel `protobuf:"varint,8,opt,name=data_model,json=dataModel,proto3,enum=backup.v1.DataModel" json:"data_model,omitempty"`
	// Backup status.
	Status BackupStatus `protobuf:"varint,9,opt,name=status,proto3,enum=backup.v1.BackupStatus" json:"status,omitempty"`
	// Artifact creation time.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Backup mode.
	Mode BackupMode `protobuf:"varint,11,opt,name=mode,proto3,enum=backup.v1.BackupMode" json:"mode,omitempty"`
	// Source database setup type.
	IsShardedCluster bool `protobuf:"varint,12,opt,name=is_sharded_cluster,json=isShardedCluster,proto3" json:"is_sharded_cluster,omitempty"`
	// Folder to store artifact on a storage.
	Folder string `protobuf:"bytes,13,opt,name=folder,proto3" json:"folder,omitempty"`
	// List of artifact metadata.
	MetadataList []*Metadata `protobuf:"bytes,14,rep,name=metadata_list,json=metadataList,proto3" json:"metadata_list,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{0}
}

func (x *Artifact) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

func (x *Artifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artifact) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *Artifact) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Artifact) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *Artifact) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Artifact) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Artifact) GetDataModel() DataModel {
	if x != nil {
		return x.DataModel
	}
	return DataModel_DATA_MODEL_INVALID
}

func (x *Artifact) GetStatus() BackupStatus {
	if x != nil {
		return x.Status
	}
	return BackupStatus_BACKUP_STATUS_INVALID
}

func (x *Artifact) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Artifact) GetMode() BackupMode {
	if x != nil {
		return x.Mode
	}
	return BackupMode_BACKUP_MODE_INVALID
}

func (x *Artifact) GetIsShardedCluster() bool {
	if x != nil {
		return x.IsShardedCluster
	}
	return false
}

func (x *Artifact) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

func (x *Artifact) GetMetadataList() []*Metadata {
	if x != nil {
		return x.MetadataList
	}
	return nil
}

type ListArtifactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListArtifactsRequest) Reset() {
	*x = ListArtifactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtifactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsRequest) ProtoMessage() {}

func (x *ListArtifactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsRequest.ProtoReflect.Descriptor instead.
func (*ListArtifactsRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{1}
}

type ListArtifactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artifacts []*Artifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
}

func (x *ListArtifactsResponse) Reset() {
	*x = ListArtifactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtifactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsResponse) ProtoMessage() {}

func (x *ListArtifactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsResponse.ProtoReflect.Descriptor instead.
func (*ListArtifactsResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{2}
}

func (x *ListArtifactsResponse) GetArtifacts() []*Artifact {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

type DeleteArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Machine-readable artifact ID.
	ArtifactId string `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	// Removes all the backup files associated with artifact if flag is set.
	RemoveFiles bool `protobuf:"varint,2,opt,name=remove_files,json=removeFiles,proto3" json:"remove_files,omitempty"`
}

func (x *DeleteArtifactRequest) Reset() {
	*x = DeleteArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtifactRequest) ProtoMessage() {}

func (x *DeleteArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtifactRequest.ProtoReflect.Descriptor instead.
func (*DeleteArtifactRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteArtifactRequest) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

func (x *DeleteArtifactRequest) GetRemoveFiles() bool {
	if x != nil {
		return x.RemoveFiles
	}
	return false
}

type DeleteArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteArtifactResponse) Reset() {
	*x = DeleteArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtifactResponse) ProtoMessage() {}

func (x *DeleteArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtifactResponse.ProtoReflect.Descriptor instead.
func (*DeleteArtifactResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{4}
}

type PitrTimerange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_timestamp is the time of the first event in the PITR chunk.
	StartTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	// end_timestamp is the time of the last event in the PITR chunk.
	EndTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
}

func (x *PitrTimerange) Reset() {
	*x = PitrTimerange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PitrTimerange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PitrTimerange) ProtoMessage() {}

func (x *PitrTimerange) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PitrTimerange.ProtoReflect.Descriptor instead.
func (*PitrTimerange) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{5}
}

func (x *PitrTimerange) GetStartTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTimestamp
	}
	return nil
}

func (x *PitrTimerange) GetEndTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTimestamp
	}
	return nil
}

type ListPitrTimerangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Artifact ID represents artifact whose location has PITR timeranges to be retrieved.
	ArtifactId string `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
}

func (x *ListPitrTimerangesRequest) Reset() {
	*x = ListPitrTimerangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPitrTimerangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPitrTimerangesRequest) ProtoMessage() {}

func (x *ListPitrTimerangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPitrTimerangesRequest.ProtoReflect.Descriptor instead.
func (*ListPitrTimerangesRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{6}
}

func (x *ListPitrTimerangesRequest) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

type ListPitrTimerangesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeranges []*PitrTimerange `protobuf:"bytes,1,rep,name=timeranges,proto3" json:"timeranges,omitempty"`
}

func (x *ListPitrTimerangesResponse) Reset() {
	*x = ListPitrTimerangesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_artifacts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPitrTimerangesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPitrTimerangesResponse) ProtoMessage() {}

func (x *ListPitrTimerangesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_artifacts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPitrTimerangesResponse.ProtoReflect.Descriptor instead.
func (*ListPitrTimerangesResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_artifacts_proto_rawDescGZIP(), []int{7}
}

func (x *ListPitrTimerangesResponse) GetTimeranges() []*PitrTimerange {
	if x != nil {
		return x.Timeranges
	}
	return nil
}

var File_managementpb_backup_artifacts_proto protoreflect.FileDescriptor

var file_managementpb_backup_artifacts_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x04, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x22, 0x5b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x50,
	0x69, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x3c, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x74, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x56, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x69, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2a, 0x98, 0x02, 0x0a, 0x0c, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x43,
	0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x43, 0x4b,
	0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16,
	0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x43, 0x4b,
	0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21,
	0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c,
	0x45, 0x41, 0x4e, 0x55, 0x50, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x08, 0x32, 0xbf, 0x03, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x12, 0x83, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x88, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x74, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x74, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x69, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x3a, 0x01,
	0x2a, 0x22, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x49, 0x54, 0x52, 0x54, 0x69, 0x6d, 0x65, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x9d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_backup_artifacts_proto_rawDescOnce sync.Once
	file_managementpb_backup_artifacts_proto_rawDescData = file_managementpb_backup_artifacts_proto_rawDesc
)

func file_managementpb_backup_artifacts_proto_rawDescGZIP() []byte {
	file_managementpb_backup_artifacts_proto_rawDescOnce.Do(func() {
		file_managementpb_backup_artifacts_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_backup_artifacts_proto_rawDescData)
	})
	return file_managementpb_backup_artifacts_proto_rawDescData
}

var (
	file_managementpb_backup_artifacts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_backup_artifacts_proto_msgTypes  = make([]protoimpl.MessageInfo, 8)
	file_managementpb_backup_artifacts_proto_goTypes   = []interface{}{
		(BackupStatus)(0),                  // 0: backup.v1.BackupStatus
		(*Artifact)(nil),                   // 1: backup.v1.Artifact
		(*ListArtifactsRequest)(nil),       // 2: backup.v1.ListArtifactsRequest
		(*ListArtifactsResponse)(nil),      // 3: backup.v1.ListArtifactsResponse
		(*DeleteArtifactRequest)(nil),      // 4: backup.v1.DeleteArtifactRequest
		(*DeleteArtifactResponse)(nil),     // 5: backup.v1.DeleteArtifactResponse
		(*PitrTimerange)(nil),              // 6: backup.v1.PitrTimerange
		(*ListPitrTimerangesRequest)(nil),  // 7: backup.v1.ListPitrTimerangesRequest
		(*ListPitrTimerangesResponse)(nil), // 8: backup.v1.ListPitrTimerangesResponse
		(DataModel)(0),                     // 9: backup.v1.DataModel
		(*timestamppb.Timestamp)(nil),      // 10: google.protobuf.Timestamp
		(BackupMode)(0),                    // 11: backup.v1.BackupMode
		(*Metadata)(nil),                   // 12: backup.v1.Metadata
	}
)

var file_managementpb_backup_artifacts_proto_depIdxs = []int32{
	9,  // 0: backup.v1.Artifact.data_model:type_name -> backup.v1.DataModel
	0,  // 1: backup.v1.Artifact.status:type_name -> backup.v1.BackupStatus
	10, // 2: backup.v1.Artifact.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: backup.v1.Artifact.mode:type_name -> backup.v1.BackupMode
	12, // 4: backup.v1.Artifact.metadata_list:type_name -> backup.v1.Metadata
	1,  // 5: backup.v1.ListArtifactsResponse.artifacts:type_name -> backup.v1.Artifact
	10, // 6: backup.v1.PitrTimerange.start_timestamp:type_name -> google.protobuf.Timestamp
	10, // 7: backup.v1.PitrTimerange.end_timestamp:type_name -> google.protobuf.Timestamp
	6,  // 8: backup.v1.ListPitrTimerangesResponse.timeranges:type_name -> backup.v1.PitrTimerange
	2,  // 9: backup.v1.Artifacts.ListArtifacts:input_type -> backup.v1.ListArtifactsRequest
	4,  // 10: backup.v1.Artifacts.DeleteArtifact:input_type -> backup.v1.DeleteArtifactRequest
	7,  // 11: backup.v1.Artifacts.ListPitrTimeranges:input_type -> backup.v1.ListPitrTimerangesRequest
	3,  // 12: backup.v1.Artifacts.ListArtifacts:output_type -> backup.v1.ListArtifactsResponse
	5,  // 13: backup.v1.Artifacts.DeleteArtifact:output_type -> backup.v1.DeleteArtifactResponse
	8,  // 14: backup.v1.Artifacts.ListPitrTimeranges:output_type -> backup.v1.ListPitrTimerangesResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_managementpb_backup_artifacts_proto_init() }
func file_managementpb_backup_artifacts_proto_init() {
	if File_managementpb_backup_artifacts_proto != nil {
		return
	}
	file_managementpb_backup_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_backup_artifacts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtifactsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtifactsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArtifactRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArtifactResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PitrTimerange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPitrTimerangesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_backup_artifacts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPitrTimerangesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_backup_artifacts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_backup_artifacts_proto_goTypes,
		DependencyIndexes: file_managementpb_backup_artifacts_proto_depIdxs,
		EnumInfos:         file_managementpb_backup_artifacts_proto_enumTypes,
		MessageInfos:      file_managementpb_backup_artifacts_proto_msgTypes,
	}.Build()
	File_managementpb_backup_artifacts_proto = out.File
	file_managementpb_backup_artifacts_proto_rawDesc = nil
	file_managementpb_backup_artifacts_proto_goTypes = nil
	file_managementpb_backup_artifacts_proto_depIdxs = nil
}

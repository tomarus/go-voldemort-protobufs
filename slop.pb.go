// Code generated by protoc-gen-go. DO NOT EDIT.
// source: slop.proto

/*
Package voldemort is a generated protocol buffer package.

It is generated from these files:
	slop.proto
	voldemort-admin.proto
	voldemort-client.proto

It has these top-level messages:
	Slop
	GetMetadataRequest
	GetMetadataResponse
	UpdateMetadataRequest
	UpdateStoreDefinitions
	UpdateMetadataPairRequest
	UpdateMetadataPairResponse
	UpdateMetadataResponse
	FileEntry
	PartitionEntry
	UpdatePartitionEntriesRequest
	UpdatePartitionEntriesResponse
	VoldemortFilter
	UpdateSlopEntriesRequest
	UpdateSlopEntriesResponse
	FetchPartitionFilesRequest
	FetchPartitionEntriesRequest
	FetchPartitionEntriesResponse
	DeletePartitionEntriesRequest
	DeletePartitionEntriesResponse
	InitiateFetchAndUpdateRequest
	AsyncOperationStatusRequest
	AsyncOperationStatusResponse
	AsyncOperationStopRequest
	AsyncOperationStopResponse
	AsyncOperationListRequest
	AsyncOperationListResponse
	ListScheduledJobsRequest
	ListScheduledJobsResponse
	GetScheduledJobStatusRequest
	GetScheduledJobStatusResponse
	StopScheduledJobRequest
	StopScheduledJobResponse
	EnableScheduledJobRequest
	EnableScheduledJobResponse
	PartitionTuple
	PerStorePartitionTuple
	RebalancePartitionInfoMap
	StoreToPartitionsIds
	RebalanceTaskInfoMap
	InitiateRebalanceNodeRequest
	InitiateRebalanceNodeOnDonorRequest
	TruncateEntriesRequest
	TruncateEntriesResponse
	AddStoreRequest
	AddStoreResponse
	DeleteStoreRequest
	DeleteStoreResponse
	FetchStoreRequest
	SwapStoreRequest
	SwapStoreResponse
	RollbackStoreRequest
	RollbackStoreResponse
	RepairJobRequest
	RepairJobResponse
	PruneJobRequest
	PruneJobResponse
	SlopPurgeJobRequest
	SlopPurgeJobResponse
	ROStoreVersionDirMap
	GetROMaxVersionDirRequest
	GetROMaxVersionDirResponse
	GetROCurrentVersionDirRequest
	GetROCurrentVersionDirResponse
	GetROStorageFormatRequest
	GetROStorageFormatResponse
	GetROStorageFileListRequest
	GetROStorageFileListResponse
	GetROStorageCompressionCodecListRequest
	GetROStorageCompressionCodecListResponse
	FailedFetchStoreRequest
	FailedFetchStoreResponse
	RebalanceStateChangeRequest
	RebalanceStateChangeResponse
	DeleteStoreRebalanceStateRequest
	DeleteStoreRebalanceStateResponse
	SetOfflineStateRequest
	SetOfflineStateResponse
	NativeBackupRequest
	ReserveMemoryRequest
	ReserveMemoryResponse
	GetHighAvailabilitySettingsRequest
	GetHighAvailabilitySettingsResponse
	DisableStoreVersionRequest
	DisableStoreVersionResponse
	HandleFetchFailureRequest
	HandleFetchFailureResponse
	GetConfigRequest
	GetConfigResponse
	MapFieldEntry
	VoldemortAdminRequest
	ClockEntry
	VectorClock
	Versioned
	Error
	KeyedVersions
	GetRequest
	GetResponse
	GetVersionResponse
	GetAllRequest
	GetAllResponse
	PutRequest
	PutResponse
	DeleteRequest
	DeleteResponse
	VoldemortRequest
*/
package voldemort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Slop struct {
	Store            *string `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	Operation        *string `protobuf:"bytes,2,opt,name=operation" json:"operation,omitempty"`
	Key              []byte  `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value            []byte  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	NodeId           *int32  `protobuf:"varint,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Arrived          *int64  `protobuf:"varint,6,opt,name=arrived" json:"arrived,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Slop) Reset()                    { *m = Slop{} }
func (m *Slop) String() string            { return proto.CompactTextString(m) }
func (*Slop) ProtoMessage()               {}
func (*Slop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Slop) GetStore() string {
	if m != nil && m.Store != nil {
		return *m.Store
	}
	return ""
}

func (m *Slop) GetOperation() string {
	if m != nil && m.Operation != nil {
		return *m.Operation
	}
	return ""
}

func (m *Slop) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Slop) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Slop) GetNodeId() int32 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *Slop) GetArrived() int64 {
	if m != nil && m.Arrived != nil {
		return *m.Arrived
	}
	return 0
}

func init() {
	proto.RegisterType((*Slop)(nil), "voldemort.Slop")
}

func init() { proto.RegisterFile("slop.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x3d, 0x8a, 0xc3, 0x30,
	0x10, 0x85, 0xd1, 0xfa, 0x0f, 0x0f, 0x5b, 0x2c, 0x62, 0xc1, 0x2a, 0xb6, 0x10, 0xdb, 0x44, 0x95,
	0x0f, 0xe1, 0x2a, 0xe9, 0x82, 0x02, 0x69, 0x83, 0x40, 0x53, 0x88, 0x28, 0x1e, 0x21, 0x2b, 0x86,
	0xe4, 0x1e, 0xb9, 0x6f, 0x90, 0x4d, 0x92, 0x6e, 0xbe, 0x37, 0x30, 0xf3, 0x3d, 0x80, 0xc9, 0x53,
	0xe8, 0x43, 0xa4, 0x44, 0xbc, 0x9d, 0xc9, 0x5b, 0xbc, 0x50, 0x4c, 0xff, 0x0f, 0x06, 0xe5, 0xc1,
	0x53, 0xe0, 0xbf, 0x50, 0x4d, 0x89, 0x22, 0x0a, 0x26, 0x99, 0x6a, 0xf5, 0x0a, 0xfc, 0x0f, 0x5a,
	0x0a, 0x18, 0x4d, 0x72, 0x34, 0x8a, 0xaf, 0x65, 0xf3, 0x09, 0xf8, 0x0f, 0x14, 0x67, 0xbc, 0x89,
	0x42, 0x32, 0xf5, 0xad, 0xf3, 0x98, 0xaf, 0xcc, 0xc6, 0x5f, 0x51, 0x94, 0x4b, 0xb6, 0x02, 0xef,
	0xa0, 0x19, 0xc9, 0xe2, 0xc9, 0x59, 0x51, 0x49, 0xa6, 0x2a, 0x5d, 0x67, 0xdc, 0x59, 0x2e, 0xa0,
	0x31, 0x31, 0xba, 0x19, 0xad, 0xa8, 0x25, 0x53, 0x85, 0x7e, 0xe1, 0xb0, 0x81, 0xee, 0x2d, 0xd9,
	0x4f, 0x18, 0x9d, 0xf1, 0xee, 0xbe, 0x7c, 0x1d, 0xe0, 0x98, 0x85, 0xf7, 0xb9, 0xc9, 0x96, 0x3d,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x85, 0x57, 0x7f, 0xe4, 0xd8, 0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/proto/resourcespb/network_interface_security_group_association.proto

package resourcespb

import (
	commonpb "github.com/multycloud/multy/api/proto/commonpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNetworkInterfaceSecurityGroupAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *NetworkInterfaceSecurityGroupAssociationArgs `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CreateNetworkInterfaceSecurityGroupAssociationRequest) Reset() {
	*x = CreateNetworkInterfaceSecurityGroupAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkInterfaceSecurityGroupAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkInterfaceSecurityGroupAssociationRequest) ProtoMessage() {}

func (x *CreateNetworkInterfaceSecurityGroupAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkInterfaceSecurityGroupAssociationRequest.ProtoReflect.Descriptor instead.
func (*CreateNetworkInterfaceSecurityGroupAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNetworkInterfaceSecurityGroupAssociationRequest) GetResource() *NetworkInterfaceSecurityGroupAssociationArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ReadNetworkInterfaceSecurityGroupAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadNetworkInterfaceSecurityGroupAssociationRequest) Reset() {
	*x = ReadNetworkInterfaceSecurityGroupAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadNetworkInterfaceSecurityGroupAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadNetworkInterfaceSecurityGroupAssociationRequest) ProtoMessage() {}

func (x *ReadNetworkInterfaceSecurityGroupAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadNetworkInterfaceSecurityGroupAssociationRequest.ProtoReflect.Descriptor instead.
func (*ReadNetworkInterfaceSecurityGroupAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP(), []int{1}
}

func (x *ReadNetworkInterfaceSecurityGroupAssociationRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateNetworkInterfaceSecurityGroupAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string                                        `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resource   *NetworkInterfaceSecurityGroupAssociationArgs `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateNetworkInterfaceSecurityGroupAssociationRequest) Reset() {
	*x = UpdateNetworkInterfaceSecurityGroupAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNetworkInterfaceSecurityGroupAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNetworkInterfaceSecurityGroupAssociationRequest) ProtoMessage() {}

func (x *UpdateNetworkInterfaceSecurityGroupAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNetworkInterfaceSecurityGroupAssociationRequest.ProtoReflect.Descriptor instead.
func (*UpdateNetworkInterfaceSecurityGroupAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNetworkInterfaceSecurityGroupAssociationRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateNetworkInterfaceSecurityGroupAssociationRequest) GetResource() *NetworkInterfaceSecurityGroupAssociationArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteNetworkInterfaceSecurityGroupAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteNetworkInterfaceSecurityGroupAssociationRequest) Reset() {
	*x = DeleteNetworkInterfaceSecurityGroupAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkInterfaceSecurityGroupAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkInterfaceSecurityGroupAssociationRequest) ProtoMessage() {}

func (x *DeleteNetworkInterfaceSecurityGroupAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkInterfaceSecurityGroupAssociationRequest.ProtoReflect.Descriptor instead.
func (*DeleteNetworkInterfaceSecurityGroupAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNetworkInterfaceSecurityGroupAssociationRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type NetworkInterfaceSecurityGroupAssociationArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters   *commonpb.ChildResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	NetworkInterfaceId string                            `protobuf:"bytes,2,opt,name=network_interface_id,json=networkInterfaceId,proto3" json:"network_interface_id,omitempty"`
	SecurityGroupId    string                            `protobuf:"bytes,3,opt,name=security_group_id,json=securityGroupId,proto3" json:"security_group_id,omitempty"`
}

func (x *NetworkInterfaceSecurityGroupAssociationArgs) Reset() {
	*x = NetworkInterfaceSecurityGroupAssociationArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterfaceSecurityGroupAssociationArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceSecurityGroupAssociationArgs) ProtoMessage() {}

func (x *NetworkInterfaceSecurityGroupAssociationArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceSecurityGroupAssociationArgs.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceSecurityGroupAssociationArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkInterfaceSecurityGroupAssociationArgs) GetCommonParameters() *commonpb.ChildResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *NetworkInterfaceSecurityGroupAssociationArgs) GetNetworkInterfaceId() string {
	if x != nil {
		return x.NetworkInterfaceId
	}
	return ""
}

func (x *NetworkInterfaceSecurityGroupAssociationArgs) GetSecurityGroupId() string {
	if x != nil {
		return x.SecurityGroupId
	}
	return ""
}

type NetworkInterfaceSecurityGroupAssociationResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters   *commonpb.CommonChildResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	NetworkInterfaceId string                                  `protobuf:"bytes,2,opt,name=network_interface_id,json=networkInterfaceId,proto3" json:"network_interface_id,omitempty"`
	SecurityGroupId    string                                  `protobuf:"bytes,3,opt,name=security_group_id,json=securityGroupId,proto3" json:"security_group_id,omitempty"`
}

func (x *NetworkInterfaceSecurityGroupAssociationResource) Reset() {
	*x = NetworkInterfaceSecurityGroupAssociationResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterfaceSecurityGroupAssociationResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceSecurityGroupAssociationResource) ProtoMessage() {}

func (x *NetworkInterfaceSecurityGroupAssociationResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceSecurityGroupAssociationResource.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceSecurityGroupAssociationResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkInterfaceSecurityGroupAssociationResource) GetCommonParameters() *commonpb.CommonChildResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *NetworkInterfaceSecurityGroupAssociationResource) GetNetworkInterfaceId() string {
	if x != nil {
		return x.NetworkInterfaceId
	}
	return ""
}

func (x *NetworkInterfaceSecurityGroupAssociationResource) GetSecurityGroupId() string {
	if x != nil {
		return x.SecurityGroupId
	}
	return ""
}

var File_api_proto_resourcespb_network_interface_security_group_association_proto protoreflect.FileDescriptor

var file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDesc = []byte{
	0x0a, 0x48, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x96, 0x01, 0x0a, 0x35, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x33, 0x52, 0x65, 0x61,
	0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x22, 0xb7, 0x01, 0x0a, 0x35, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x5d, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67,
	0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x35, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xe4, 0x01, 0x0a, 0x2c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x56, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x10, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30,
	0x0a, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xee, 0x01, 0x0a,
	0x30, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x5c, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x5e, 0x0a,
	0x17, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescOnce sync.Once
	file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescData = file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDesc
)

func file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescGZIP() []byte {
	file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescOnce.Do(func() {
		file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescData)
	})
	return file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDescData
}

var file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_resourcespb_network_interface_security_group_association_proto_goTypes = []interface{}{
	(*CreateNetworkInterfaceSecurityGroupAssociationRequest)(nil), // 0: dev.multy.resources.CreateNetworkInterfaceSecurityGroupAssociationRequest
	(*ReadNetworkInterfaceSecurityGroupAssociationRequest)(nil),   // 1: dev.multy.resources.ReadNetworkInterfaceSecurityGroupAssociationRequest
	(*UpdateNetworkInterfaceSecurityGroupAssociationRequest)(nil), // 2: dev.multy.resources.UpdateNetworkInterfaceSecurityGroupAssociationRequest
	(*DeleteNetworkInterfaceSecurityGroupAssociationRequest)(nil), // 3: dev.multy.resources.DeleteNetworkInterfaceSecurityGroupAssociationRequest
	(*NetworkInterfaceSecurityGroupAssociationArgs)(nil),          // 4: dev.multy.resources.NetworkInterfaceSecurityGroupAssociationArgs
	(*NetworkInterfaceSecurityGroupAssociationResource)(nil),      // 5: dev.multy.resources.NetworkInterfaceSecurityGroupAssociationResource
	(*commonpb.ChildResourceCommonArgs)(nil),                      // 6: dev.multy.common.ChildResourceCommonArgs
	(*commonpb.CommonChildResourceParameters)(nil),                // 7: dev.multy.common.CommonChildResourceParameters
}
var file_api_proto_resourcespb_network_interface_security_group_association_proto_depIdxs = []int32{
	4, // 0: dev.multy.resources.CreateNetworkInterfaceSecurityGroupAssociationRequest.resource:type_name -> dev.multy.resources.NetworkInterfaceSecurityGroupAssociationArgs
	4, // 1: dev.multy.resources.UpdateNetworkInterfaceSecurityGroupAssociationRequest.resource:type_name -> dev.multy.resources.NetworkInterfaceSecurityGroupAssociationArgs
	6, // 2: dev.multy.resources.NetworkInterfaceSecurityGroupAssociationArgs.common_parameters:type_name -> dev.multy.common.ChildResourceCommonArgs
	7, // 3: dev.multy.resources.NetworkInterfaceSecurityGroupAssociationResource.common_parameters:type_name -> dev.multy.common.CommonChildResourceParameters
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_resourcespb_network_interface_security_group_association_proto_init() }
func file_api_proto_resourcespb_network_interface_security_group_association_proto_init() {
	if File_api_proto_resourcespb_network_interface_security_group_association_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkInterfaceSecurityGroupAssociationRequest); i {
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
		file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadNetworkInterfaceSecurityGroupAssociationRequest); i {
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
		file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNetworkInterfaceSecurityGroupAssociationRequest); i {
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
		file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkInterfaceSecurityGroupAssociationRequest); i {
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
		file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterfaceSecurityGroupAssociationArgs); i {
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
		file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInterfaceSecurityGroupAssociationResource); i {
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
			RawDescriptor: file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resourcespb_network_interface_security_group_association_proto_goTypes,
		DependencyIndexes: file_api_proto_resourcespb_network_interface_security_group_association_proto_depIdxs,
		MessageInfos:      file_api_proto_resourcespb_network_interface_security_group_association_proto_msgTypes,
	}.Build()
	File_api_proto_resourcespb_network_interface_security_group_association_proto = out.File
	file_api_proto_resourcespb_network_interface_security_group_association_proto_rawDesc = nil
	file_api_proto_resourcespb_network_interface_security_group_association_proto_goTypes = nil
	file_api_proto_resourcespb_network_interface_security_group_association_proto_depIdxs = nil
}
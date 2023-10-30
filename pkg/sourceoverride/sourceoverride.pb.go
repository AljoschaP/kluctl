// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: sourceoverride.proto

package sourceoverride

import (
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

type ProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *AuthMsg                 `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Response *ResolveOverrideResponse `protobuf:"bytes,2,opt,name=response,proto3,oneof" json:"response,omitempty"`
}

func (x *ProxyResponse) Reset() {
	*x = ProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourceoverride_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyResponse) ProtoMessage() {}

func (x *ProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sourceoverride_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyResponse.ProtoReflect.Descriptor instead.
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return file_sourceoverride_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyResponse) GetAuth() *AuthMsg {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ProxyResponse) GetResponse() *ResolveOverrideResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth    *AuthMsg                `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Request *ResolveOverrideRequest `protobuf:"bytes,2,opt,name=request,proto3,oneof" json:"request,omitempty"`
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourceoverride_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sourceoverride_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRequest.ProtoReflect.Descriptor instead.
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return file_sourceoverride_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyRequest) GetAuth() *AuthMsg {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ProxyRequest) GetRequest() *ResolveOverrideRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type AuthMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey     []byte  `protobuf:"bytes,1,opt,name=pubKey,proto3,oneof" json:"pubKey,omitempty"`
	Challenge  []byte  `protobuf:"bytes,2,opt,name=challenge,proto3,oneof" json:"challenge,omitempty"`
	Pop        []byte  `protobuf:"bytes,3,opt,name=pop,proto3,oneof" json:"pop,omitempty"`
	AuthError  *string `protobuf:"bytes,4,opt,name=authError,proto3,oneof" json:"authError,omitempty"`
	PubKeyHash *string `protobuf:"bytes,5,opt,name=pubKeyHash,proto3,oneof" json:"pubKeyHash,omitempty"`
}

func (x *AuthMsg) Reset() {
	*x = AuthMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourceoverride_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMsg) ProtoMessage() {}

func (x *AuthMsg) ProtoReflect() protoreflect.Message {
	mi := &file_sourceoverride_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMsg.ProtoReflect.Descriptor instead.
func (*AuthMsg) Descriptor() ([]byte, []int) {
	return file_sourceoverride_proto_rawDescGZIP(), []int{2}
}

func (x *AuthMsg) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *AuthMsg) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

func (x *AuthMsg) GetPop() []byte {
	if x != nil {
		return x.Pop
	}
	return nil
}

func (x *AuthMsg) GetAuthError() string {
	if x != nil && x.AuthError != nil {
		return *x.AuthError
	}
	return ""
}

func (x *AuthMsg) GetPubKeyHash() string {
	if x != nil && x.PubKeyHash != nil {
		return *x.PubKeyHash
	}
	return ""
}

type ResolveOverrideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoKey string `protobuf:"bytes,1,opt,name=repoKey,proto3" json:"repoKey,omitempty"`
}

func (x *ResolveOverrideRequest) Reset() {
	*x = ResolveOverrideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourceoverride_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveOverrideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveOverrideRequest) ProtoMessage() {}

func (x *ResolveOverrideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sourceoverride_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveOverrideRequest.ProtoReflect.Descriptor instead.
func (*ResolveOverrideRequest) Descriptor() ([]byte, []int) {
	return file_sourceoverride_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveOverrideRequest) GetRepoKey() string {
	if x != nil {
		return x.RepoKey
	}
	return ""
}

type ResolveOverrideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    *string `protobuf:"bytes,1,opt,name=error,proto3,oneof" json:"error,omitempty"`
	Artifact []byte  `protobuf:"bytes,2,opt,name=artifact,proto3,oneof" json:"artifact,omitempty"`
}

func (x *ResolveOverrideResponse) Reset() {
	*x = ResolveOverrideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourceoverride_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveOverrideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveOverrideResponse) ProtoMessage() {}

func (x *ResolveOverrideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sourceoverride_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveOverrideResponse.ProtoReflect.Descriptor instead.
func (*ResolveOverrideResponse) Descriptor() ([]byte, []int) {
	return file_sourceoverride_proto_rawDescGZIP(), []int{4}
}

func (x *ResolveOverrideResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

func (x *ResolveOverrideResponse) GetArtifact() []byte {
	if x != nil {
		return x.Artifact
	}
	return nil
}

var File_sourceoverride_proto protoreflect.FileDescriptor

var file_sourceoverride_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x73, 0x67, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x48, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8e, 0x01, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x45, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe6, 0x01,
	0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x70, 0x6f, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x03, 0x70, 0x6f, 0x70, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x70, 0x6f, 0x70, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x4b, 0x65, 0x79, 0x22, 0x6c, 0x0a, 0x17, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x01, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x32, 0xb5, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x12, 0x50, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1d, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x6c, 0x75, 0x63, 0x74, 0x6c, 0x2f, 0x6b, 0x6c, 0x75, 0x63, 0x74, 0x6c, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sourceoverride_proto_rawDescOnce sync.Once
	file_sourceoverride_proto_rawDescData = file_sourceoverride_proto_rawDesc
)

func file_sourceoverride_proto_rawDescGZIP() []byte {
	file_sourceoverride_proto_rawDescOnce.Do(func() {
		file_sourceoverride_proto_rawDescData = protoimpl.X.CompressGZIP(file_sourceoverride_proto_rawDescData)
	})
	return file_sourceoverride_proto_rawDescData
}

var file_sourceoverride_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sourceoverride_proto_goTypes = []interface{}{
	(*ProxyResponse)(nil),           // 0: sourceoverride.ProxyResponse
	(*ProxyRequest)(nil),            // 1: sourceoverride.ProxyRequest
	(*AuthMsg)(nil),                 // 2: sourceoverride.AuthMsg
	(*ResolveOverrideRequest)(nil),  // 3: sourceoverride.ResolveOverrideRequest
	(*ResolveOverrideResponse)(nil), // 4: sourceoverride.ResolveOverrideResponse
}
var file_sourceoverride_proto_depIdxs = []int32{
	2, // 0: sourceoverride.ProxyResponse.auth:type_name -> sourceoverride.AuthMsg
	4, // 1: sourceoverride.ProxyResponse.response:type_name -> sourceoverride.ResolveOverrideResponse
	2, // 2: sourceoverride.ProxyRequest.auth:type_name -> sourceoverride.AuthMsg
	3, // 3: sourceoverride.ProxyRequest.request:type_name -> sourceoverride.ResolveOverrideRequest
	0, // 4: sourceoverride.Proxy.ProxyStream:input_type -> sourceoverride.ProxyResponse
	1, // 5: sourceoverride.Proxy.ResolveOverride:input_type -> sourceoverride.ProxyRequest
	1, // 6: sourceoverride.Proxy.ProxyStream:output_type -> sourceoverride.ProxyRequest
	4, // 7: sourceoverride.Proxy.ResolveOverride:output_type -> sourceoverride.ResolveOverrideResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sourceoverride_proto_init() }
func file_sourceoverride_proto_init() {
	if File_sourceoverride_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sourceoverride_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyResponse); i {
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
		file_sourceoverride_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRequest); i {
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
		file_sourceoverride_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMsg); i {
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
		file_sourceoverride_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveOverrideRequest); i {
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
		file_sourceoverride_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveOverrideResponse); i {
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
	file_sourceoverride_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_sourceoverride_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_sourceoverride_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_sourceoverride_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sourceoverride_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sourceoverride_proto_goTypes,
		DependencyIndexes: file_sourceoverride_proto_depIdxs,
		MessageInfos:      file_sourceoverride_proto_msgTypes,
	}.Build()
	File_sourceoverride_proto = out.File
	file_sourceoverride_proto_rawDesc = nil
	file_sourceoverride_proto_goTypes = nil
	file_sourceoverride_proto_depIdxs = nil
}

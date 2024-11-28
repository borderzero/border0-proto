// sets the .proto file syntax version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: device.proto

// sets the protobuf package name (i.e. definitions namespace)

package device

import (
	common "github.com/borderzero/border0-proto/common"
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

// messages from devices to the server (api)
type DeviceToServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*DeviceToServerMessage_AuthChallengeSolution
	//	*DeviceToServerMessage_DiscoveryDetails
	//	*DeviceToServerMessage_Heartbeat
	//	*DeviceToServerMessage_Stats
	Message isDeviceToServerMessage_Message `protobuf_oneof:"Message"`
}

func (x *DeviceToServerMessage) Reset() {
	*x = DeviceToServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceToServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceToServerMessage) ProtoMessage() {}

func (x *DeviceToServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceToServerMessage.ProtoReflect.Descriptor instead.
func (*DeviceToServerMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{0}
}

func (m *DeviceToServerMessage) GetMessage() isDeviceToServerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *DeviceToServerMessage) GetAuthChallengeSolution() *AuthChallengeSolutionMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_AuthChallengeSolution); ok {
		return x.AuthChallengeSolution
	}
	return nil
}

func (x *DeviceToServerMessage) GetDiscoveryDetails() *common.DiscoveryDetailsMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_DiscoveryDetails); ok {
		return x.DiscoveryDetails
	}
	return nil
}

func (x *DeviceToServerMessage) GetHeartbeat() *common.HeartbeatMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *DeviceToServerMessage) GetStats() *common.StatsMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_Stats); ok {
		return x.Stats
	}
	return nil
}

type isDeviceToServerMessage_Message interface {
	isDeviceToServerMessage_Message()
}

type DeviceToServerMessage_AuthChallengeSolution struct {
	AuthChallengeSolution *AuthChallengeSolutionMessage `protobuf:"bytes,1,opt,name=auth_challenge_solution,json=authChallengeSolution,proto3,oneof"`
}

type DeviceToServerMessage_DiscoveryDetails struct {
	DiscoveryDetails *common.DiscoveryDetailsMessage `protobuf:"bytes,2,opt,name=discovery_details,json=discoveryDetails,proto3,oneof"`
}

type DeviceToServerMessage_Heartbeat struct {
	Heartbeat *common.HeartbeatMessage `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof"`
}

type DeviceToServerMessage_Stats struct {
	Stats *common.StatsMessage `protobuf:"bytes,4,opt,name=stats,proto3,oneof"`
}

func (*DeviceToServerMessage_AuthChallengeSolution) isDeviceToServerMessage_Message() {}

func (*DeviceToServerMessage_DiscoveryDetails) isDeviceToServerMessage_Message() {}

func (*DeviceToServerMessage_Heartbeat) isDeviceToServerMessage_Message() {}

func (*DeviceToServerMessage_Stats) isDeviceToServerMessage_Message() {}

// messages from the server (api) to devices
type ServerToDeviceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ServerToDeviceMessage_AuthChallenge
	//	*ServerToDeviceMessage_Heartbeat
	//	*ServerToDeviceMessage_NetworkState
	//	*ServerToDeviceMessage_PeerOnline
	//	*ServerToDeviceMessage_PeerOffline
	//	*ServerToDeviceMessage_Disconnect
	//	*ServerToDeviceMessage_Service
	Message isServerToDeviceMessage_Message `protobuf_oneof:"Message"`
}

func (x *ServerToDeviceMessage) Reset() {
	*x = ServerToDeviceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerToDeviceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerToDeviceMessage) ProtoMessage() {}

func (x *ServerToDeviceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerToDeviceMessage.ProtoReflect.Descriptor instead.
func (*ServerToDeviceMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{1}
}

func (m *ServerToDeviceMessage) GetMessage() isServerToDeviceMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ServerToDeviceMessage) GetAuthChallenge() *AuthChallengeMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_AuthChallenge); ok {
		return x.AuthChallenge
	}
	return nil
}

func (x *ServerToDeviceMessage) GetHeartbeat() *common.HeartbeatMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *ServerToDeviceMessage) GetNetworkState() *common.NetworkStateMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_NetworkState); ok {
		return x.NetworkState
	}
	return nil
}

func (x *ServerToDeviceMessage) GetPeerOnline() *common.PeerOnlineMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_PeerOnline); ok {
		return x.PeerOnline
	}
	return nil
}

func (x *ServerToDeviceMessage) GetPeerOffline() *common.PeerOfflineMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_PeerOffline); ok {
		return x.PeerOffline
	}
	return nil
}

func (x *ServerToDeviceMessage) GetDisconnect() *common.DisconnectMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_Disconnect); ok {
		return x.Disconnect
	}
	return nil
}

func (x *ServerToDeviceMessage) GetService() *Service {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_Service); ok {
		return x.Service
	}
	return nil
}

type isServerToDeviceMessage_Message interface {
	isServerToDeviceMessage_Message()
}

type ServerToDeviceMessage_AuthChallenge struct {
	AuthChallenge *AuthChallengeMessage `protobuf:"bytes,1,opt,name=auth_challenge,json=authChallenge,proto3,oneof"`
}

type ServerToDeviceMessage_Heartbeat struct {
	Heartbeat *common.HeartbeatMessage `protobuf:"bytes,2,opt,name=heartbeat,proto3,oneof"`
}

type ServerToDeviceMessage_NetworkState struct {
	NetworkState *common.NetworkStateMessage `protobuf:"bytes,3,opt,name=network_state,json=networkState,proto3,oneof"`
}

type ServerToDeviceMessage_PeerOnline struct {
	PeerOnline *common.PeerOnlineMessage `protobuf:"bytes,4,opt,name=peer_online,json=peerOnline,proto3,oneof"`
}

type ServerToDeviceMessage_PeerOffline struct {
	PeerOffline *common.PeerOfflineMessage `protobuf:"bytes,5,opt,name=peer_offline,json=peerOffline,proto3,oneof"`
}

type ServerToDeviceMessage_Disconnect struct {
	Disconnect *common.DisconnectMessage `protobuf:"bytes,6,opt,name=disconnect,proto3,oneof"`
}

type ServerToDeviceMessage_Service struct {
	Service *Service `protobuf:"bytes,7,opt,name=service,proto3,oneof"`
}

func (*ServerToDeviceMessage_AuthChallenge) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_Heartbeat) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_NetworkState) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_PeerOnline) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_PeerOffline) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_Disconnect) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_Service) isServerToDeviceMessage_Message() {}

type AuthChallengeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerPublicKey string `protobuf:"bytes,1,opt,name=serverPublicKey,proto3" json:"serverPublicKey,omitempty"`
	Challenge       []byte `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
	ChallengeNonce  []byte `protobuf:"bytes,3,opt,name=challengeNonce,proto3" json:"challengeNonce,omitempty"`
}

func (x *AuthChallengeMessage) Reset() {
	*x = AuthChallengeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthChallengeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthChallengeMessage) ProtoMessage() {}

func (x *AuthChallengeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthChallengeMessage.ProtoReflect.Descriptor instead.
func (*AuthChallengeMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{2}
}

func (x *AuthChallengeMessage) GetServerPublicKey() string {
	if x != nil {
		return x.ServerPublicKey
	}
	return ""
}

func (x *AuthChallengeMessage) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

func (x *AuthChallengeMessage) GetChallengeNonce() []byte {
	if x != nil {
		return x.ChallengeNonce
	}
	return nil
}

type AuthChallengeSolutionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solved      []byte `protobuf:"bytes,1,opt,name=solved,proto3" json:"solved,omitempty"`
	SolvedNonce []byte `protobuf:"bytes,2,opt,name=solvedNonce,proto3" json:"solvedNonce,omitempty"`
}

func (x *AuthChallengeSolutionMessage) Reset() {
	*x = AuthChallengeSolutionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthChallengeSolutionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthChallengeSolutionMessage) ProtoMessage() {}

func (x *AuthChallengeSolutionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthChallengeSolutionMessage.ProtoReflect.Descriptor instead.
func (*AuthChallengeSolutionMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{3}
}

func (x *AuthChallengeSolutionMessage) GetSolved() []byte {
	if x != nil {
		return x.Solved
	}
	return nil
}

func (x *AuthChallengeSolutionMessage) GetSolvedNonce() []byte {
	if x != nil {
		return x.SolvedNonce
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId     string   `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Ipv4          string   `protobuf:"bytes,4,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6          string   `protobuf:"bytes,5,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	SubnetRoutes  []string `protobuf:"bytes,6,rep,name=subnet_routes,json=subnetRoutes,proto3" json:"subnet_routes,omitempty"`
	PeerPublicKey []string `protobuf:"bytes,7,rep,name=peer_public_key,json=peerPublicKey,proto3" json:"peer_public_key,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{4}
}

func (x *Service) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Service) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *Service) GetIpv6() string {
	if x != nil {
		return x.Ipv6
	}
	return ""
}

func (x *Service) GetSubnetRoutes() []string {
	if x != nil {
		return x.SubnetRoutes
	}
	return nil
}

func (x *Service) GetPeerPublicKey() []string {
	if x != nil {
		return x.PeerPublicKey
	}
	return nil
}

var File_device_proto protoreflect.FileDescriptor

var file_device_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x15, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x69, 0x0a, 0x17, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x15, 0x61, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a,
	0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x37, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x9d, 0x04, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d,
	0x61, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x43, 0x0a,
	0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x4d, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x47, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0a,
	0x70, 0x65, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x36,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x1c, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x70, 0x76, 0x36, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x65, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x32, 0x82, 0x01, 0x0a,
	0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x2e, 0x62, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x28, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x7a, 0x65, 0x72, 0x6f, 0x2f, 0x62, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x30, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_proto_rawDescOnce sync.Once
	file_device_proto_rawDescData = file_device_proto_rawDesc
)

func file_device_proto_rawDescGZIP() []byte {
	file_device_proto_rawDescOnce.Do(func() {
		file_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_proto_rawDescData)
	})
	return file_device_proto_rawDescData
}

var file_device_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_device_proto_goTypes = []any{
	(*DeviceToServerMessage)(nil),          // 0: border0.device.v1.DeviceToServerMessage
	(*ServerToDeviceMessage)(nil),          // 1: border0.device.v1.ServerToDeviceMessage
	(*AuthChallengeMessage)(nil),           // 2: border0.device.v1.AuthChallengeMessage
	(*AuthChallengeSolutionMessage)(nil),   // 3: border0.device.v1.AuthChallengeSolutionMessage
	(*Service)(nil),                        // 4: border0.device.v1.Service
	(*common.DiscoveryDetailsMessage)(nil), // 5: border0.common.v1.DiscoveryDetailsMessage
	(*common.HeartbeatMessage)(nil),        // 6: border0.common.v1.HeartbeatMessage
	(*common.StatsMessage)(nil),            // 7: border0.common.v1.StatsMessage
	(*common.NetworkStateMessage)(nil),     // 8: border0.common.v1.NetworkStateMessage
	(*common.PeerOnlineMessage)(nil),       // 9: border0.common.v1.PeerOnlineMessage
	(*common.PeerOfflineMessage)(nil),      // 10: border0.common.v1.PeerOfflineMessage
	(*common.DisconnectMessage)(nil),       // 11: border0.common.v1.DisconnectMessage
}
var file_device_proto_depIdxs = []int32{
	3,  // 0: border0.device.v1.DeviceToServerMessage.auth_challenge_solution:type_name -> border0.device.v1.AuthChallengeSolutionMessage
	5,  // 1: border0.device.v1.DeviceToServerMessage.discovery_details:type_name -> border0.common.v1.DiscoveryDetailsMessage
	6,  // 2: border0.device.v1.DeviceToServerMessage.heartbeat:type_name -> border0.common.v1.HeartbeatMessage
	7,  // 3: border0.device.v1.DeviceToServerMessage.stats:type_name -> border0.common.v1.StatsMessage
	2,  // 4: border0.device.v1.ServerToDeviceMessage.auth_challenge:type_name -> border0.device.v1.AuthChallengeMessage
	6,  // 5: border0.device.v1.ServerToDeviceMessage.heartbeat:type_name -> border0.common.v1.HeartbeatMessage
	8,  // 6: border0.device.v1.ServerToDeviceMessage.network_state:type_name -> border0.common.v1.NetworkStateMessage
	9,  // 7: border0.device.v1.ServerToDeviceMessage.peer_online:type_name -> border0.common.v1.PeerOnlineMessage
	10, // 8: border0.device.v1.ServerToDeviceMessage.peer_offline:type_name -> border0.common.v1.PeerOfflineMessage
	11, // 9: border0.device.v1.ServerToDeviceMessage.disconnect:type_name -> border0.common.v1.DisconnectMessage
	4,  // 10: border0.device.v1.ServerToDeviceMessage.service:type_name -> border0.device.v1.Service
	0,  // 11: border0.device.v1.DeviceManagementService.ControlStream:input_type -> border0.device.v1.DeviceToServerMessage
	1,  // 12: border0.device.v1.DeviceManagementService.ControlStream:output_type -> border0.device.v1.ServerToDeviceMessage
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_device_proto_init() }
func file_device_proto_init() {
	if File_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceToServerMessage); i {
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
		file_device_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ServerToDeviceMessage); i {
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
		file_device_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AuthChallengeMessage); i {
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
		file_device_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AuthChallengeSolutionMessage); i {
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
		file_device_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Service); i {
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
	file_device_proto_msgTypes[0].OneofWrappers = []any{
		(*DeviceToServerMessage_AuthChallengeSolution)(nil),
		(*DeviceToServerMessage_DiscoveryDetails)(nil),
		(*DeviceToServerMessage_Heartbeat)(nil),
		(*DeviceToServerMessage_Stats)(nil),
	}
	file_device_proto_msgTypes[1].OneofWrappers = []any{
		(*ServerToDeviceMessage_AuthChallenge)(nil),
		(*ServerToDeviceMessage_Heartbeat)(nil),
		(*ServerToDeviceMessage_NetworkState)(nil),
		(*ServerToDeviceMessage_PeerOnline)(nil),
		(*ServerToDeviceMessage_PeerOffline)(nil),
		(*ServerToDeviceMessage_Disconnect)(nil),
		(*ServerToDeviceMessage_Service)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_proto_goTypes,
		DependencyIndexes: file_device_proto_depIdxs,
		MessageInfos:      file_device_proto_msgTypes,
	}.Build()
	File_device_proto = out.File
	file_device_proto_rawDesc = nil
	file_device_proto_goTypes = nil
	file_device_proto_depIdxs = nil
}

// sets the .proto file syntax version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: device.proto

// sets the protobuf package name (i.e. definitions namespace)

package device

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

// messages from devices to the server (api)
type DeviceToServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*DeviceToServerMessage_AuthChallengeSolutionMessage
	//	*DeviceToServerMessage_DiscoveryDetailsMessage
	//	*DeviceToServerMessage_HeartbeatMessage
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

func (x *DeviceToServerMessage) GetAuthChallengeSolutionMessage() *AuthChallengeSolutionMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_AuthChallengeSolutionMessage); ok {
		return x.AuthChallengeSolutionMessage
	}
	return nil
}

func (x *DeviceToServerMessage) GetDiscoveryDetailsMessage() *DiscoveryDetailsMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_DiscoveryDetailsMessage); ok {
		return x.DiscoveryDetailsMessage
	}
	return nil
}

func (x *DeviceToServerMessage) GetHeartbeatMessage() *HeartbeatMessage {
	if x, ok := x.GetMessage().(*DeviceToServerMessage_HeartbeatMessage); ok {
		return x.HeartbeatMessage
	}
	return nil
}

type isDeviceToServerMessage_Message interface {
	isDeviceToServerMessage_Message()
}

type DeviceToServerMessage_AuthChallengeSolutionMessage struct {
	AuthChallengeSolutionMessage *AuthChallengeSolutionMessage `protobuf:"bytes,1,opt,name=authChallengeSolutionMessage,proto3,oneof"`
}

type DeviceToServerMessage_DiscoveryDetailsMessage struct {
	DiscoveryDetailsMessage *DiscoveryDetailsMessage `protobuf:"bytes,2,opt,name=discoveryDetailsMessage,proto3,oneof"`
}

type DeviceToServerMessage_HeartbeatMessage struct {
	HeartbeatMessage *HeartbeatMessage `protobuf:"bytes,3,opt,name=heartbeatMessage,proto3,oneof"`
}

func (*DeviceToServerMessage_AuthChallengeSolutionMessage) isDeviceToServerMessage_Message() {}

func (*DeviceToServerMessage_DiscoveryDetailsMessage) isDeviceToServerMessage_Message() {}

func (*DeviceToServerMessage_HeartbeatMessage) isDeviceToServerMessage_Message() {}

// messages from the server (api) to devices
type ServerToDeviceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ServerToDeviceMessage_AuthChallengeMessage
	//	*ServerToDeviceMessage_HeartbeatMessage
	//	*ServerToDeviceMessage_NetworkStateMessage
	//	*ServerToDeviceMessage_PeerOnlineMessage
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

func (x *ServerToDeviceMessage) GetAuthChallengeMessage() *AuthChallengeMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_AuthChallengeMessage); ok {
		return x.AuthChallengeMessage
	}
	return nil
}

func (x *ServerToDeviceMessage) GetHeartbeatMessage() *HeartbeatMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_HeartbeatMessage); ok {
		return x.HeartbeatMessage
	}
	return nil
}

func (x *ServerToDeviceMessage) GetNetworkStateMessage() *NetworkStateMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_NetworkStateMessage); ok {
		return x.NetworkStateMessage
	}
	return nil
}

func (x *ServerToDeviceMessage) GetPeerOnlineMessage() *PeerOnlineMessage {
	if x, ok := x.GetMessage().(*ServerToDeviceMessage_PeerOnlineMessage); ok {
		return x.PeerOnlineMessage
	}
	return nil
}

type isServerToDeviceMessage_Message interface {
	isServerToDeviceMessage_Message()
}

type ServerToDeviceMessage_AuthChallengeMessage struct {
	AuthChallengeMessage *AuthChallengeMessage `protobuf:"bytes,1,opt,name=authChallengeMessage,proto3,oneof"`
}

type ServerToDeviceMessage_HeartbeatMessage struct {
	HeartbeatMessage *HeartbeatMessage `protobuf:"bytes,2,opt,name=heartbeatMessage,proto3,oneof"`
}

type ServerToDeviceMessage_NetworkStateMessage struct {
	NetworkStateMessage *NetworkStateMessage `protobuf:"bytes,3,opt,name=networkStateMessage,proto3,oneof"`
}

type ServerToDeviceMessage_PeerOnlineMessage struct {
	PeerOnlineMessage *PeerOnlineMessage `protobuf:"bytes,4,opt,name=peerOnlineMessage,proto3,oneof"`
}

func (*ServerToDeviceMessage_AuthChallengeMessage) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_HeartbeatMessage) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_NetworkStateMessage) isServerToDeviceMessage_Message() {}

func (*ServerToDeviceMessage_PeerOnlineMessage) isServerToDeviceMessage_Message() {}

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

type DiscoveryDetailsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discoverable bool   `protobuf:"varint,1,opt,name=discoverable,proto3" json:"discoverable,omitempty"`                    // whether the peer should be returned along with the networks its in
	EndpointUdp4 string `protobuf:"bytes,2,opt,name=endpoint_udp4,json=endpointUdp4,proto3" json:"endpoint_udp4,omitempty"` // the endpoint (IPv4 + port) to send packets to for the peer
}

func (x *DiscoveryDetailsMessage) Reset() {
	*x = DiscoveryDetailsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryDetailsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryDetailsMessage) ProtoMessage() {}

func (x *DiscoveryDetailsMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DiscoveryDetailsMessage.ProtoReflect.Descriptor instead.
func (*DiscoveryDetailsMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{4}
}

func (x *DiscoveryDetailsMessage) GetDiscoverable() bool {
	if x != nil {
		return x.Discoverable
	}
	return false
}

func (x *DiscoveryDetailsMessage) GetEndpointUdp4() string {
	if x != nil {
		return x.EndpointUdp4
	}
	return ""
}

type HeartbeatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatMessage) Reset() {
	*x = HeartbeatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatMessage) ProtoMessage() {}

func (x *HeartbeatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatMessage.ProtoReflect.Descriptor instead.
func (*HeartbeatMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{5}
}

type PeerOnlineMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkID string         `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Peer      *WireGuardPeer `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *PeerOnlineMessage) Reset() {
	*x = PeerOnlineMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerOnlineMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerOnlineMessage) ProtoMessage() {}

func (x *PeerOnlineMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerOnlineMessage.ProtoReflect.Descriptor instead.
func (*PeerOnlineMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{6}
}

func (x *PeerOnlineMessage) GetNetworkID() string {
	if x != nil {
		return x.NetworkID
	}
	return ""
}

func (x *PeerOnlineMessage) GetPeer() *WireGuardPeer {
	if x != nil {
		return x.Peer
	}
	return nil
}

type NetworkStateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkID     string           `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	NetworkCIDRV4 string           `protobuf:"bytes,2,opt,name=networkCIDRV4,proto3" json:"networkCIDRV4,omitempty"`
	NetworkCIDRV6 string           `protobuf:"bytes,3,opt,name=networkCIDRV6,proto3" json:"networkCIDRV6,omitempty"`
	SelfIPv4      string           `protobuf:"bytes,4,opt,name=selfIPv4,proto3" json:"selfIPv4,omitempty"`
	SelfIPv6      string           `protobuf:"bytes,5,opt,name=selfIPv6,proto3" json:"selfIPv6,omitempty"`
	OnlinePeers   []*WireGuardPeer `protobuf:"bytes,6,rep,name=onlinePeers,proto3" json:"onlinePeers,omitempty"`
}

func (x *NetworkStateMessage) Reset() {
	*x = NetworkStateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkStateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkStateMessage) ProtoMessage() {}

func (x *NetworkStateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkStateMessage.ProtoReflect.Descriptor instead.
func (*NetworkStateMessage) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{7}
}

func (x *NetworkStateMessage) GetNetworkID() string {
	if x != nil {
		return x.NetworkID
	}
	return ""
}

func (x *NetworkStateMessage) GetNetworkCIDRV4() string {
	if x != nil {
		return x.NetworkCIDRV4
	}
	return ""
}

func (x *NetworkStateMessage) GetNetworkCIDRV6() string {
	if x != nil {
		return x.NetworkCIDRV6
	}
	return ""
}

func (x *NetworkStateMessage) GetSelfIPv4() string {
	if x != nil {
		return x.SelfIPv4
	}
	return ""
}

func (x *NetworkStateMessage) GetSelfIPv6() string {
	if x != nil {
		return x.SelfIPv6
	}
	return ""
}

func (x *NetworkStateMessage) GetOnlinePeers() []*WireGuardPeer {
	if x != nil {
		return x.OnlinePeers
	}
	return nil
}

type WireGuardPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey                          string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`                                                                                 // the public key of the peer, used for identification and encryption
	IPv4                               string   `protobuf:"bytes,2,opt,name=IPv4,proto3" json:"IPv4,omitempty"`                                                                                                            // the peer's (private) IPv4 address in the wireguard network
	IPv6                               string   `protobuf:"bytes,3,opt,name=IPv6,proto3" json:"IPv6,omitempty"`                                                                                                            // the peer's (private) IPv6 address in the wireguard network
	AllowedIps                         []string `protobuf:"bytes,4,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`                                                                              // list of routes (CIDRs) to be routed through this peer (most peers will just have their own IP/32)
	Endpoints                          []string `protobuf:"bytes,5,rep,name=endpoints,proto3" json:"endpoints,omitempty"`                                                                                                  // the endpoints to try send packets to for the peer e.g. updv4, udpv6, relayv4, relayv6, etc...
	PersistentKeepaliveIntervalSeconds uint32   `protobuf:"varint,6,opt,name=persistent_keepalive_interval_seconds,json=persistentKeepaliveIntervalSeconds,proto3" json:"persistent_keepalive_interval_seconds,omitempty"` // the interval for sending keepalive packets (0 means disabled)
}

func (x *WireGuardPeer) Reset() {
	*x = WireGuardPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireGuardPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireGuardPeer) ProtoMessage() {}

func (x *WireGuardPeer) ProtoReflect() protoreflect.Message {
	mi := &file_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireGuardPeer.ProtoReflect.Descriptor instead.
func (*WireGuardPeer) Descriptor() ([]byte, []int) {
	return file_device_proto_rawDescGZIP(), []int{8}
}

func (x *WireGuardPeer) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *WireGuardPeer) GetIPv4() string {
	if x != nil {
		return x.IPv4
	}
	return ""
}

func (x *WireGuardPeer) GetIPv6() string {
	if x != nil {
		return x.IPv6
	}
	return ""
}

func (x *WireGuardPeer) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

func (x *WireGuardPeer) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *WireGuardPeer) GetPersistentKeepaliveIntervalSeconds() uint32 {
	if x != nil {
		return x.PersistentKeepaliveIntervalSeconds
	}
	return 0
}

var File_device_proto protoreflect.FileDescriptor

var file_device_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0xd4, 0x02, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x75, 0x0a, 0x1c, 0x61,
	0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x1c, 0x61, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x66, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x17, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x10, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x86, 0x03, 0x0a, 0x15, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x14, 0x61, 0x75, 0x74,
	0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x51, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x10, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x13, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x54, 0x0a, 0x11, 0x70, 0x65, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x11, 0x70, 0x65, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
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
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0x62, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x75, 0x64, 0x70, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x55, 0x64, 0x70, 0x34, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x67, 0x0a, 0x11,
	0x50, 0x65, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12,
	0x34, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x47, 0x75, 0x61, 0x72, 0x64, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0xfb, 0x01, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52, 0x56, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52, 0x56,
	0x34, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52,
	0x56, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x49, 0x44, 0x52, 0x56, 0x36, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x49,
	0x50, 0x76, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x49,
	0x50, 0x76, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x50, 0x76, 0x36, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x50, 0x76, 0x36, 0x12,
	0x42, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x47, 0x75, 0x61,
	0x72, 0x64, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x0d, 0x57, 0x69, 0x72, 0x65, 0x47, 0x75, 0x61, 0x72,
	0x64, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x49, 0x50, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x36,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x50, 0x76, 0x36, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x51, 0x0a, 0x25, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x22, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x32, 0x82,
	0x01, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x2e, 0x62, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x28, 0x2e, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x7a, 0x65, 0x72, 0x6f, 0x2f, 0x62, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x30, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_device_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_device_proto_goTypes = []any{
	(*DeviceToServerMessage)(nil),        // 0: border0.device.v1.DeviceToServerMessage
	(*ServerToDeviceMessage)(nil),        // 1: border0.device.v1.ServerToDeviceMessage
	(*AuthChallengeMessage)(nil),         // 2: border0.device.v1.AuthChallengeMessage
	(*AuthChallengeSolutionMessage)(nil), // 3: border0.device.v1.AuthChallengeSolutionMessage
	(*DiscoveryDetailsMessage)(nil),      // 4: border0.device.v1.DiscoveryDetailsMessage
	(*HeartbeatMessage)(nil),             // 5: border0.device.v1.HeartbeatMessage
	(*PeerOnlineMessage)(nil),            // 6: border0.device.v1.PeerOnlineMessage
	(*NetworkStateMessage)(nil),          // 7: border0.device.v1.NetworkStateMessage
	(*WireGuardPeer)(nil),                // 8: border0.device.v1.WireGuardPeer
}
var file_device_proto_depIdxs = []int32{
	3,  // 0: border0.device.v1.DeviceToServerMessage.authChallengeSolutionMessage:type_name -> border0.device.v1.AuthChallengeSolutionMessage
	4,  // 1: border0.device.v1.DeviceToServerMessage.discoveryDetailsMessage:type_name -> border0.device.v1.DiscoveryDetailsMessage
	5,  // 2: border0.device.v1.DeviceToServerMessage.heartbeatMessage:type_name -> border0.device.v1.HeartbeatMessage
	2,  // 3: border0.device.v1.ServerToDeviceMessage.authChallengeMessage:type_name -> border0.device.v1.AuthChallengeMessage
	5,  // 4: border0.device.v1.ServerToDeviceMessage.heartbeatMessage:type_name -> border0.device.v1.HeartbeatMessage
	7,  // 5: border0.device.v1.ServerToDeviceMessage.networkStateMessage:type_name -> border0.device.v1.NetworkStateMessage
	6,  // 6: border0.device.v1.ServerToDeviceMessage.peerOnlineMessage:type_name -> border0.device.v1.PeerOnlineMessage
	8,  // 7: border0.device.v1.PeerOnlineMessage.peer:type_name -> border0.device.v1.WireGuardPeer
	8,  // 8: border0.device.v1.NetworkStateMessage.onlinePeers:type_name -> border0.device.v1.WireGuardPeer
	0,  // 9: border0.device.v1.DeviceManagementService.ControlStream:input_type -> border0.device.v1.DeviceToServerMessage
	1,  // 10: border0.device.v1.DeviceManagementService.ControlStream:output_type -> border0.device.v1.ServerToDeviceMessage
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
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
			switch v := v.(*DiscoveryDetailsMessage); i {
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
		file_device_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*HeartbeatMessage); i {
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
		file_device_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PeerOnlineMessage); i {
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
		file_device_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkStateMessage); i {
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
		file_device_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*WireGuardPeer); i {
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
		(*DeviceToServerMessage_AuthChallengeSolutionMessage)(nil),
		(*DeviceToServerMessage_DiscoveryDetailsMessage)(nil),
		(*DeviceToServerMessage_HeartbeatMessage)(nil),
	}
	file_device_proto_msgTypes[1].OneofWrappers = []any{
		(*ServerToDeviceMessage_AuthChallengeMessage)(nil),
		(*ServerToDeviceMessage_HeartbeatMessage)(nil),
		(*ServerToDeviceMessage_NetworkStateMessage)(nil),
		(*ServerToDeviceMessage_PeerOnlineMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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

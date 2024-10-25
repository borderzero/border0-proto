// DO NOT EDIT.
// swift-format-ignore-file
// swiftlint:disable all
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: messages.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

/// sets the .proto file syntax version

import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

enum Border0_Common_V1_PeerType: SwiftProtobuf.Enum, Swift.CaseIterable {
  typealias RawValue = Int
  case unknown // = 0
  case device // = 1
  case connector // = 2
  case UNRECOGNIZED(Int)

  init() {
    self = .unknown
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unknown
    case 1: self = .device
    case 2: self = .connector
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .unknown: return 0
    case .device: return 1
    case .connector: return 2
    case .UNRECOGNIZED(let i): return i
    }
  }

  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static let allCases: [Border0_Common_V1_PeerType] = [
    .unknown,
    .device,
    .connector,
  ]

}

enum Border0_Common_V1_ServiceType: SwiftProtobuf.Enum, Swift.CaseIterable {
  typealias RawValue = Int
  case unknown // = 0
  case http // = 1
  case ssh // = 2
  case database // = 3
  case tls // = 4
  case vnc // = 5
  case rdp // = 6
  case kubernetes // = 7
  case subnetRoutes // = 8
  case UNRECOGNIZED(Int)

  init() {
    self = .unknown
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unknown
    case 1: self = .http
    case 2: self = .ssh
    case 3: self = .database
    case 4: self = .tls
    case 5: self = .vnc
    case 6: self = .rdp
    case 7: self = .kubernetes
    case 8: self = .subnetRoutes
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .unknown: return 0
    case .http: return 1
    case .ssh: return 2
    case .database: return 3
    case .tls: return 4
    case .vnc: return 5
    case .rdp: return 6
    case .kubernetes: return 7
    case .subnetRoutes: return 8
    case .UNRECOGNIZED(let i): return i
    }
  }

  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static let allCases: [Border0_Common_V1_ServiceType] = [
    .unknown,
    .http,
    .ssh,
    .database,
    .tls,
    .vnc,
    .rdp,
    .kubernetes,
    .subnetRoutes,
  ]

}

enum Border0_Common_V1_DisconnectionReason: SwiftProtobuf.Enum, Swift.CaseIterable {
  typealias RawValue = Int
  case unknown // = 0
  case serverShutdown // = 1
  case UNRECOGNIZED(Int)

  init() {
    self = .unknown
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .unknown
    case 1: self = .serverShutdown
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .unknown: return 0
    case .serverShutdown: return 1
    case .UNRECOGNIZED(let i): return i
    }
  }

  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static let allCases: [Border0_Common_V1_DisconnectionReason] = [
    .unknown,
    .serverShutdown,
  ]

}

struct Border0_Common_V1_DiscoveryDetailsMessage: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// whether the peer should be returned along with the networks it's in
  var discoverable: Bool = false

  /// the endpoint (IPv4 + port) to send packets to for the peer
  var endpointPublicUdp4: String = String()

  /// the endpoint (IPv6 + port) to send packets to for the peer
  var endpointPublicUdp6: String = String()

  /// the public key of the peer, only used in connector
  var publicKey: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_HeartbeatMessage: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_PeerOnlineMessage: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var networkID: String = String()

  var peer: Border0_Common_V1_WireGuardPeer {
    get {return _peer ?? Border0_Common_V1_WireGuardPeer()}
    set {_peer = newValue}
  }
  /// Returns true if `peer` has been explicitly set.
  var hasPeer: Bool {return self._peer != nil}
  /// Clears the value of `peer`. Subsequent reads from it will return its default value.
  mutating func clearPeer() {self._peer = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _peer: Border0_Common_V1_WireGuardPeer? = nil
}

struct Border0_Common_V1_PeerOfflineMessage: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var networkID: String = String()

  var peerPublicKey: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_NetworkStateMessage: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var networkID: String = String()

  var networkCidrV4: String = String()

  var networkCidrV6: String = String()

  var selfIpv4: String = String()

  var selfIpv6: String = String()

  var onlinePeers: [Border0_Common_V1_WireGuardPeer] = []

  var networkResourcesCidrV4: String = String()

  var networkResourcesCidrV6: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_WireGuardPeer: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  /// the public key of the peer, used for identification and encryption
  var publicKey: String = String()

  /// the peer's (private) IPv4 address in the WireGuard network
  var ipv4: String = String()

  /// the peer's (private) IPv6 address in the WireGuard network
  var ipv6: String = String()

  /// list of routes (CIDRs) to be routed through this peer (most peers will just have their own IP/32). (this field is now deprecated in favor of building the allowed_ips list from ipv4 + ipv6 + service ips + subnet routes)
  ///
  /// NOTE: This field was marked as deprecated in the .proto file.
  var allowedIps: [String] = []

  /// the interval for sending keepalive packets (0 means disabled)
  var persistentKeepaliveIntervalSeconds: UInt32 = 0

  /// endpoint for UDP peer-to-peer communication over IPv4 (public IPv4 + port as seen from the Internet)
  var publicUdp4Endpoint: String = String()

  /// endpoint for UDP peer-to-peer communication over IPv6 (public IPv6 + port as seen from the Internet)
  var publicUdp6Endpoint: String = String()

  /// client or connector
  var type: Border0_Common_V1_PeerType = .unknown

  /// details about the wireguard peer's underlying device
  var deviceDetails: Border0_Common_V1_DeviceDetails {
    get {return _deviceDetails ?? Border0_Common_V1_DeviceDetails()}
    set {_deviceDetails = newValue}
  }
  /// Returns true if `deviceDetails` has been explicitly set.
  var hasDeviceDetails: Bool {return self._deviceDetails != nil}
  /// Clears the value of `deviceDetails`. Subsequent reads from it will return its default value.
  mutating func clearDeviceDetails() {self._deviceDetails = nil}

  /// applicable only when PeerType == PEER_TYPE_CONNECTOR
  var connectorDetails: Border0_Common_V1_ConnectorDetails {
    get {return _connectorDetails ?? Border0_Common_V1_ConnectorDetails()}
    set {_connectorDetails = newValue}
  }
  /// Returns true if `connectorDetails` has been explicitly set.
  var hasConnectorDetails: Bool {return self._connectorDetails != nil}
  /// Clears the value of `connectorDetails`. Subsequent reads from it will return its default value.
  mutating func clearConnectorDetails() {self._connectorDetails = nil}

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}

  fileprivate var _deviceDetails: Border0_Common_V1_DeviceDetails? = nil
  fileprivate var _connectorDetails: Border0_Common_V1_ConnectorDetails? = nil
}

struct Border0_Common_V1_DeviceDetails: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var name: String = String()

  var uuid: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_ConnectorDetails: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var name: String = String()

  var uuid: String = String()

  var services: [Border0_Common_V1_Service] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_Service: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var name: String = String()

  var type: Border0_Common_V1_ServiceType = .unknown

  var ipv4: String = String()

  var ipv6: String = String()

  /// applicable only to services of type SUBNET_ROUTES
  var subnetRoutes: [String] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Border0_Common_V1_DisconnectMessage: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var reason: Border0_Common_V1_DisconnectionReason = .unknown

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "border0.common.v1"

extension Border0_Common_V1_PeerType: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "PEER_TYPE_UNKNOWN"),
    1: .same(proto: "PEER_TYPE_DEVICE"),
    2: .same(proto: "PEER_TYPE_CONNECTOR"),
  ]
}

extension Border0_Common_V1_ServiceType: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "SERVICE_TYPE_UNKNOWN"),
    1: .same(proto: "SERVICE_TYPE_HTTP"),
    2: .same(proto: "SERVICE_TYPE_SSH"),
    3: .same(proto: "SERVICE_TYPE_DATABASE"),
    4: .same(proto: "SERVICE_TYPE_TLS"),
    5: .same(proto: "SERVICE_TYPE_VNC"),
    6: .same(proto: "SERVICE_TYPE_RDP"),
    7: .same(proto: "SERVICE_TYPE_KUBERNETES"),
    8: .same(proto: "SERVICE_TYPE_SUBNET_ROUTES"),
  ]
}

extension Border0_Common_V1_DisconnectionReason: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "UNKNOWN"),
    1: .same(proto: "SERVER_SHUTDOWN"),
  ]
}

extension Border0_Common_V1_DiscoveryDetailsMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".DiscoveryDetailsMessage"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "discoverable"),
    2: .standard(proto: "endpoint_public_udp4"),
    3: .standard(proto: "endpoint_public_udp6"),
    4: .standard(proto: "public_key"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularBoolField(value: &self.discoverable) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.endpointPublicUdp4) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.endpointPublicUdp6) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.publicKey) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.discoverable != false {
      try visitor.visitSingularBoolField(value: self.discoverable, fieldNumber: 1)
    }
    if !self.endpointPublicUdp4.isEmpty {
      try visitor.visitSingularStringField(value: self.endpointPublicUdp4, fieldNumber: 2)
    }
    if !self.endpointPublicUdp6.isEmpty {
      try visitor.visitSingularStringField(value: self.endpointPublicUdp6, fieldNumber: 3)
    }
    if !self.publicKey.isEmpty {
      try visitor.visitSingularStringField(value: self.publicKey, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_DiscoveryDetailsMessage, rhs: Border0_Common_V1_DiscoveryDetailsMessage) -> Bool {
    if lhs.discoverable != rhs.discoverable {return false}
    if lhs.endpointPublicUdp4 != rhs.endpointPublicUdp4 {return false}
    if lhs.endpointPublicUdp6 != rhs.endpointPublicUdp6 {return false}
    if lhs.publicKey != rhs.publicKey {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_HeartbeatMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".HeartbeatMessage"
  static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    // Load everything into unknown fields
    while try decoder.nextFieldNumber() != nil {}
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_HeartbeatMessage, rhs: Border0_Common_V1_HeartbeatMessage) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_PeerOnlineMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".PeerOnlineMessage"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "network_id"),
    2: .same(proto: "peer"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.networkID) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._peer) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.networkID.isEmpty {
      try visitor.visitSingularStringField(value: self.networkID, fieldNumber: 1)
    }
    try { if let v = self._peer {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_PeerOnlineMessage, rhs: Border0_Common_V1_PeerOnlineMessage) -> Bool {
    if lhs.networkID != rhs.networkID {return false}
    if lhs._peer != rhs._peer {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_PeerOfflineMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".PeerOfflineMessage"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "network_id"),
    2: .standard(proto: "peer_public_key"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.networkID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.peerPublicKey) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.networkID.isEmpty {
      try visitor.visitSingularStringField(value: self.networkID, fieldNumber: 1)
    }
    if !self.peerPublicKey.isEmpty {
      try visitor.visitSingularStringField(value: self.peerPublicKey, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_PeerOfflineMessage, rhs: Border0_Common_V1_PeerOfflineMessage) -> Bool {
    if lhs.networkID != rhs.networkID {return false}
    if lhs.peerPublicKey != rhs.peerPublicKey {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_NetworkStateMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".NetworkStateMessage"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "network_id"),
    2: .standard(proto: "network_cidr_v4"),
    3: .standard(proto: "network_cidr_v6"),
    4: .standard(proto: "self_ipv4"),
    5: .standard(proto: "self_ipv6"),
    6: .standard(proto: "online_peers"),
    7: .standard(proto: "network_resources_cidr_v4"),
    8: .standard(proto: "network_resources_cidr_v6"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.networkID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.networkCidrV4) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.networkCidrV6) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.selfIpv4) }()
      case 5: try { try decoder.decodeSingularStringField(value: &self.selfIpv6) }()
      case 6: try { try decoder.decodeRepeatedMessageField(value: &self.onlinePeers) }()
      case 7: try { try decoder.decodeSingularStringField(value: &self.networkResourcesCidrV4) }()
      case 8: try { try decoder.decodeSingularStringField(value: &self.networkResourcesCidrV6) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.networkID.isEmpty {
      try visitor.visitSingularStringField(value: self.networkID, fieldNumber: 1)
    }
    if !self.networkCidrV4.isEmpty {
      try visitor.visitSingularStringField(value: self.networkCidrV4, fieldNumber: 2)
    }
    if !self.networkCidrV6.isEmpty {
      try visitor.visitSingularStringField(value: self.networkCidrV6, fieldNumber: 3)
    }
    if !self.selfIpv4.isEmpty {
      try visitor.visitSingularStringField(value: self.selfIpv4, fieldNumber: 4)
    }
    if !self.selfIpv6.isEmpty {
      try visitor.visitSingularStringField(value: self.selfIpv6, fieldNumber: 5)
    }
    if !self.onlinePeers.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.onlinePeers, fieldNumber: 6)
    }
    if !self.networkResourcesCidrV4.isEmpty {
      try visitor.visitSingularStringField(value: self.networkResourcesCidrV4, fieldNumber: 7)
    }
    if !self.networkResourcesCidrV6.isEmpty {
      try visitor.visitSingularStringField(value: self.networkResourcesCidrV6, fieldNumber: 8)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_NetworkStateMessage, rhs: Border0_Common_V1_NetworkStateMessage) -> Bool {
    if lhs.networkID != rhs.networkID {return false}
    if lhs.networkCidrV4 != rhs.networkCidrV4 {return false}
    if lhs.networkCidrV6 != rhs.networkCidrV6 {return false}
    if lhs.selfIpv4 != rhs.selfIpv4 {return false}
    if lhs.selfIpv6 != rhs.selfIpv6 {return false}
    if lhs.onlinePeers != rhs.onlinePeers {return false}
    if lhs.networkResourcesCidrV4 != rhs.networkResourcesCidrV4 {return false}
    if lhs.networkResourcesCidrV6 != rhs.networkResourcesCidrV6 {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_WireGuardPeer: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".WireGuardPeer"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "public_key"),
    2: .same(proto: "ipv4"),
    3: .same(proto: "ipv6"),
    4: .standard(proto: "allowed_ips"),
    5: .standard(proto: "persistent_keepalive_interval_seconds"),
    6: .standard(proto: "public_udp4_endpoint"),
    7: .standard(proto: "public_udp6_endpoint"),
    8: .same(proto: "type"),
    9: .standard(proto: "device_details"),
    10: .standard(proto: "connector_details"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.publicKey) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.ipv4) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.ipv6) }()
      case 4: try { try decoder.decodeRepeatedStringField(value: &self.allowedIps) }()
      case 5: try { try decoder.decodeSingularUInt32Field(value: &self.persistentKeepaliveIntervalSeconds) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.publicUdp4Endpoint) }()
      case 7: try { try decoder.decodeSingularStringField(value: &self.publicUdp6Endpoint) }()
      case 8: try { try decoder.decodeSingularEnumField(value: &self.type) }()
      case 9: try { try decoder.decodeSingularMessageField(value: &self._deviceDetails) }()
      case 10: try { try decoder.decodeSingularMessageField(value: &self._connectorDetails) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.publicKey.isEmpty {
      try visitor.visitSingularStringField(value: self.publicKey, fieldNumber: 1)
    }
    if !self.ipv4.isEmpty {
      try visitor.visitSingularStringField(value: self.ipv4, fieldNumber: 2)
    }
    if !self.ipv6.isEmpty {
      try visitor.visitSingularStringField(value: self.ipv6, fieldNumber: 3)
    }
    if !self.allowedIps.isEmpty {
      try visitor.visitRepeatedStringField(value: self.allowedIps, fieldNumber: 4)
    }
    if self.persistentKeepaliveIntervalSeconds != 0 {
      try visitor.visitSingularUInt32Field(value: self.persistentKeepaliveIntervalSeconds, fieldNumber: 5)
    }
    if !self.publicUdp4Endpoint.isEmpty {
      try visitor.visitSingularStringField(value: self.publicUdp4Endpoint, fieldNumber: 6)
    }
    if !self.publicUdp6Endpoint.isEmpty {
      try visitor.visitSingularStringField(value: self.publicUdp6Endpoint, fieldNumber: 7)
    }
    if self.type != .unknown {
      try visitor.visitSingularEnumField(value: self.type, fieldNumber: 8)
    }
    try { if let v = self._deviceDetails {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 9)
    } }()
    try { if let v = self._connectorDetails {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 10)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_WireGuardPeer, rhs: Border0_Common_V1_WireGuardPeer) -> Bool {
    if lhs.publicKey != rhs.publicKey {return false}
    if lhs.ipv4 != rhs.ipv4 {return false}
    if lhs.ipv6 != rhs.ipv6 {return false}
    if lhs.allowedIps != rhs.allowedIps {return false}
    if lhs.persistentKeepaliveIntervalSeconds != rhs.persistentKeepaliveIntervalSeconds {return false}
    if lhs.publicUdp4Endpoint != rhs.publicUdp4Endpoint {return false}
    if lhs.publicUdp6Endpoint != rhs.publicUdp6Endpoint {return false}
    if lhs.type != rhs.type {return false}
    if lhs._deviceDetails != rhs._deviceDetails {return false}
    if lhs._connectorDetails != rhs._connectorDetails {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_DeviceDetails: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".DeviceDetails"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "uuid"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.uuid) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if !self.uuid.isEmpty {
      try visitor.visitSingularStringField(value: self.uuid, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_DeviceDetails, rhs: Border0_Common_V1_DeviceDetails) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.uuid != rhs.uuid {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_ConnectorDetails: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".ConnectorDetails"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "uuid"),
    3: .same(proto: "services"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.uuid) }()
      case 3: try { try decoder.decodeRepeatedMessageField(value: &self.services) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if !self.uuid.isEmpty {
      try visitor.visitSingularStringField(value: self.uuid, fieldNumber: 2)
    }
    if !self.services.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.services, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_ConnectorDetails, rhs: Border0_Common_V1_ConnectorDetails) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.uuid != rhs.uuid {return false}
    if lhs.services != rhs.services {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_Service: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Service"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
    2: .same(proto: "type"),
    3: .same(proto: "ipv4"),
    4: .same(proto: "ipv6"),
    5: .standard(proto: "subnet_routes"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 2: try { try decoder.decodeSingularEnumField(value: &self.type) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.ipv4) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.ipv6) }()
      case 5: try { try decoder.decodeRepeatedStringField(value: &self.subnetRoutes) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    if self.type != .unknown {
      try visitor.visitSingularEnumField(value: self.type, fieldNumber: 2)
    }
    if !self.ipv4.isEmpty {
      try visitor.visitSingularStringField(value: self.ipv4, fieldNumber: 3)
    }
    if !self.ipv6.isEmpty {
      try visitor.visitSingularStringField(value: self.ipv6, fieldNumber: 4)
    }
    if !self.subnetRoutes.isEmpty {
      try visitor.visitRepeatedStringField(value: self.subnetRoutes, fieldNumber: 5)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_Service, rhs: Border0_Common_V1_Service) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.type != rhs.type {return false}
    if lhs.ipv4 != rhs.ipv4 {return false}
    if lhs.ipv6 != rhs.ipv6 {return false}
    if lhs.subnetRoutes != rhs.subnetRoutes {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Border0_Common_V1_DisconnectMessage: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".DisconnectMessage"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "reason"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.reason) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.reason != .unknown {
      try visitor.visitSingularEnumField(value: self.reason, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Border0_Common_V1_DisconnectMessage, rhs: Border0_Common_V1_DisconnectMessage) -> Bool {
    if lhs.reason != rhs.reason {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

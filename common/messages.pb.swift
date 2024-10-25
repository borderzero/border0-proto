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

  /// list of routes (CIDRs) to be routed through this peer (most peers will just have their own IP/32)
  var allowedIps: [String] = []

  /// the interval for sending keepalive packets (0 means disabled)
  var persistentKeepaliveIntervalSeconds: UInt32 = 0

  /// endpoint for UDP peer-to-peer communication over IPv4 (public IPv4 + port as seen from the Internet)
  var publicUdp4Endpoint: String = String()

  /// endpoint for UDP peer-to-peer communication over IPv6 (public IPv6 + port as seen from the Internet)
  var publicUdp6Endpoint: String = String()

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
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
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

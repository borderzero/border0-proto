//
// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the protocol buffer compiler.
// Source: device.proto
//
import GRPC
import NIO
import NIOConcurrencyHelpers
import SwiftProtobuf


/// the stream (rpc) service for the Border0 api to manage devices
///
/// Usage: instantiate `Border0_Device_V1_DeviceManagementServiceClient`, then call methods of this protocol to make API calls.
internal protocol Border0_Device_V1_DeviceManagementServiceClientProtocol: GRPCClient {
  var serviceName: String { get }
  var interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? { get }

  func controlStream(
    callOptions: CallOptions?,
    handler: @escaping (Border0_Device_V1_ServerToDeviceMessage) -> Void
  ) -> BidirectionalStreamingCall<Border0_Device_V1_DeviceToServerMessage, Border0_Device_V1_ServerToDeviceMessage>
}

extension Border0_Device_V1_DeviceManagementServiceClientProtocol {
  internal var serviceName: String {
    return "border0.device.v1.DeviceManagementService"
  }

  /// Bidirectional streaming call to ControlStream
  ///
  /// Callers should use the `send` method on the returned object to send messages
  /// to the server. The caller should send an `.end` after the final message has been sent.
  ///
  /// - Parameters:
  ///   - callOptions: Call options.
  ///   - handler: A closure called when each response is received from the server.
  /// - Returns: A `ClientStreamingCall` with futures for the metadata and status.
  internal func controlStream(
    callOptions: CallOptions? = nil,
    handler: @escaping (Border0_Device_V1_ServerToDeviceMessage) -> Void
  ) -> BidirectionalStreamingCall<Border0_Device_V1_DeviceToServerMessage, Border0_Device_V1_ServerToDeviceMessage> {
    return self.makeBidirectionalStreamingCall(
      path: Border0_Device_V1_DeviceManagementServiceClientMetadata.Methods.controlStream.path,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeControlStreamInterceptors() ?? [],
      handler: handler
    )
  }
}

@available(*, deprecated)
extension Border0_Device_V1_DeviceManagementServiceClient: @unchecked Sendable {}

@available(*, deprecated, renamed: "Border0_Device_V1_DeviceManagementServiceNIOClient")
internal final class Border0_Device_V1_DeviceManagementServiceClient: Border0_Device_V1_DeviceManagementServiceClientProtocol {
  private let lock = Lock()
  private var _defaultCallOptions: CallOptions
  private var _interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol?
  internal let channel: GRPCChannel
  internal var defaultCallOptions: CallOptions {
    get { self.lock.withLock { return self._defaultCallOptions } }
    set { self.lock.withLockVoid { self._defaultCallOptions = newValue } }
  }
  internal var interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? {
    get { self.lock.withLock { return self._interceptors } }
    set { self.lock.withLockVoid { self._interceptors = newValue } }
  }

  /// Creates a client for the border0.device.v1.DeviceManagementService service.
  ///
  /// - Parameters:
  ///   - channel: `GRPCChannel` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  ///   - interceptors: A factory providing interceptors for each RPC.
  internal init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self._defaultCallOptions = defaultCallOptions
    self._interceptors = interceptors
  }
}

internal struct Border0_Device_V1_DeviceManagementServiceNIOClient: Border0_Device_V1_DeviceManagementServiceClientProtocol {
  internal var channel: GRPCChannel
  internal var defaultCallOptions: CallOptions
  internal var interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol?

  /// Creates a client for the border0.device.v1.DeviceManagementService service.
  ///
  /// - Parameters:
  ///   - channel: `GRPCChannel` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  ///   - interceptors: A factory providing interceptors for each RPC.
  internal init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self.defaultCallOptions = defaultCallOptions
    self.interceptors = interceptors
  }
}

/// the stream (rpc) service for the Border0 api to manage devices
@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
internal protocol Border0_Device_V1_DeviceManagementServiceAsyncClientProtocol: GRPCClient {
  static var serviceDescriptor: GRPCServiceDescriptor { get }
  var interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? { get }

  func makeControlStreamCall(
    callOptions: CallOptions?
  ) -> GRPCAsyncBidirectionalStreamingCall<Border0_Device_V1_DeviceToServerMessage, Border0_Device_V1_ServerToDeviceMessage>
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
extension Border0_Device_V1_DeviceManagementServiceAsyncClientProtocol {
  internal static var serviceDescriptor: GRPCServiceDescriptor {
    return Border0_Device_V1_DeviceManagementServiceClientMetadata.serviceDescriptor
  }

  internal var interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? {
    return nil
  }

  internal func makeControlStreamCall(
    callOptions: CallOptions? = nil
  ) -> GRPCAsyncBidirectionalStreamingCall<Border0_Device_V1_DeviceToServerMessage, Border0_Device_V1_ServerToDeviceMessage> {
    return self.makeAsyncBidirectionalStreamingCall(
      path: Border0_Device_V1_DeviceManagementServiceClientMetadata.Methods.controlStream.path,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeControlStreamInterceptors() ?? []
    )
  }
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
extension Border0_Device_V1_DeviceManagementServiceAsyncClientProtocol {
  internal func controlStream<RequestStream>(
    _ requests: RequestStream,
    callOptions: CallOptions? = nil
  ) -> GRPCAsyncResponseStream<Border0_Device_V1_ServerToDeviceMessage> where RequestStream: Sequence, RequestStream.Element == Border0_Device_V1_DeviceToServerMessage {
    return self.performAsyncBidirectionalStreamingCall(
      path: Border0_Device_V1_DeviceManagementServiceClientMetadata.Methods.controlStream.path,
      requests: requests,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeControlStreamInterceptors() ?? []
    )
  }

  internal func controlStream<RequestStream>(
    _ requests: RequestStream,
    callOptions: CallOptions? = nil
  ) -> GRPCAsyncResponseStream<Border0_Device_V1_ServerToDeviceMessage> where RequestStream: AsyncSequence & Sendable, RequestStream.Element == Border0_Device_V1_DeviceToServerMessage {
    return self.performAsyncBidirectionalStreamingCall(
      path: Border0_Device_V1_DeviceManagementServiceClientMetadata.Methods.controlStream.path,
      requests: requests,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeControlStreamInterceptors() ?? []
    )
  }
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
internal struct Border0_Device_V1_DeviceManagementServiceAsyncClient: Border0_Device_V1_DeviceManagementServiceAsyncClientProtocol {
  internal var channel: GRPCChannel
  internal var defaultCallOptions: CallOptions
  internal var interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol?

  internal init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self.defaultCallOptions = defaultCallOptions
    self.interceptors = interceptors
  }
}

internal protocol Border0_Device_V1_DeviceManagementServiceClientInterceptorFactoryProtocol: Sendable {

  /// - Returns: Interceptors to use when invoking 'controlStream'.
  func makeControlStreamInterceptors() -> [ClientInterceptor<Border0_Device_V1_DeviceToServerMessage, Border0_Device_V1_ServerToDeviceMessage>]
}

internal enum Border0_Device_V1_DeviceManagementServiceClientMetadata {
  internal static let serviceDescriptor = GRPCServiceDescriptor(
    name: "DeviceManagementService",
    fullName: "border0.device.v1.DeviceManagementService",
    methods: [
      Border0_Device_V1_DeviceManagementServiceClientMetadata.Methods.controlStream,
    ]
  )

  internal enum Methods {
    internal static let controlStream = GRPCMethodDescriptor(
      name: "ControlStream",
      path: "/border0.device.v1.DeviceManagementService/ControlStream",
      type: GRPCCallType.bidirectionalStreaming
    )
  }
}

/// the stream (rpc) service for the Border0 api to manage devices
///
/// To build a server, implement a class that conforms to this protocol.
internal protocol Border0_Device_V1_DeviceManagementServiceProvider: CallHandlerProvider {
  var interceptors: Border0_Device_V1_DeviceManagementServiceServerInterceptorFactoryProtocol? { get }

  func controlStream(context: StreamingResponseCallContext<Border0_Device_V1_ServerToDeviceMessage>) -> EventLoopFuture<(StreamEvent<Border0_Device_V1_DeviceToServerMessage>) -> Void>
}

extension Border0_Device_V1_DeviceManagementServiceProvider {
  internal var serviceName: Substring {
    return Border0_Device_V1_DeviceManagementServiceServerMetadata.serviceDescriptor.fullName[...]
  }

  /// Determines, calls and returns the appropriate request handler, depending on the request's method.
  /// Returns nil for methods not handled by this service.
  internal func handle(
    method name: Substring,
    context: CallHandlerContext
  ) -> GRPCServerHandlerProtocol? {
    switch name {
    case "ControlStream":
      return BidirectionalStreamingServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Border0_Device_V1_DeviceToServerMessage>(),
        responseSerializer: ProtobufSerializer<Border0_Device_V1_ServerToDeviceMessage>(),
        interceptors: self.interceptors?.makeControlStreamInterceptors() ?? [],
        observerFactory: self.controlStream(context:)
      )

    default:
      return nil
    }
  }
}

/// the stream (rpc) service for the Border0 api to manage devices
///
/// To implement a server, implement an object which conforms to this protocol.
@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
internal protocol Border0_Device_V1_DeviceManagementServiceAsyncProvider: CallHandlerProvider, Sendable {
  static var serviceDescriptor: GRPCServiceDescriptor { get }
  var interceptors: Border0_Device_V1_DeviceManagementServiceServerInterceptorFactoryProtocol? { get }

  func controlStream(
    requestStream: GRPCAsyncRequestStream<Border0_Device_V1_DeviceToServerMessage>,
    responseStream: GRPCAsyncResponseStreamWriter<Border0_Device_V1_ServerToDeviceMessage>,
    context: GRPCAsyncServerCallContext
  ) async throws
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
extension Border0_Device_V1_DeviceManagementServiceAsyncProvider {
  internal static var serviceDescriptor: GRPCServiceDescriptor {
    return Border0_Device_V1_DeviceManagementServiceServerMetadata.serviceDescriptor
  }

  internal var serviceName: Substring {
    return Border0_Device_V1_DeviceManagementServiceServerMetadata.serviceDescriptor.fullName[...]
  }

  internal var interceptors: Border0_Device_V1_DeviceManagementServiceServerInterceptorFactoryProtocol? {
    return nil
  }

  internal func handle(
    method name: Substring,
    context: CallHandlerContext
  ) -> GRPCServerHandlerProtocol? {
    switch name {
    case "ControlStream":
      return GRPCAsyncServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Border0_Device_V1_DeviceToServerMessage>(),
        responseSerializer: ProtobufSerializer<Border0_Device_V1_ServerToDeviceMessage>(),
        interceptors: self.interceptors?.makeControlStreamInterceptors() ?? [],
        wrapping: { try await self.controlStream(requestStream: $0, responseStream: $1, context: $2) }
      )

    default:
      return nil
    }
  }
}

internal protocol Border0_Device_V1_DeviceManagementServiceServerInterceptorFactoryProtocol: Sendable {

  /// - Returns: Interceptors to use when handling 'controlStream'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeControlStreamInterceptors() -> [ServerInterceptor<Border0_Device_V1_DeviceToServerMessage, Border0_Device_V1_ServerToDeviceMessage>]
}

internal enum Border0_Device_V1_DeviceManagementServiceServerMetadata {
  internal static let serviceDescriptor = GRPCServiceDescriptor(
    name: "DeviceManagementService",
    fullName: "border0.device.v1.DeviceManagementService",
    methods: [
      Border0_Device_V1_DeviceManagementServiceServerMetadata.Methods.controlStream,
    ]
  )

  internal enum Methods {
    internal static let controlStream = GRPCMethodDescriptor(
      name: "ControlStream",
      path: "/border0.device.v1.DeviceManagementService/ControlStream",
      type: GRPCCallType.bidirectionalStreaming
    )
  }
}

package border0.device.v1

import border0.device.v1.DeviceManagementServiceGrpc.getServiceDescriptor
import io.grpc.CallOptions
import io.grpc.CallOptions.DEFAULT
import io.grpc.Channel
import io.grpc.Metadata
import io.grpc.MethodDescriptor
import io.grpc.ServerServiceDefinition
import io.grpc.ServerServiceDefinition.builder
import io.grpc.ServiceDescriptor
import io.grpc.Status.UNIMPLEMENTED
import io.grpc.StatusException
import io.grpc.kotlin.AbstractCoroutineServerImpl
import io.grpc.kotlin.AbstractCoroutineStub
import io.grpc.kotlin.ClientCalls.bidiStreamingRpc
import io.grpc.kotlin.ServerCalls.bidiStreamingServerMethodDefinition
import io.grpc.kotlin.StubFor
import kotlin.String
import kotlin.coroutines.CoroutineContext
import kotlin.coroutines.EmptyCoroutineContext
import kotlin.jvm.JvmOverloads
import kotlin.jvm.JvmStatic
import kotlinx.coroutines.flow.Flow

/**
 * Holder for Kotlin coroutine-based client and server APIs for
 * border0.device.v1.DeviceManagementService.
 */
public object DeviceManagementServiceGrpcKt {
  public const val SERVICE_NAME: String = DeviceManagementServiceGrpc.SERVICE_NAME

  @JvmStatic
  public val serviceDescriptor: ServiceDescriptor
    get() = getServiceDescriptor()

  public val controlStreamMethod:
      MethodDescriptor<Device.DeviceToServerMessage, Device.ServerToDeviceMessage>
    @JvmStatic
    get() = DeviceManagementServiceGrpc.getControlStreamMethod()

  /**
   * A stub for issuing RPCs to a(n) border0.device.v1.DeviceManagementService service as suspending
   * coroutines.
   */
  @StubFor(DeviceManagementServiceGrpc::class)
  public class DeviceManagementServiceCoroutineStub @JvmOverloads constructor(
    channel: Channel,
    callOptions: CallOptions = DEFAULT,
  ) : AbstractCoroutineStub<DeviceManagementServiceCoroutineStub>(channel, callOptions) {
    override fun build(channel: Channel, callOptions: CallOptions):
        DeviceManagementServiceCoroutineStub = DeviceManagementServiceCoroutineStub(channel,
        callOptions)

    /**
     * Returns a [Flow] that, when collected, executes this RPC and emits responses from the
     * server as they arrive.  That flow finishes normally if the server closes its response with
     * [`Status.OK`][io.grpc.Status], and fails by throwing a [StatusException] otherwise.  If
     * collecting the flow downstream fails exceptionally (including via cancellation), the RPC
     * is cancelled with that exception as a cause.
     *
     * The [Flow] of requests is collected once each time the [Flow] of responses is
     * collected. If collection of the [Flow] of responses completes normally or
     * exceptionally before collection of `requests` completes, the collection of
     * `requests` is cancelled.  If the collection of `requests` completes
     * exceptionally for any other reason, then the collection of the [Flow] of responses
     * completes exceptionally for the same reason and the RPC is cancelled with that reason.
     *
     * @param requests A [Flow] of request messages.
     *
     * @param headers Metadata to attach to the request.  Most users will not need this.
     *
     * @return A flow that, when collected, emits the responses from the server.
     */
    public fun controlStream(requests: Flow<Device.DeviceToServerMessage>, headers: Metadata =
        Metadata()): Flow<Device.ServerToDeviceMessage> = bidiStreamingRpc(
      channel,
      DeviceManagementServiceGrpc.getControlStreamMethod(),
      requests,
      callOptions,
      headers
    )
  }

  /**
   * Skeletal implementation of the border0.device.v1.DeviceManagementService service based on
   * Kotlin coroutines.
   */
  public abstract class DeviceManagementServiceCoroutineImplBase(
    coroutineContext: CoroutineContext = EmptyCoroutineContext,
  ) : AbstractCoroutineServerImpl(coroutineContext) {
    /**
     * Returns a [Flow] of responses to an RPC for
     * border0.device.v1.DeviceManagementService.ControlStream.
     *
     * If creating or collecting the returned flow fails with a [StatusException], the RPC
     * will fail with the corresponding [io.grpc.Status].  If it fails with a
     * [java.util.concurrent.CancellationException], the RPC will fail with status
     * `Status.CANCELLED`.  If creating
     * or collecting the returned flow fails for any other reason, the RPC will fail with
     * `Status.UNKNOWN` with the exception as a cause.
     *
     * @param requests A [Flow] of requests from the client.  This flow can be
     *        collected only once and throws [java.lang.IllegalStateException] on attempts to
     * collect
     *        it more than once.
     */
    public open fun controlStream(requests: Flow<Device.DeviceToServerMessage>):
        Flow<Device.ServerToDeviceMessage> = throw
        StatusException(UNIMPLEMENTED.withDescription("Method border0.device.v1.DeviceManagementService.ControlStream is unimplemented"))

    final override fun bindService(): ServerServiceDefinition = builder(getServiceDescriptor())
      .addMethod(bidiStreamingServerMethodDefinition(
      context = this.context,
      descriptor = DeviceManagementServiceGrpc.getControlStreamMethod(),
      implementation = ::controlStream
    )).build()
  }
}

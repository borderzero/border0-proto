//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: messages.proto

package border0.common.v1;

@kotlin.jvm.JvmName("-initializewireGuardEndpointAddr")
public inline fun wireGuardEndpointAddr(block: border0.common.v1.WireGuardEndpointAddrKt.Dsl.() -> kotlin.Unit): border0.common.v1.Messages.WireGuardEndpointAddr =
  border0.common.v1.WireGuardEndpointAddrKt.Dsl._create(border0.common.v1.Messages.WireGuardEndpointAddr.newBuilder()).apply { block() }._build()
public object WireGuardEndpointAddrKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.common.v1.Messages.WireGuardEndpointAddr.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.common.v1.Messages.WireGuardEndpointAddr.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.common.v1.Messages.WireGuardEndpointAddr = _builder.build()

    /**
     * <pre>
     * true if address was discovered via STUN
     * </pre>
     *
     * <code>bool from_stun = 1;</code>
     */
    public var fromStun: kotlin.Boolean
      @JvmName("getFromStun")
      get() = _builder.getFromStun()
      @JvmName("setFromStun")
      set(value) {
        _builder.setFromStun(value)
      }
    /**
     * <pre>
     * true if address was discovered via STUN
     * </pre>
     *
     * <code>bool from_stun = 1;</code>
     */
    public fun clearFromStun() {
      _builder.clearFromStun()
    }

    /**
     * <pre>
     * the name of the local interface that has the address (e.g., "eth0"), empty if from_stun is true.
     * </pre>
     *
     * <code>string iface_name = 2;</code>
     */
    public var ifaceName: kotlin.String
      @JvmName("getIfaceName")
      get() = _builder.getIfaceName()
      @JvmName("setIfaceName")
      set(value) {
        _builder.setIfaceName(value)
      }
    /**
     * <pre>
     * the name of the local interface that has the address (e.g., "eth0"), empty if from_stun is true.
     * </pre>
     *
     * <code>string iface_name = 2;</code>
     */
    public fun clearIfaceName() {
      _builder.clearIfaceName()
    }

    /**
     * <pre>
     * the CIDR block this IP belongs to (e.g., "192.168.0.0/24"), empty if from_stun is true.
     * </pre>
     *
     * <code>string iface_cidr = 3;</code>
     */
    public var ifaceCidr: kotlin.String
      @JvmName("getIfaceCidr")
      get() = _builder.getIfaceCidr()
      @JvmName("setIfaceCidr")
      set(value) {
        _builder.setIfaceCidr(value)
      }
    /**
     * <pre>
     * the CIDR block this IP belongs to (e.g., "192.168.0.0/24"), empty if from_stun is true.
     * </pre>
     *
     * <code>string iface_cidr = 3;</code>
     */
    public fun clearIfaceCidr() {
      _builder.clearIfaceCidr()
    }

    /**
     * <pre>
     * the actual IP address
     * </pre>
     *
     * <code>string ip_address = 4;</code>
     */
    public var ipAddress: kotlin.String
      @JvmName("getIpAddress")
      get() = _builder.getIpAddress()
      @JvmName("setIpAddress")
      set(value) {
        _builder.setIpAddress(value)
      }
    /**
     * <pre>
     * the actual IP address
     * </pre>
     *
     * <code>string ip_address = 4;</code>
     */
    public fun clearIpAddress() {
      _builder.clearIpAddress()
    }

    /**
     * <pre>
     * the UDP port number used for WireGuard
     * </pre>
     *
     * <code>uint32 port = 5;</code>
     */
    public var port: kotlin.Int
      @JvmName("getPort")
      get() = _builder.getPort()
      @JvmName("setPort")
      set(value) {
        _builder.setPort(value)
      }
    /**
     * <pre>
     * the UDP port number used for WireGuard
     * </pre>
     *
     * <code>uint32 port = 5;</code>
     */
    public fun clearPort() {
      _builder.clearPort()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.common.v1.Messages.WireGuardEndpointAddr.copy(block: border0.common.v1.WireGuardEndpointAddrKt.Dsl.() -> kotlin.Unit): border0.common.v1.Messages.WireGuardEndpointAddr =
  border0.common.v1.WireGuardEndpointAddrKt.Dsl._create(this.toBuilder()).apply { block() }._build()


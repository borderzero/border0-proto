//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: messages.proto

package border0.common.v1;

@kotlin.jvm.JvmName("-initializediscoveryDetailsMessage")
public inline fun discoveryDetailsMessage(block: border0.common.v1.DiscoveryDetailsMessageKt.Dsl.() -> kotlin.Unit): border0.common.v1.Messages.DiscoveryDetailsMessage =
  border0.common.v1.DiscoveryDetailsMessageKt.Dsl._create(border0.common.v1.Messages.DiscoveryDetailsMessage.newBuilder()).apply { block() }._build()
public object DiscoveryDetailsMessageKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.common.v1.Messages.DiscoveryDetailsMessage.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.common.v1.Messages.DiscoveryDetailsMessage.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.common.v1.Messages.DiscoveryDetailsMessage = _builder.build()

    /**
     * <pre>
     * whether the peer should be returned along with the networks it's in
     * </pre>
     *
     * <code>bool discoverable = 1;</code>
     */
    public var discoverable: kotlin.Boolean
      @JvmName("getDiscoverable")
      get() = _builder.getDiscoverable()
      @JvmName("setDiscoverable")
      set(value) {
        _builder.setDiscoverable(value)
      }
    /**
     * <pre>
     * whether the peer should be returned along with the networks it's in
     * </pre>
     *
     * <code>bool discoverable = 1;</code>
     */
    public fun clearDiscoverable() {
      _builder.clearDiscoverable()
    }

    /**
     * <pre>
     * the peer's public IP WireGuard endpoint (IPv4 + port) discovered via STUN
     * </pre>
     *
     * <code>string endpoint_public_udp4 = 2;</code>
     */
    public var endpointPublicUdp4: kotlin.String
      @JvmName("getEndpointPublicUdp4")
      get() = _builder.getEndpointPublicUdp4()
      @JvmName("setEndpointPublicUdp4")
      set(value) {
        _builder.setEndpointPublicUdp4(value)
      }
    /**
     * <pre>
     * the peer's public IP WireGuard endpoint (IPv4 + port) discovered via STUN
     * </pre>
     *
     * <code>string endpoint_public_udp4 = 2;</code>
     */
    public fun clearEndpointPublicUdp4() {
      _builder.clearEndpointPublicUdp4()
    }

    /**
     * <pre>
     * the peer's public IP WireGuard endpoint (IPv6 + port) discovered via STUN
     * </pre>
     *
     * <code>string endpoint_public_udp6 = 3;</code>
     */
    public var endpointPublicUdp6: kotlin.String
      @JvmName("getEndpointPublicUdp6")
      get() = _builder.getEndpointPublicUdp6()
      @JvmName("setEndpointPublicUdp6")
      set(value) {
        _builder.setEndpointPublicUdp6(value)
      }
    /**
     * <pre>
     * the peer's public IP WireGuard endpoint (IPv6 + port) discovered via STUN
     * </pre>
     *
     * <code>string endpoint_public_udp6 = 3;</code>
     */
    public fun clearEndpointPublicUdp6() {
      _builder.clearEndpointPublicUdp6()
    }

    /**
     * <pre>
     * the public key of the peer, only used in connector
     * </pre>
     *
     * <code>string public_key = 4;</code>
     */
    public var publicKey: kotlin.String
      @JvmName("getPublicKey")
      get() = _builder.getPublicKey()
      @JvmName("setPublicKey")
      set(value) {
        _builder.setPublicKey(value)
      }
    /**
     * <pre>
     * the public key of the peer, only used in connector
     * </pre>
     *
     * <code>string public_key = 4;</code>
     */
    public fun clearPublicKey() {
      _builder.clearPublicKey()
    }

    /**
     * An uninstantiable, behaviorless type to represent the field in
     * generics.
     */
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
    public class WgEpAddrsProxy private constructor() : com.google.protobuf.kotlin.DslProxy()
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     */
     public val wgEpAddrs: com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>
      @kotlin.jvm.JvmSynthetic
      get() = com.google.protobuf.kotlin.DslList(
        _builder.getWgEpAddrsList()
      )
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     * @param value The wgEpAddrs to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("addWgEpAddrs")
    public fun com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>.add(value: border0.common.v1.Messages.WireGuardEndpointAddr) {
      _builder.addWgEpAddrs(value)
    }
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     * @param value The wgEpAddrs to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("plusAssignWgEpAddrs")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>.plusAssign(value: border0.common.v1.Messages.WireGuardEndpointAddr) {
      add(value)
    }
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     * @param values The wgEpAddrs to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("addAllWgEpAddrs")
    public fun com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>.addAll(values: kotlin.collections.Iterable<border0.common.v1.Messages.WireGuardEndpointAddr>) {
      _builder.addAllWgEpAddrs(values)
    }
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     * @param values The wgEpAddrs to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("plusAssignAllWgEpAddrs")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>.plusAssign(values: kotlin.collections.Iterable<border0.common.v1.Messages.WireGuardEndpointAddr>) {
      addAll(values)
    }
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     * @param index The index to set the value at.
     * @param value The wgEpAddrs to set.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("setWgEpAddrs")
    public operator fun com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>.set(index: kotlin.Int, value: border0.common.v1.Messages.WireGuardEndpointAddr) {
      _builder.setWgEpAddrs(index, value)
    }
    /**
     * <pre>
     * all ip address and port combinations that the peer can use to tx/rx traffic
     * </pre>
     *
     * <code>repeated .border0.common.v1.WireGuardEndpointAddr wg_ep_addrs = 5;</code>
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("clearWgEpAddrs")
    public fun com.google.protobuf.kotlin.DslList<border0.common.v1.Messages.WireGuardEndpointAddr, WgEpAddrsProxy>.clear() {
      _builder.clearWgEpAddrs()
    }

  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.common.v1.Messages.DiscoveryDetailsMessage.copy(block: border0.common.v1.DiscoveryDetailsMessageKt.Dsl.() -> kotlin.Unit): border0.common.v1.Messages.DiscoveryDetailsMessage =
  border0.common.v1.DiscoveryDetailsMessageKt.Dsl._create(this.toBuilder()).apply { block() }._build()


//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: connector.proto

package border0.v1;

@kotlin.jvm.JvmName("-initializedisconnect")
public inline fun disconnect(block: border0.v1.DisconnectKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.Disconnect =
  border0.v1.DisconnectKt.Dsl._create(border0.v1.Connector.Disconnect.newBuilder()).apply { block() }._build()
public object DisconnectKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.v1.Connector.Disconnect.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.v1.Connector.Disconnect.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.v1.Connector.Disconnect = _builder.build()

    /**
     * <code>string reason = 1;</code>
     */
    public var reason: kotlin.String
      @JvmName("getReason")
      get() = _builder.getReason()
      @JvmName("setReason")
      set(value) {
        _builder.setReason(value)
      }
    /**
     * <code>string reason = 1;</code>
     */
    public fun clearReason() {
      _builder.clearReason()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.v1.Connector.Disconnect.copy(block: border0.v1.DisconnectKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.Disconnect =
  border0.v1.DisconnectKt.Dsl._create(this.toBuilder()).apply { block() }._build()


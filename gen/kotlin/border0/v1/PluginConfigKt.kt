//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: connector.proto

package border0.v1;

@kotlin.jvm.JvmName("-initializepluginConfig")
public inline fun pluginConfig(block: border0.v1.PluginConfigKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.PluginConfig =
  border0.v1.PluginConfigKt.Dsl._create(border0.v1.Connector.PluginConfig.newBuilder()).apply { block() }._build()
public object PluginConfigKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.v1.Connector.PluginConfig.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.v1.Connector.PluginConfig.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.v1.Connector.PluginConfig = _builder.build()

    /**
     * <code>string id = 1;</code>
     */
    public var id: kotlin.String
      @JvmName("getId")
      get() = _builder.getId()
      @JvmName("setId")
      set(value) {
        _builder.setId(value)
      }
    /**
     * <code>string id = 1;</code>
     */
    public fun clearId() {
      _builder.clearId()
    }

    /**
     * <code>string name = 2;</code>
     */
    public var name: kotlin.String
      @JvmName("getName")
      get() = _builder.getName()
      @JvmName("setName")
      set(value) {
        _builder.setName(value)
      }
    /**
     * <code>string name = 2;</code>
     */
    public fun clearName() {
      _builder.clearName()
    }

    /**
     * <code>string type = 3;</code>
     */
    public var type: kotlin.String
      @JvmName("getType")
      get() = _builder.getType()
      @JvmName("setType")
      set(value) {
        _builder.setType(value)
      }
    /**
     * <code>string type = 3;</code>
     */
    public fun clearType() {
      _builder.clearType()
    }

    /**
     * <code>.google.protobuf.Struct config = 4;</code>
     */
    public var config: com.google.protobuf.Struct
      @JvmName("getConfig")
      get() = _builder.getConfig()
      @JvmName("setConfig")
      set(value) {
        _builder.setConfig(value)
      }
    /**
     * <code>.google.protobuf.Struct config = 4;</code>
     */
    public fun clearConfig() {
      _builder.clearConfig()
    }
    /**
     * <code>.google.protobuf.Struct config = 4;</code>
     * @return Whether the config field is set.
     */
    public fun hasConfig(): kotlin.Boolean {
      return _builder.hasConfig()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.v1.Connector.PluginConfig.copy(block: border0.v1.PluginConfigKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.PluginConfig =
  border0.v1.PluginConfigKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val border0.v1.Connector.PluginConfigOrBuilder.configOrNull: com.google.protobuf.Struct?
  get() = if (hasConfig()) getConfig() else null


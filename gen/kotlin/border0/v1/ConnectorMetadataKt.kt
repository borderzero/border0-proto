//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: connector.proto

package border0.v1;

@kotlin.jvm.JvmName("-initializeconnectorMetadata")
public inline fun connectorMetadata(block: border0.v1.ConnectorMetadataKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.ConnectorMetadata =
  border0.v1.ConnectorMetadataKt.Dsl._create(border0.v1.Connector.ConnectorMetadata.newBuilder()).apply { block() }._build()
public object ConnectorMetadataKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.v1.Connector.ConnectorMetadata.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.v1.Connector.ConnectorMetadata.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.v1.Connector.ConnectorMetadata = _builder.build()

    /**
     * <code>.google.protobuf.Struct data = 1;</code>
     */
    public var data: com.google.protobuf.Struct
      @JvmName("getData")
      get() = _builder.getData()
      @JvmName("setData")
      set(value) {
        _builder.setData(value)
      }
    /**
     * <code>.google.protobuf.Struct data = 1;</code>
     */
    public fun clearData() {
      _builder.clearData()
    }
    /**
     * <code>.google.protobuf.Struct data = 1;</code>
     * @return Whether the data field is set.
     */
    public fun hasData(): kotlin.Boolean {
      return _builder.hasData()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.v1.Connector.ConnectorMetadata.copy(block: border0.v1.ConnectorMetadataKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.ConnectorMetadata =
  border0.v1.ConnectorMetadataKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val border0.v1.Connector.ConnectorMetadataOrBuilder.dataOrNull: com.google.protobuf.Struct?
  get() = if (hasData()) getData() else null


//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: connector.proto

package border0.v1;

@kotlin.jvm.JvmName("-initializetag")
public inline fun tag(block: border0.v1.TagKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.Tag =
  border0.v1.TagKt.Dsl._create(border0.v1.Connector.Tag.newBuilder()).apply { block() }._build()
public object TagKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.v1.Connector.Tag.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.v1.Connector.Tag.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.v1.Connector.Tag = _builder.build()

    /**
     * <code>string key = 1;</code>
     */
    public var key: kotlin.String
      @JvmName("getKey")
      get() = _builder.getKey()
      @JvmName("setKey")
      set(value) {
        _builder.setKey(value)
      }
    /**
     * <code>string key = 1;</code>
     */
    public fun clearKey() {
      _builder.clearKey()
    }

    /**
     * <code>string value = 2;</code>
     */
    public var value: kotlin.String
      @JvmName("getValue")
      get() = _builder.getValue()
      @JvmName("setValue")
      set(value) {
        _builder.setValue(value)
      }
    /**
     * <code>string value = 2;</code>
     */
    public fun clearValue() {
      _builder.clearValue()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.v1.Connector.Tag.copy(block: border0.v1.TagKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.Tag =
  border0.v1.TagKt.Dsl._create(this.toBuilder()).apply { block() }._build()


//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: device.proto

package border0.device.v1;

@kotlin.jvm.JvmName("-initializeauthChallengeSolutionMessage")
public inline fun authChallengeSolutionMessage(block: border0.device.v1.AuthChallengeSolutionMessageKt.Dsl.() -> kotlin.Unit): border0.device.v1.Device.AuthChallengeSolutionMessage =
  border0.device.v1.AuthChallengeSolutionMessageKt.Dsl._create(border0.device.v1.Device.AuthChallengeSolutionMessage.newBuilder()).apply { block() }._build()
public object AuthChallengeSolutionMessageKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.device.v1.Device.AuthChallengeSolutionMessage.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.device.v1.Device.AuthChallengeSolutionMessage.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.device.v1.Device.AuthChallengeSolutionMessage = _builder.build()

    /**
     * <code>bytes solved = 1;</code>
     */
    public var solved: com.google.protobuf.ByteString
      @JvmName("getSolved")
      get() = _builder.getSolved()
      @JvmName("setSolved")
      set(value) {
        _builder.setSolved(value)
      }
    /**
     * <code>bytes solved = 1;</code>
     */
    public fun clearSolved() {
      _builder.clearSolved()
    }

    /**
     * <code>bytes solvedNonce = 2;</code>
     */
    public var solvedNonce: com.google.protobuf.ByteString
      @JvmName("getSolvedNonce")
      get() = _builder.getSolvedNonce()
      @JvmName("setSolvedNonce")
      set(value) {
        _builder.setSolvedNonce(value)
      }
    /**
     * <code>bytes solvedNonce = 2;</code>
     */
    public fun clearSolvedNonce() {
      _builder.clearSolvedNonce()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.device.v1.Device.AuthChallengeSolutionMessage.copy(block: border0.device.v1.AuthChallengeSolutionMessageKt.Dsl.() -> kotlin.Unit): border0.device.v1.Device.AuthChallengeSolutionMessage =
  border0.device.v1.AuthChallengeSolutionMessageKt.Dsl._create(this.toBuilder()).apply { block() }._build()


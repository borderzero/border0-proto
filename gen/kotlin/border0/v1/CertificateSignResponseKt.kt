//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: connector.proto

package border0.v1;

@kotlin.jvm.JvmName("-initializecertificateSignResponse")
public inline fun certificateSignResponse(block: border0.v1.CertificateSignResponseKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.CertificateSignResponse =
  border0.v1.CertificateSignResponseKt.Dsl._create(border0.v1.Connector.CertificateSignResponse.newBuilder()).apply { block() }._build()
public object CertificateSignResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.v1.Connector.CertificateSignResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.v1.Connector.CertificateSignResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.v1.Connector.CertificateSignResponse = _builder.build()

    /**
     * <code>string request_id = 1;</code>
     */
    public var requestId: kotlin.String
      @JvmName("getRequestId")
      get() = _builder.getRequestId()
      @JvmName("setRequestId")
      set(value) {
        _builder.setRequestId(value)
      }
    /**
     * <code>string request_id = 1;</code>
     */
    public fun clearRequestId() {
      _builder.clearRequestId()
    }

    /**
     * <code>bytes certificate = 2;</code>
     */
    public var certificate: com.google.protobuf.ByteString
      @JvmName("getCertificate")
      get() = _builder.getCertificate()
      @JvmName("setCertificate")
      set(value) {
        _builder.setCertificate(value)
      }
    /**
     * <code>bytes certificate = 2;</code>
     */
    public fun clearCertificate() {
      _builder.clearCertificate()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.v1.Connector.CertificateSignResponse.copy(block: border0.v1.CertificateSignResponseKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.CertificateSignResponse =
  border0.v1.CertificateSignResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()


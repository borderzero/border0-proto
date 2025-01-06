//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: connector.proto

package border0.v1;

@kotlin.jvm.JvmName("-initializecertificateSignRequest")
public inline fun certificateSignRequest(block: border0.v1.CertificateSignRequestKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.CertificateSignRequest =
  border0.v1.CertificateSignRequestKt.Dsl._create(border0.v1.Connector.CertificateSignRequest.newBuilder()).apply { block() }._build()
public object CertificateSignRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: border0.v1.Connector.CertificateSignRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: border0.v1.Connector.CertificateSignRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): border0.v1.Connector.CertificateSignRequest = _builder.build()

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
     * <code>bytes certificate_signing_request = 2;</code>
     */
    public var certificateSigningRequest: com.google.protobuf.ByteString
      @JvmName("getCertificateSigningRequest")
      get() = _builder.getCertificateSigningRequest()
      @JvmName("setCertificateSigningRequest")
      set(value) {
        _builder.setCertificateSigningRequest(value)
      }
    /**
     * <code>bytes certificate_signing_request = 2;</code>
     */
    public fun clearCertificateSigningRequest() {
      _builder.clearCertificateSigningRequest()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun border0.v1.Connector.CertificateSignRequest.copy(block: border0.v1.CertificateSignRequestKt.Dsl.() -> kotlin.Unit): border0.v1.Connector.CertificateSignRequest =
  border0.v1.CertificateSignRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()


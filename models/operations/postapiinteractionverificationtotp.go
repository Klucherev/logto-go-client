// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

// PostAPIInteractionVerificationTotpResponseBody - OK
type PostAPIInteractionVerificationTotpResponseBody struct {
	Secret       string `json:"secret"`
	SecretQrCode string `json:"secretQrCode"`
}

func (o *PostAPIInteractionVerificationTotpResponseBody) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

func (o *PostAPIInteractionVerificationTotpResponseBody) GetSecretQrCode() string {
	if o == nil {
		return ""
	}
	return o.SecretQrCode
}

type PostAPIInteractionVerificationTotpResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Object *PostAPIInteractionVerificationTotpResponseBody
}

func (o *PostAPIInteractionVerificationTotpResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostAPIInteractionVerificationTotpResponse) GetObject() *PostAPIInteractionVerificationTotpResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

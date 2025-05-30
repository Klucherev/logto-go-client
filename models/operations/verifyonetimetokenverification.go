// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

// VerifyOneTimeTokenVerificationIdentifier - The unique user identifier.  <br/> Currently, only `email` is accepted.
type VerifyOneTimeTokenVerificationIdentifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (o *VerifyOneTimeTokenVerificationIdentifier) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *VerifyOneTimeTokenVerificationIdentifier) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type VerifyOneTimeTokenVerificationRequest struct {
	// The unique user identifier.  <br/> Currently, only `email` is accepted.
	Identifier VerifyOneTimeTokenVerificationIdentifier `json:"identifier"`
	// The one-time token to be verified.
	Token string `json:"token"`
}

func (o *VerifyOneTimeTokenVerificationRequest) GetIdentifier() VerifyOneTimeTokenVerificationIdentifier {
	if o == nil {
		return VerifyOneTimeTokenVerificationIdentifier{}
	}
	return o.Identifier
}

func (o *VerifyOneTimeTokenVerificationRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// VerifyOneTimeTokenVerificationResponseBody - The one-time token was successfully verified.
type VerifyOneTimeTokenVerificationResponseBody struct {
	// The unique ID of the verification record. Required for user identification via the `Identification` API or to bind the identifier to the user's account via the `Profile` API.
	VerificationID string `json:"verificationId"`
}

func (o *VerifyOneTimeTokenVerificationResponseBody) GetVerificationID() string {
	if o == nil {
		return ""
	}
	return o.VerificationID
}

type VerifyOneTimeTokenVerificationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The one-time token was successfully verified.
	Object *VerifyOneTimeTokenVerificationResponseBody
}

func (o *VerifyOneTimeTokenVerificationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VerifyOneTimeTokenVerificationResponse) GetObject() *VerifyOneTimeTokenVerificationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

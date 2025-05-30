// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/klucherev/logto-go-client/models/components"
)

type CreateSamlApplicationSecretRequestBody struct {
	// The lifetime of the certificate in years (minimum 1 year).
	LifeSpanInYears int64 `json:"lifeSpanInYears"`
}

func (o *CreateSamlApplicationSecretRequestBody) GetLifeSpanInYears() int64 {
	if o == nil {
		return 0
	}
	return o.LifeSpanInYears
}

type CreateSamlApplicationSecretRequest struct {
	// The unique identifier of the saml application.
	ID          string                                 `pathParam:"style=simple,explode=false,name=id"`
	RequestBody CreateSamlApplicationSecretRequestBody `request:"mediaType=application/json"`
}

func (o *CreateSamlApplicationSecretRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateSamlApplicationSecretRequest) GetRequestBody() CreateSamlApplicationSecretRequestBody {
	if o == nil {
		return CreateSamlApplicationSecretRequestBody{}
	}
	return o.RequestBody
}

type CreateSamlApplicationSecretSha256 struct {
	Formatted   string `json:"formatted"`
	Unformatted string `json:"unformatted"`
}

func (o *CreateSamlApplicationSecretSha256) GetFormatted() string {
	if o == nil {
		return ""
	}
	return o.Formatted
}

func (o *CreateSamlApplicationSecretSha256) GetUnformatted() string {
	if o == nil {
		return ""
	}
	return o.Unformatted
}

type CreateSamlApplicationSecretFingerprints struct {
	Sha256 CreateSamlApplicationSecretSha256 `json:"sha256"`
}

func (o *CreateSamlApplicationSecretFingerprints) GetSha256() CreateSamlApplicationSecretSha256 {
	if o == nil {
		return CreateSamlApplicationSecretSha256{}
	}
	return o.Sha256
}

// CreateSamlApplicationSecretResponseBody - The signing certificate was created successfully.
type CreateSamlApplicationSecretResponseBody struct {
	ID           string                                  `json:"id"`
	Certificate  string                                  `json:"certificate"`
	CreatedAt    float64                                 `json:"createdAt"`
	ExpiresAt    float64                                 `json:"expiresAt"`
	Active       bool                                    `json:"active"`
	Fingerprints CreateSamlApplicationSecretFingerprints `json:"fingerprints"`
}

func (o *CreateSamlApplicationSecretResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateSamlApplicationSecretResponseBody) GetCertificate() string {
	if o == nil {
		return ""
	}
	return o.Certificate
}

func (o *CreateSamlApplicationSecretResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *CreateSamlApplicationSecretResponseBody) GetExpiresAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ExpiresAt
}

func (o *CreateSamlApplicationSecretResponseBody) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *CreateSamlApplicationSecretResponseBody) GetFingerprints() CreateSamlApplicationSecretFingerprints {
	if o == nil {
		return CreateSamlApplicationSecretFingerprints{}
	}
	return o.Fingerprints
}

type CreateSamlApplicationSecretResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The signing certificate was created successfully.
	Object *CreateSamlApplicationSecretResponseBody
}

func (o *CreateSamlApplicationSecretResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateSamlApplicationSecretResponse) GetObject() *CreateSamlApplicationSecretResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
